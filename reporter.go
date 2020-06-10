package zapreporter

import (
	"os"

	"github.com/setare/services"
	"go.uber.org/zap"
)

type reporter struct {
	logger *zap.Logger
}

// NewReporter returns a services.Reporter that uses the zap logging library to
// output the process actions.
func NewReporter(logger *zap.Logger, options ...zap.Option) services.RetrierReporter {
	l := logger
	if len(options) > 0 {
		l = logger.WithOptions(zap.AddCallerSkip(2))
	}
	return &reporter{
		logger: l,
	}
}

func (reporter *reporter) BeforeStart(service services.Service) {
	if service == nil {
		return
	}
	reporter.logger.Info("Starting", zap.String("service", service.Name()))
}

func (reporter *reporter) AfterStart(service services.Service, err error) {
	if service == nil {
		return
	}
	if err != nil {
		reporter.logger.Error("Start failed", zap.String("service", service.Name()), zap.Error(err))
		return
	}
	reporter.logger.Info("Started", zap.String("service", service.Name()))
}

func (reporter *reporter) BeforeStop(service services.Service) {
	if service == nil {
		reporter.logger.Info("Stopping services")
		return
	}
	reporter.logger.Info("Stopping", zap.String("service", service.Name()))
}

func (reporter *reporter) AfterStop(service services.Service, err error) {
	if service == nil {
		return
	}
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

func (reporter *reporter) SignalReceived(sig os.Signal) {
	reporter.logger.Info("signal received", zap.String("signal", sig.String()))
}

func (reporter *reporter) BeforeRetry(service services.Service, try int) {
	serviceName := service.Name()
	reporter.logger.Info("Retrying service", zap.String("service", serviceName), zap.Int("try", try))
}

func (reporter *reporter) AfterGiveUp(service services.Service, try int, err error) {
	serviceName := service.Name()
	reporter.logger.Info("Giving up", zap.String("service", serviceName), zap.Int("try", try), zap.Error(err))
}
