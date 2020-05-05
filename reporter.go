package zapreporter

import (
	"github.com/setare/services"
	"go.uber.org/zap"
)

type reporter struct {
	logger *zap.Logger
}

// NewReporter returns a services.Reporter that uses the zap logging library to
// output the process actions.
func NewReporter(logger *zap.Logger) services.Reporter {
	return &reporter{
		logger: logger,
	}
}

func (reporter *reporter) BeforeStart(service services.Service) {
	reporter.logger.Info("Starting", zap.String("service", service.Name()))
}

func (reporter *reporter) AfterStart(service services.Service, err error) {
	if err != nil {
		reporter.logger.Error("Start failed", zap.String("service", service.Name()), zap.Error(err))
		return
	}
	reporter.logger.Info("Started", zap.String("service", service.Name()))
}

func (reporter *reporter) BeforeStop(service services.Service) {
	reporter.logger.Info("Stopping", zap.String("service", service.Name()))
}

func (reporter *reporter) AfterStop(service services.Service, err error) {
	if err != nil {
		reporter.logger.Error("Stop failed", zap.String("service", service.Name()), zap.Error(err))
		return
	}
	reporter.logger.Info("Stopped", zap.String("service", service.Name()))
}

func (reporter *reporter) BeforeLoad(configurable services.Configurable) {
	serviceName := "unknown"
	if srv, ok := configurable.(services.Service); ok {
		serviceName = srv.Name()
	}
	reporter.logger.Info("Loading configuration", zap.String("service", serviceName))
}

func (reporter *reporter) AfterLoad(configurable services.Configurable, err error) {
	serviceName := "unknown"
	if srv, ok := configurable.(services.Service); ok {
		serviceName = srv.Name()
	}
	if err != nil {
		reporter.logger.Error("Load configuration failed", zap.String("service", serviceName), zap.Error(err))
		return
	}
	reporter.logger.Info("Configuration loaded", zap.String("service", serviceName))
}
