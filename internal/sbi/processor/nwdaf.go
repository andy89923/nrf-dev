package processor

import (
	"context"

	"github.com/free5gc/nrf/internal/logger"
	"github.com/free5gc/openapi/models"
)

func (p *Processor) ReceiveNfLoadLevelAnalytics(notification *[]models.NnwdafEventsSubscriptionNotification) {
	if (notification == nil) || (len(*notification) == 0) {
		logger.ProcessorLog.Warnln("ReceiveNfLoadLevelAnalytics: notification is nil or empty")
		return
	}
	if len((*notification)[0].EventNotifications) == 0 {
		logger.ProcessorLog.Warnln("ReceiveNfLoadLevelAnalytics: EventNotifications is nil or empty")
		return
	}
	eventNotification := (*notification)[0].EventNotifications[0]
	if len(eventNotification.NfLoadLevelInfos) == 0 {
		logger.ProcessorLog.Warnln("ReceiveNfLoadLevelAnalytics: NfLoadLevelInfos is nil or empty")
		return
	}

	nfLoadLevelInfo := eventNotification.NfLoadLevelInfos[0]
	if nfLoadLevelInfo.Confidence == 0 {
		// Confidence is 0, ignore this notification (This notification is not prediction)
		logger.ProcessorLog.Warnln("ReceiveNfLoadLevelAnalytics: Confidence is 0, ignore this notification")
		return
	}

	logger.ProcessorLog.Warnf("ReceiveNfLoadLevelAnalytics: NfLoadLevelInfo: %+v", nfLoadLevelInfo)
	logger.ProcessorLog.Warnf("LoadLevel Peak: %+v", nfLoadLevelInfo.NfLoadLevelpeak)

	// TODO: Process the nfLoadLevelInfo
	// If the NfLoadLevelPeak is greater than the threshold, adjust the token period
}

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
