package consumer

import (
	"context"
	"fmt"
	"sync"

	"github.com/free5gc/nrf/internal/logger"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/openapi/nwdaf/EventsSubscription"
)

type nwdafService struct {
	consumer *Consumer

	eventSubscriptionMu     sync.RWMutex
	eventSubscriptionClient map[string]*EventsSubscription.APIClient
}

func (s *nwdafService) getEventSubscriptionClient(uri string) *EventsSubscription.APIClient {
	if uri == "" {
		return nil
	}
	s.eventSubscriptionMu.RLock()
	client, ok := s.eventSubscriptionClient[uri]
	if ok {
		s.eventSubscriptionMu.RUnlock()
		return client
	}

	configuration := EventsSubscription.NewConfiguration()
	configuration.SetBasePath(uri)
	client = EventsSubscription.NewAPIClient(configuration)

	s.eventSubscriptionMu.RUnlock()
	s.eventSubscriptionMu.Lock()
	defer s.eventSubscriptionMu.Unlock()

	s.eventSubscriptionClient[uri] = client
	return client
}

func (s *nwdafService) CreateEventSubscription(
	ctx context.Context,
	uri string,
) (string, error) {
	client := s.getEventSubscriptionClient(uri)
	if client == nil {
		return "", fmt.Errorf("Can't Get/New Client for url for [%+v]", uri)
	}
	s.eventSubscriptionMu.RLock()
	defer s.eventSubscriptionMu.RUnlock()

	callbackUri := fmt.Sprintf("%s://%s:%d/nwdaf-oam/callback",
		s.consumer.Config().Configuration.Sbi.Scheme,
		s.consumer.Config().Configuration.Sbi.BindingIPv4,
		s.consumer.Config().Configuration.Sbi.Port)

	subscriptionRequest := &EventsSubscription.CreateNWDAFEventsSubscriptionRequest{
		NnwdafEventsSubscription: &models.NnwdafEventsSubscription{
			EventSubscriptions: []models.NwdafEventsSubscriptionEventSubscription{
				{
					Event: models.NwdafEvent_NF_LOAD,
					NfTypes: []models.NrfNfManagementNfType{
						models.NrfNfManagementNfType_NRF,
					},
				},
			},
			NotificationURI: callbackUri,
		},
	}

	resp, err := client.NWDAFEventsSubscriptionsCollectionApi.CreateNWDAFEventsSubscription(ctx, subscriptionRequest)
	if err != nil {
		return "", fmt.Errorf("CreateNWDAFEventsSubscription failed: %w", err)
	}
	logger.ConsumerLog.Infoln("CreateNWDAFEventsSubscription location:", resp.Location)

	return resp.Location, nil
}

func (s *nwdafService) DeleteEventSubscription(
	ctx context.Context,
	uri string,
	subscriptionId string,
) error {
	client := s.getEventSubscriptionClient(uri)
	if client == nil {
		return fmt.Errorf("Can't Get/New Client for url for [%+v]", uri)
	}
	s.eventSubscriptionMu.RLock()
	defer s.eventSubscriptionMu.RUnlock()

	deleyRequest := &EventsSubscription.DeleteNWDAFEventsSubscriptionRequest{
		SubscriptionId: &subscriptionId,
	}

	resp, err := client.IndividualNWDAFEventsSubscriptionDocumentApi.DeleteNWDAFEventsSubscription(ctx, deleyRequest)
	if err != nil {
		return fmt.Errorf("DeleteNWDAFEventsSubscription failed: %w", err)
	}
	logger.ConsumerLog.Infoln("DeleteNWDAFEventsSubscription response:", resp)

	return nil
}
