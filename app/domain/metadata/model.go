package metadata

import (
	"fmt"
	"strconv"
	"strings"
)

const POLYMORPH_IMAGE_URL string = "https://storage.googleapis.com/polymorph-images/"
const EXTERNAL_URL string = "https://universe.xyz/polymorphs/"
const GENES_COUNT = 9
const BACKGROUND_GENE_COUNT int = 12
const BASE_GENES_COUNT int = 12
const PANTS_GENES_COUNT int = 24
const TORSO_GENES_COUNT int = 22
const SHOES_GENES_COUNT int = 17
const FACE_GENES_COUNT int = 9
const HEAD_GENES_COUNT int = 28
const WEAPON_RIGHT_GENES_COUNT int = 5
const WEAPON_LEFT_GENES_COUNT int = 17

type Genome string
type Gene int

func (g Gene) toPath() string {
	if g < 10 {
		return fmt.Sprintf("0%s", strconv.Itoa(int(g)))
	}

	return strconv.Itoa(int(g))
}

func getGeneInt(g string, start, end, count int) int {
	genomeLen := len(g)
	geneStr := g[genomeLen+start : genomeLen+end]
	gene, _ := strconv.Atoi(geneStr)
	return gene % count
}

func getGene(g string, start, end, count int) string {
	gene := getGeneInt(g, start, end, count)
	return Gene(gene).toPath()
}

func getWeaponLeftGene(g string) string {
	return getGene(g, -18, -16, WEAPON_LEFT_GENES_COUNT)
}

func getWeaponRightGene(g string) string {
	return getGene(g, -16, -14, WEAPON_RIGHT_GENES_COUNT)
}

func getHeadGene(g string) string {
	return getGene(g, -14, -12, HEAD_GENES_COUNT)
}

func getFaceGene(g string) string {
	return getGene(g, -12, -10, FACE_GENES_COUNT)
}

func getShoesGene(g string) string {
	return getGene(g, -10, -8, SHOES_GENES_COUNT)
}

func getTorsoGene(g string) string {
	return getGene(g, -8, -6, TORSO_GENES_COUNT)
}

func getPantsGene(g string) string {
	return getGene(g, -6, -4, PANTS_GENES_COUNT)
}

func getBackgroundGene(g string) string {
	return getGene(g, -4, -2, BACKGROUND_GENE_COUNT)
}

func getBaseGene(g string) string {
	return getGene(g, -2, 0, BASE_GENES_COUNT)
}

func (g *Genome) genes() []string {
	res := make([]string, 0, GENES_COUNT)
	gStr := string(*g)

	res = append(res, getWeaponLeftGene(gStr))
	res = append(res, getWeaponRightGene(gStr))
	res = append(res, getHeadGene(gStr))
	res = append(res, getFaceGene(gStr))
	res = append(res, getTorsoGene(gStr))
	res = append(res, getPantsGene(gStr))
	res = append(res, getShoesGene(gStr))
	res = append(res, getBaseGene(gStr))
	res = append(res, getBackgroundGene(gStr))

	return res
}

func (g *Genome) Metadata(tokenId string) Metadata {
	var m Metadata
	genes := g.genes()

	m.Name = fmt.Sprintf("Polymorph with Id %v", tokenId)
	m.Description = fmt.Sprintf("Polymorph with Id %v and Genome %v", tokenId, string(*g))
	m.Attributes = make([]string, 0, 0)
	m.ExternalUrl = fmt.Sprintf("%s%s", EXTERNAL_URL, tokenId)

	b := strings.Builder{}

	b.WriteString(POLYMORPH_IMAGE_URL) // Start with base url

	for _, gene := range genes {
		b.WriteString(gene)
	}

	b.WriteString(".jpg") // Finish with jpg extension

	imageURL := b.String()

	imageExists := imageExists(imageURL)

	if !imageExists {
		generateAndSaveImage(genes)
	}

	m.Image = imageURL
	return m
}

type Metadata struct {
	Description string   `json:"description"`
	Name        string   `json:"name"`
	Image       string   `json:"image"`
	Attributes  []string `json:"attributes"`
	ExternalUrl string   `json:"external_url"`
}
