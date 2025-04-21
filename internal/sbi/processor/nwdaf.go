package processor

import (
	"context"

	"github.com/free5gc/nrf/internal/logger"
)

func (p *Processor) subscribeNfLoadLevelAnalytics(
	ctx context.Context,
	uri string,
) {
	p.Lock()
	defer p.Unlock()

	logger.NfmLog.Infof("CreateEventSubscription to %s", uri)

	subscriptionId, err := p.Consumer().CreateEventSubscription(ctx, uri)
	if err != nil {
		logger.NfmLog.Errorf("CreateEventSubscription failed: %+v", err)
		return
	}

	logger.NfmLog.Infof("CreateEventSubscription success: %s", subscriptionId)
	p.NwdafSubscriptionId = subscriptionId
	p.NwdafUri = uri
}

func (p *Processor) deleteSubscriptions(
	ctx context.Context,
	uri string,
) {
	p.Lock()
	defer p.Unlock()

	if p.NwdafSubscriptionId == "" {
		logger.NfmLog.Warn("NwdafSubscriptionId is empty")
		return
	}

	err := p.Consumer().DeleteEventSubscription(ctx, uri, p.NwdafSubscriptionId)
	if err != nil {
		logger.NfmLog.Errorf("DeleteEventSubscription failed: %+v", err)
		return
	}
	logger.NfmLog.Infof("DeleteEventSubscription success: %s", p.NwdafSubscriptionId)
	p.NwdafSubscriptionId = ""
}
