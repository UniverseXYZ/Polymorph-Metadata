package metadata

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/polymorph-metadata/app/config"
)

const POLYMORPH_IMAGE_URL string = "https://storage.googleapis.com/polymorph-images/"
const EXTERNAL_URL string = "https://universe.xyz/polymorphs/"
const GENES_COUNT = 9
const BACKGROUND_GENE_COUNT int = 12
const BASE_GENES_COUNT int = 11
const SHOES_GENES_COUNT int = 25
const PANTS_GENES_COUNT int = 33
const TORSO_GENES_COUNT int = 34
const EYEWEAR_GENES_COUNT int = 13
const HEAD_GENES_COUNT int = 31
const WEAPON_RIGHT_GENES_COUNT int = 32
const WEAPON_LEFT_GENES_COUNT int = 32

type Genome string
type Gene int
type Attribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

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

func getWeaponLeftGene(g string) int {
	return getGeneInt(g, -18, -16, WEAPON_LEFT_GENES_COUNT)
}

func getWeaponLeftGeneAttribute(g string, configService *config.ConfigService) Attribute {
	gene := getWeaponLeftGene(g)
	return Attribute{
		TraitType: "WeaponLeft",
		Value:     configService.WeaponLeft[gene],
	}
}

func getWeaponLeftGenePath(g string) string {
	gene := getWeaponLeftGene(g)
	return Gene(gene).toPath()
}

func getWeaponRightGene(g string) int {
	return getGeneInt(g, -16, -14, WEAPON_RIGHT_GENES_COUNT)
}

func getWeaponRightGeneAttribute(g string, configService *config.ConfigService) Attribute {
	gene := getWeaponRightGene(g)
	return Attribute{
		TraitType: "WeaponRight",
		Value:     configService.WeaponRight[gene],
	}
}

func getWeaponRightGenePath(g string) string {
	gene := getWeaponRightGene(g)
	return Gene(gene).toPath()
}

func getHeadGene(g string) int {
	return getGeneInt(g, -14, -12, HEAD_GENES_COUNT)
}

func getHeadGeneAttribute(g string, configService *config.ConfigService) Attribute {
	gene := getHeadGene(g)
	return Attribute{
		TraitType: "Headwear",
		Value:     configService.Headwear[gene],
	}
}

func getHeadGenePath(g string) string {
	gene := getHeadGene(g)
	return Gene(gene).toPath()
}

func getEyewearGene(g string) int {
	return getGeneInt(g, -12, -10, EYEWEAR_GENES_COUNT)
}

func getEyewearGeneAttribute(g string, configService *config.ConfigService) Attribute {
	gene := getEyewearGene(g)
	return Attribute{
		TraitType: "Eyewear",
		Value:     configService.Eyewear[gene],
	}
}

func getEyewearGenePath(g string) string {
	gene := getEyewearGene(g)
	return Gene(gene).toPath()
}

func getShoesGene(g string) int {
	return getGeneInt(g, -10, -8, SHOES_GENES_COUNT)
}

func getShoesGeneAttribute(g string, configService *config.ConfigService) Attribute {
	gene := getShoesGene(g)
	return Attribute{
		TraitType: "Footwear",
		Value:     configService.Footwear[gene],
	}
}

func getShoesGenePath(g string) string {
	gene := getShoesGene(g)
	return Gene(gene).toPath()
}

func getTorsoGene(g string) int {
	return getGeneInt(g, -8, -6, TORSO_GENES_COUNT)
}

func getTorsoGeneAttribute(g string, configService *config.ConfigService) Attribute {
	gene := getTorsoGene(g)
	return Attribute{
		TraitType: "Torso",
		Value:     configService.Torso[gene],
	}
}

func getTorsoGenePath(g string) string {
	gene := getTorsoGene(g)
	return Gene(gene).toPath()
}

func getPantsGene(g string) int {
	return getGeneInt(g, -6, -4, PANTS_GENES_COUNT)
}

func getPantsGeneAttribute(g string, configService *config.ConfigService) Attribute {
	gene := getPantsGene(g)
	return Attribute{
		TraitType: "Pants",
		Value:     configService.Pants[gene],
	}
}

func getPantsGenePath(g string) string {
	gene := getPantsGene(g)
	return Gene(gene).toPath()
}

func getBackgroundGene(g string) int {
	return getGeneInt(g, -2, 0, BASE_GENES_COUNT)
}

func getBackgroundGenePath(g string) string {
	gene := getBackgroundGene(g)
	return Gene(gene).toPath()
}

func getBaseGene(g string) int {
	return getGeneInt(g, -2, 0, BASE_GENES_COUNT)
}

func getBaseGeneAttribute(g string, configService *config.ConfigService) Attribute {
	gene := getBaseGene(g)
	return Attribute{
		TraitType: "Character",
		Value:     configService.Character[gene],
	}
}

func getBaseGenePath(g string) string {
	gene := getBaseGene(g)
	return Gene(gene).toPath()
}

func (g *Genome) name(configService *config.ConfigService, tokenId string) string {
	gStr := string(*g)
	gene := getBaseGene(gStr)
	return fmt.Sprintf("%v #%v", configService.Character[gene], tokenId)
}

func (g *Genome) description(configService *config.ConfigService, tokenId string) string {
	gStr := string(*g)
	gene := getBaseGene(gStr)
	return fmt.Sprintf("The %v named %v #%v is a citizen of the Polymorph Universe and has a unique genetic code! You can scramble your Polymorph at anytime.", configService.Type[gene], configService.Character[gene], tokenId)
}

func (g *Genome) genes() []string {
	gStr := string(*g)

	res := make([]string, 0, GENES_COUNT)

	res = append(res, getWeaponLeftGenePath(gStr))
	res = append(res, getWeaponRightGenePath(gStr))
	res = append(res, getHeadGenePath(gStr))
	res = append(res, getEyewearGenePath(gStr))
	res = append(res, getTorsoGenePath(gStr))
	res = append(res, getPantsGenePath(gStr))
	res = append(res, getShoesGenePath(gStr))
	res = append(res, getBaseGenePath(gStr))
	res = append(res, getBackgroundGenePath(gStr))

	return res
}

func (g *Genome) attributes(configService *config.ConfigService) []Attribute {
	gStr := string(*g)

	res := make([]Attribute, 0, GENES_COUNT)
	res = append(res, getBaseGeneAttribute(gStr, configService))
	res = append(res, getShoesGeneAttribute(gStr, configService))
	res = append(res, getPantsGeneAttribute(gStr, configService))
	res = append(res, getTorsoGeneAttribute(gStr, configService))
	res = append(res, getEyewearGeneAttribute(gStr, configService))
	res = append(res, getHeadGeneAttribute(gStr, configService))
	res = append(res, getWeaponLeftGeneAttribute(gStr, configService))
	res = append(res, getWeaponRightGeneAttribute(gStr, configService))
	return res
}

func (g *Genome) Metadata(tokenId string, configService *config.ConfigService) Metadata {
	var m Metadata
	genes := g.genes()

	m.Attributes = g.attributes(configService)
	m.Name = g.name(configService, tokenId)
	m.Description = g.description(configService, tokenId)
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
	Description string      `json:"description"`
	Name        string      `json:"name"`
	Image       string      `json:"image"`
	Attributes  []Attribute `json:"attributes"`
	ExternalUrl string      `json:"external_url"`
}
