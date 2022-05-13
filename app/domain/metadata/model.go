package metadata

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/polymorph-metadata/app/config"
	"github.com/polymorph-metadata/structs"
)

const POLYMORPH_IMAGE_URL string = "https://storage.googleapis.com/polymorph-images/"
const EXTERNAL_URL string = "https://universe.xyz/polymorphs/"
const GENES_COUNT = 13
const BACKGROUND_GENE_COUNT int = 5
const HAIR_LEFT int = 33
const HAIR_RIGHT int = 33
const EAR_LEFT int = 33
const EAR_RIGHT int = 33
const BEARD_TOP_LEFT int = 33
const BEARD_TOP_RIGHT int = 33
const LIPS_LEFT int = 33
const LIPS_RIGHT int = 33
const BEARD_BOTTOM_LEFT int = 33
const BEARD_BOTTOM_RIGHT int = 33
const EYE_RIGHT int = 33
const EYE_LEFT int =33

type Genome string
type Gene int
type StringAttribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type IntegerAttribute struct {
	TraitType string `json:"trait_type"`
	Value     int    `json:"value"`
}

type FloatAttribute struct {
	TraitType   string  `json:"trait_type"`
	Value       float64 `json:"value"`
	DisplayType string  `json:"display_type"`
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


func getEyeRightGene(g string) int {
	return getGeneInt(g, -28, -26, LIPS_RIGHT)
}

func getEyeRightGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getEyeRightGene(g)
	return StringAttribute{
		TraitType: "Eye Right",
		Value:     configService.Traits[gene],
	}
}

func getEyeRightPath(g string) string {
	gene := getEyeRightGene(g)
	return Gene(gene).toPath()
}

func getEyeLeftGene(g string) int {
	return getGeneInt(g, -26, -24, LIPS_LEFT)
}

func getEyeLeftGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getEyeLeftGene(g)
	return StringAttribute{
		TraitType: "Eye Left",
		Value:     configService.Traits[gene],
	}
}

func getEyeLeftPath(g string) string {
	gene := getEyeLeftGene(g)
	return Gene(gene).toPath()
}


func getBeardBottomRightGene(g string) int {
	return getGeneInt(g, -24, -22, BEARD_BOTTOM_RIGHT)
}

func getBeardBottomRightGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getBeardBottomRightGene(g)
	return StringAttribute{
		TraitType: "Beard Bottom Right",
		Value:     configService.Traits[gene],
	}
}

func getBeardBottomRightPath(g string) string {
	gene := getBeardBottomRightGene(g)
	return Gene(gene).toPath()
}

func getBeardBottomLeftGene(g string) int {
	return getGeneInt(g, -22, -20, BEARD_BOTTOM_LEFT)
}

func getBeardBottomLeftGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getBeardBottomLeftGene(g)
	return StringAttribute{
		TraitType: "Beard Bottom Left",
		Value:     configService.Traits[gene],
	}
}

func getBeardBottomLeftPath(g string) string {
	gene := getBeardBottomLeftGene(g)
	return Gene(gene).toPath()
}

func getLipsRightGene(g string) int {
	return getGeneInt(g, -20, -18, LIPS_RIGHT)
}

func getLipsRightGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getLipsRightGene(g)
	return StringAttribute{
		TraitType: "Lips Right",
		Value:     configService.Traits[gene],
	}
}

func getLipsRightPath(g string) string {
	gene := getLipsRightGene(g)
	return Gene(gene).toPath()
}

func getLipsLeftGene(g string) int {
	return getGeneInt(g, -18, -16, LIPS_LEFT)
}

func getLipsLeftGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getLipsLeftGene(g)
	return StringAttribute{
		TraitType: "Lips Left",
		Value:     configService.Traits[gene],
	}
}

func getLipsLeftPath(g string) string {
	gene := getLipsLeftGene(g)
	return Gene(gene).toPath()
}


func getBeardTopRightGene(g string) int {
	return getGeneInt(g, -16, -14, BEARD_TOP_RIGHT)
}

func getBeardTopRightGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getBeardTopRightGene(g)
	return StringAttribute{
		TraitType: "Beard Top Right",
		Value:     configService.Traits[gene],
	}
}

func getBeardTopRightPath(g string) string {
	gene := getBeardTopRightGene(g)
	return Gene(gene).toPath()
}

func getBeardTopLeftGene(g string) int {
	return getGeneInt(g, -14, -12, BEARD_BOTTOM_LEFT)
}

func getBeardTopLeftGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getBeardTopLeftGene(g)
	return StringAttribute{
		TraitType: "Beard Top Left",
		Value:     configService.Traits[gene],
	}
}

func getBeardTopLeftPath(g string) string {
	gene := getBeardTopLeftGene(g)
	return Gene(gene).toPath()
}


func getEarsRightGene(g string) int {
	return getGeneInt(g, -12, -10, EAR_RIGHT)
}

func getEarsRightGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getEarsRightGene(g)
	return StringAttribute{
		TraitType: "Ear Right",
		Value:     configService.Traits[gene],
	}
}

func getEarsRightPath(g string) string {
	gene := getEarsRightGene(g)
	return Gene(gene).toPath()
}

func getEarsLeftGene(g string) int {
	return getGeneInt(g, -10, -8, EAR_LEFT)
}

func getEarsLeftGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getEarsLeftGene(g)
	return StringAttribute{
		TraitType: "Ear Left",
		Value:     configService.Traits[gene],
	}
}

func getEarsLeftPath(g string) string {
	gene := getEarsLeftGene(g)
	return Gene(gene).toPath()
}


func getHairRightGene(g string) int {
	return getGeneInt(g, -6, -4, HAIR_RIGHT)
}

func getHairRightGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getHairRightGene(g)
	return StringAttribute{
		TraitType: "Hair Right",
		Value:     configService.Traits[gene],
	}
}

func getHairRightPath(g string) string {
	gene := getHairRightGene(g)
	return Gene(gene).toPath()
}

func getHairLeftGene(g string) int {
	return getGeneInt(g, -4, -2, HAIR_LEFT)
}

func getHairLeftGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getHairLeftGene(g)
	return StringAttribute{
		TraitType: "Hair Left",
		Value:     configService.Traits[gene],
	}
}

func getHairLeftPath(g string) string {
	gene := getHairLeftGene(g)
	return Gene(gene).toPath()
}


func getBackgroundGene(g string) int {
	return getGeneInt(g, -2, 0, BACKGROUND_GENE_COUNT)
}

func getBackgroundGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getBackgroundGene(g)
	return StringAttribute{
		TraitType: "Background",
		Value:     configService.Background[gene],
	}
}

func getBackgroundGenePath(g string) string {
	gene := getBackgroundGene(g)
	return Gene(gene).toPath()
}


func (g *Genome) name(configService *config.ConfigService, tokenId string) string {
	return fmt.Sprintf("Polymorphic Face #%v", tokenId)
}

func (g *Genome) description(configService *config.ConfigService, tokenId string) string {
	return fmt.Sprintf("Polymorphic Face Description")
}


// const BASE_GENES_COUNT int = 11
// const SHOES_GENES_COUNT int = 25
// const PANTS_GENES_COUNT int = 33
// const TORSO_GENES_COUNT int = 34
// const EYEWEAR_GENES_COUNT int = 13
// const HEAD_GENES_COUNT int = 31
// const WEAPON_RIGHT_GENES_COUNT int = 32
// const WEAPON_LEFT_GENES_COUNT int = 32


// const HAIR_LEFT int = 33
// const HAIR_RIGHT int = 33
// const EAR_LEFT int = 33
// const EAR_RIGHT int = 33
// const BEARD_TOP_LEFT int = 33
// const BEARD_TOP_RIGHT int = 33
// const LIPS_LEFT int = 33
// const LIPS_RIGHT int = 33
// const BEARD_BOTTOM_LEFT int = 33
// const BEARD_BOTTOM_RIGHT int = 33
// const EYE_RIGHT int = 33
// const EYE_LEFT int =33

func (g *Genome) genes() []string {
	gStr := string(*g)

	res := make([]string, 0, GENES_COUNT)

	res = append(res, getEyeRightPath(gStr))
	res = append(res, getEyeLeftPath(gStr))
	res = append(res, getBeardBottomRightPath(gStr))
	res = append(res, getBeardBottomLeftPath(gStr))
	res = append(res, getLipsRightPath(gStr))
	res = append(res, getLipsLeftPath(gStr))
	res = append(res, getBeardTopRightPath(gStr))
	res = append(res, getBeardTopLeftPath(gStr))
	res = append(res, getEarsRightPath(gStr))
	res = append(res, getEarsLeftPath(gStr))
	res = append(res, getHairRightPath(gStr))
	res = append(res, getHairLeftPath(gStr))
	res = append(res, getBackgroundGenePath(gStr))

	return res
}

func getRarityScoreAttribute(rarity float64) FloatAttribute {
	return FloatAttribute{
		TraitType:   "Rarity Score",
		DisplayType: "number",
		Value:       math.Round(rarity*100) / 100,
	}
}

func getRankAttribute(rank int) IntegerAttribute {
	return IntegerAttribute{
		TraitType: "Rank",
		Value:     rank,
	}
}

func (g *Genome) attributes(configService *config.ConfigService, rarityResponse structs.RarityServiceResponse) []interface{} {
	gStr := string(*g)

	res := []interface{}{}

	res = append(res, getEyeRightGeneAttribute(gStr, configService))
	res = append(res, getEyeLeftGeneAttribute(gStr, configService))
	res = append(res, getBeardBottomRightGeneAttribute(gStr, configService))
	res = append(res, getBeardBottomLeftGeneAttribute(gStr, configService))
	res = append(res, getLipsRightGeneAttribute(gStr, configService))
	res = append(res, getLipsLeftGeneAttribute(gStr, configService))
	res = append(res, getBeardTopRightGeneAttribute(gStr, configService))
	res = append(res, getBeardTopLeftGeneAttribute(gStr, configService))
	res = append(res, getEarsRightGeneAttribute(gStr, configService))
	res = append(res, getEarsLeftGeneAttribute(gStr, configService))
	res = append(res, getHairRightGeneAttribute(gStr, configService))
	res = append(res, getHairLeftGeneAttribute(gStr, configService))
	res = append(res, getBackgroundGeneAttribute(gStr, configService))
	res = append(res, getRarityScoreAttribute(rarityResponse.RarityScore))
	res = append(res, getRankAttribute(rarityResponse.Rank))

	return res
}

func (g *Genome) Metadata(tokenId string, configService *config.ConfigService, rarityResponse structs.RarityServiceResponse) Metadata {
	var m Metadata
	genes := g.genes()

	m.Attributes = g.attributes(configService, rarityResponse)
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

	// imageExists := imageExists(imageURL)

	// if !imageExists {
	// 	generateAndSaveImage(genes)
	// }

	m.Image = imageURL
	return m
}

type Metadata struct {
	Description string      `json:"description"`
	Name        string      `json:"name"`
	Image       string      `json:"image"`
	Attributes  interface{} `json:"attributes"`
	ExternalUrl string      `json:"external_url"`
}
