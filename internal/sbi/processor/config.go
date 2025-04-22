package processor

import (
	"fmt"

	"github.com/free5gc/nrf/internal/logger"
	"github.com/free5gc/nrf/pkg/configurations"
)

func (p *Processor) ConfigUpdate(config *configurations.DynamicConfig) error {
	if config == nil {
		return fmt.Errorf("config is nil")
	}

	// Update OAuth2 configuration
	if config.Sbi.OAuth2.Enable {
		p.Lock()
		defer p.Unlock()

		p.Config().Configuration.Sbi.OAuth = config.Sbi.OAuth2.Enable
		p.TokenExpiration = config.Sbi.OAuth2.Period

		logger.ProcessorLog.Infof("OAuth2 Settings: Enable: %v, Period: %d",
			p.Config().Configuration.Sbi.OAuth, p.TokenExpiration)
	}
	return nil
}
