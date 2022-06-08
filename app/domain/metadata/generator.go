package metadata

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/json"
	"fmt"
	"github.com/disintegration/imaging"
	log "github.com/sirupsen/logrus"
	"image"
	"image/color"
	"net/http"
	"strings"
)

const IMG_SIZE = 4000
const GCLOUD_UPLOAD_BUCKET_NAME = "polymorphs-v1-test"
const GCLOUD_SOURCE_BUCKET_NAME = "polymorph-source-images"
const IFRAME_CLOUD_FUNCTION_URL = "https://us-central1-polymorphmetadata.cloudfunctions.net/rinkeby-iframe?id="
const IFRAME_HTMLS_BUCKET_NAME = "iframe-htmls"

func imageExists(imageURL string) bool {
	resp, err := http.Get(imageURL)
	if err != nil {
		log.Fatalln(err)
	}

	return resp.StatusCode != 404
}

func objectExists(objectURL string) bool {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(IFRAME_HTMLS_BUCKET_NAME)
	stats, err := bucket.Object(objectURL).Attrs(ctx)

	return stats != nil
}

// Queries iframe cloud function so that the animation_url is automatically generated from metadata function
func generateIframeAnimation(polymorphID *string) string {
	getUrl := IFRAME_CLOUD_FUNCTION_URL + *polymorphID
	resp, err := http.Get(getUrl)
	if err != nil {
		log.Errorf("Couldn't query iframe-cloud-function. Original error: %v", err)
	}

	type iframeResponse struct {
		AnimationUrl string `json:"animation_url"`
	}

	iframeResp := iframeResponse{}

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&iframeResp)
	return iframeResp.AnimationUrl
}

func combineRemoteImages(bucket *storage.BucketHandle, basePath string, overlayPaths ...string) *image.NRGBA {

	ctx := context.Background()

	baseReader, err := bucket.Object(basePath).NewReader(ctx)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	defer baseReader.Close()
	base, err := imaging.Decode(baseReader)
	if err != nil {
		log.Fatalf("failed to decode image: %v", err)
	}
	dst := imaging.New(IMG_SIZE, IMG_SIZE, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, base, image.Pt(0, 0))

	for _, op := range overlayPaths {
		r, err := bucket.Object(op).NewReader(ctx)
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}
		defer r.Close()
		o, err := imaging.Decode(r)
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

	bucket := client.Bucket(GCLOUD_UPLOAD_BUCKET_NAME).Object(name).NewWriter(ctx)
	// f, err := imaging.FormatFromFilename(name)
	// if err != nil {
	// 	log.Errorf("Format from filename: %v", err)
	// }
	err = imaging.Encode(bucket, i, imaging.JPEG, imaging.JPEGQuality(80))

	if err != nil {
		log.Errorf("Upload: %v", err)
	}

	if err = bucket.Close(); err != nil {
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

	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(GCLOUD_SOURCE_BUCKET_NAME)

	i := combineRemoteImages(bucket, f[0], f[1:]...)

	b := strings.Builder{}

	for _, gene := range genes {
		b.WriteString(gene)
	}

	b.WriteString(".jpg") // Finish with jpg extension

	saveToGCloud(i, b.String())
}
