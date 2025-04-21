package processor

import (
	"sync"

	"github.com/free5gc/nrf/internal/sbi/consumer"
	"github.com/free5gc/nrf/pkg/app"
)

type ProcessorNrf interface {
	app.App
	Consumer() *consumer.Consumer
}

type Processor struct {
	ProcessorNrf

	sync.Mutex

	NwdafUri            string
	NwdafSubscriptionId string
	TokenExpiration     int32 // milliseconds
}

func NewProcessor(nrf ProcessorNrf) (*Processor, error) {
	p := &Processor{
		ProcessorNrf:        nrf,
		TokenExpiration:     1000, // default expiration time
		NwdafUri:            "",
		NwdafSubscriptionId: "",
	}
	return p, nil
}
