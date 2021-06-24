package processor

import (
	"log"
	"math/big"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/polymorph-metadata/app/config"
	"github.com/polymorph-metadata/app/contracts"
	"github.com/polymorph-metadata/app/domain/metadata"
	"github.com/polymorph-metadata/app/interface/dlt/ethereum"
)

type ProcessRequest struct {
	tokenId   string
	responseC chan *metadata.Metadata
	errorC    chan error
}

type Processor struct {
	requestChannel chan *ProcessRequest
	contract       *contracts.Polymorph
	configService  *config.ConfigService
	queueLength    int
}

func (p *Processor) Start() {
	go p.handle()
}

func (p *Processor) process(wg *sync.WaitGroup, r *ProcessRequest) {
	iTokenId, err := strconv.Atoi(r.tokenId)
	if err != nil {
		r.errorC <- err
		return
	}

	genomeInt, err := p.contract.GeneOf(nil, big.NewInt(int64(iTokenId)))
	if err != nil {
		r.errorC <- err
		return
	}

	g := metadata.Genome(genomeInt.String())
	metadata := g.Metadata(r.tokenId, p.configService)

	r.responseC <- &metadata

	wg.Done()
}

func (p *Processor) handle() {
	var wg sync.WaitGroup
	for {
		for i := 1; i <= p.queueLength; i++ {
			r := <-p.requestChannel
			wg.Add(1)
			go p.process(&wg, r)
		}

		wg.Wait()

	}
}

func (p *Processor) HandleProcessRequest(tokenId string, responseC chan *metadata.Metadata, errorC chan error) {
	p.requestChannel <- &ProcessRequest{
		tokenId,
		responseC,
		errorC,
	}
}

func NewProcessor(queue int, ethClient *ethereum.EthereumClient, address string, configService *config.ConfigService) *Processor {

	instance, err := contracts.NewPolymorph(common.HexToAddress(address), ethClient.Client)
	if err != nil {
		log.Fatalln(err)
	}

	return &Processor{
		requestChannel: make(chan *ProcessRequest, queue),
		contract:       instance,
		configService:  configService,
		queueLength:    queue,
	}
}
