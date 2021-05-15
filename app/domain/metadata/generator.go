package metadata

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/disintegration/imaging"
	log "github.com/sirupsen/logrus"
	"image"
	"image/color"
	"net/http"
	"strings"
)

const IMG_SIZE = 4000
const GCLOUD_BUCKET_NAME = "polymorph-images"

func imageExists(imageURL string) bool {
	resp, err := http.Get(imageURL)
	if err != nil {
		log.Fatalln(err)
	}

	return resp.StatusCode != 404
}

func combineImages(basePath string, overlayPaths ...string) *image.NRGBA {
	base, err := imaging.Open(basePath)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	dst := imaging.New(IMG_SIZE, IMG_SIZE, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, base, image.Pt(0, 0))

	for _, op := range overlayPaths {
		o, err := imaging.Open(op)
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}
		dst = imaging.Overlay(dst, o, image.Pt(0, 0), 1)
	}

	return dst
}

func reverseGenesOrder(genes []string) []string {
	res := make([]string, 0, len(genes))
	for i := len(genes) - 1; i >= 0; i-- {
		res = append(res, genes[i])
	}
	return res
}

func saveToGCloud(i *image.NRGBA, name string) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	wc := client.Bucket(GCLOUD_BUCKET_NAME).Object(name).NewWriter(ctx)
	f, err := imaging.FormatFromFilename(name)
	if err != nil {
		log.Errorf("Format from filename: %v", err)
	}
	err = imaging.Encode(wc, i, f, imaging.JPEGQuality(100))

	if err != nil {
		log.Errorf("Upload: %v", err)
	}

	if err = wc.Close(); err != nil {
		log.Errorf("Writer.Close: %v", err)
	}

}

func generateAndSaveImage(genes []string) {
	// Reverse
	revGenes := reverseGenesOrder(genes)

	f := make([]string, len(genes))

	for i, gene := range revGenes {
		f[i] = fmt.Sprintf("./images/%v/%s.png", i, gene)
	}
	i := combineImages(f[0], f[1:]...)

	b := strings.Builder{}

	for _, gene := range genes {
		b.WriteString(gene)
	}

	b.WriteString(".png") // Finish with png extension

	saveToGCloud(i, b.String())
}
