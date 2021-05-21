package metadata

import (
	"fmt"
	"strconv"
	"strings"
)

const POLYMORPH_IMAGE_URL string = "https://storage.googleapis.com/polymorph-images/"
const EXTERNAL_URL string = "https://kek.dao/"
const GENES_COUNT = 7
const BASE_GENES_COUNT int = 12
const HEAD_GENES_COUNT int = 10
const ARMOR_GENES_COUNT int = 10
const PANTS_GENES_COUNT int = 10
const SHOES_GENES_COUNT int = 8
const WEAPON_GENES_COUNT int = 9
const FACE_GENES_COUNT int = 5

type Genome string

func getGene(g string, start, end, count int) string {
	genomeLen := len(g)
	geneStr := g[genomeLen+start : genomeLen+end]
	gene, _ := strconv.Atoi(geneStr)
	gene = gene % count
	if gene < 10 {
		return fmt.Sprintf("0%s", strconv.Itoa(gene))
	}

	return strconv.Itoa(gene)
}

func getFaceGene(g string) string {
	return getGene(g, -14, -12, FACE_GENES_COUNT)
}

func getWeaponGene(g string) string {
	return getGene(g, -12, -10, WEAPON_GENES_COUNT)
}

func getShoesGene(g string) string {
	return getGene(g, -10, -8, SHOES_GENES_COUNT)
}

func getPantsGene(g string) string {
	return getGene(g, -8, -6, PANTS_GENES_COUNT)
}

func getArmorGene(g string) string {
	return getGene(g, -6, -4, ARMOR_GENES_COUNT)
}

func getHeadGene(g string) string {
	return getGene(g, -4, -2, HEAD_GENES_COUNT)
}

func getBaseGene(g string) string {
	return getGene(g, -2, 0, BASE_GENES_COUNT)
}

func (g *Genome) genes() []string {
	res := make([]string, 0, GENES_COUNT)
	gStr := string(*g)

	res = append(res, getFaceGene(gStr))
	res = append(res, getWeaponGene(gStr))
	res = append(res, getShoesGene(gStr))
	res = append(res, getPantsGene(gStr))
	res = append(res, getArmorGene(gStr))
	res = append(res, getHeadGene(gStr))
	res = append(res, getBaseGene(gStr))

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

	b.WriteString(".png") // Finish with png extension

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
