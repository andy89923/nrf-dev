package consumer

import (
	"github.com/free5gc/nrf/pkg/app"
	"github.com/free5gc/openapi/nrf/NFManagement"
	"github.com/free5gc/openapi/nwdaf/EventsSubscription"
)

type ConsumerNrf interface {
	app.App
}

type Consumer struct {
	ConsumerNrf

	*nnrfService
	*nwdafService
}

func NewConsumer(nrf ConsumerNrf) (*Consumer, error) {
	c := &Consumer{
		ConsumerNrf: nrf,
	}

	c.nnrfService = &nnrfService{
		consumer:        c,
		nfMngmntClients: make(map[string]*NFManagement.APIClient),
	}
	c.nwdafService = &nwdafService{
		consumer:                c,
		eventSubscriptionClient: make(map[string]*EventsSubscription.APIClient),
	}
	return c, nil
}
