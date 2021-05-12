package handlers

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/render"
	"github.com/polymorph-metadata/contracts"
	log "github.com/sirupsen/logrus"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const POLYMORPH_IMAGE_URL string = "https://storage.googleapis.com/polymorph-images/"
const EXTERNAL_URL string = "https://kek.dao/"
const BASE_GENES_COUNT int = 6
const HEAD_GENES_COUNT int = 4
const WEAPON_GENES_COUNT int = 3

type Metadata struct {
	Description string   `json:"description"`
	Name        string   `json:"name"`
	Image       string   `json:"image"`
	Attributes  []string `json:"attributes"`
	ExternalUrl string   `json:"external_url"`
}

type Genome string

func getGene(g string, start, end, count int) string {
	genomeLen := len(g)
	geneStr := g[genomeLen+start : genomeLen+end]
	gene, _ := strconv.Atoi(geneStr)
	gene = gene % count
	if gene < 10 {
		geneStr = fmt.Sprintf("0%s", strconv.Itoa(gene))
	}

	return geneStr
}

func getWeaponGene(g string) string {
	return getGene(g, -6, -4, WEAPON_GENES_COUNT)
}

func getHeadGene(g string) string {
	return getGene(g, -4, -2, HEAD_GENES_COUNT)
}

func getBaseGene(g string) string {
	return getGene(g, -2, 0, BASE_GENES_COUNT)
}

func (g *Genome) genomeToImagePath() string {
	gStr := string(*g)
	b := strings.Builder{}

	b.WriteString(POLYMORPH_IMAGE_URL) // Start with base url

	weaponGene := getWeaponGene(gStr) // add weapon gene
	b.WriteString(weaponGene)

	headGene := getHeadGene(gStr) // add weapon gene
	b.WriteString(headGene)

	baseGenome := getBaseGene(gStr) // add base gene
	b.WriteString(baseGenome)

	b.WriteString(".png") // Finish with png extension
	return b.String()
}

func (g *Genome) Metadata(tokenId string) Metadata {
	var m Metadata
	m.Name = fmt.Sprintf("Polymorph with Id %v", tokenId)
	m.Description = fmt.Sprintf("Polymorph with Id %v and Genome %v", tokenId, string(*g))
	m.Image = g.genomeToImagePath()
	m.Attributes = make([]string, 0, 0)
	m.ExternalUrl = fmt.Sprintf("%s%s", EXTERNAL_URL, tokenId)
	return m
}

func TokenMetadata(w http.ResponseWriter, r *http.Request) {
	client, err := ethclient.Dial(os.Getenv("NODE_URL"))
	if err != nil {
		render.Status(r, 500)
		log.Fatal(err)
	}

	address := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := contracts.NewPolymorph(address, client)
	if err != nil {
		render.Status(r, 500)
		log.Fatal(err)
	}

	tokenId := r.URL.Query().Get("id")

	iTokenId, err := strconv.Atoi(tokenId)
	if err != nil {
		render.Status(r, 500)
		log.Fatal(err)
	}

	genomeInt, err := instance.GeneOf(nil, big.NewInt(int64(iTokenId)))
	if err != nil {
		render.Status(r, 500)
		log.Fatal(err)
	}

	var g Genome
	g = Genome(genomeInt.String())
	render.JSON(w, r, (&g).Metadata(tokenId))
}
