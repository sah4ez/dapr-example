package config

import (
	"io"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"go.uber.org/automaxprocs/maxprocs"
)

const formatJSON = "json"

type ServiceConfig struct {
	LogLevel     string `envconfig:"LOGGER_LEVEL" default:"trace"`
	LogFormat    string `envconfig:"LOGGER_FORMAT" default:"console"`
	ReportCaller bool   `envconfig:"LOG_REPORT_CALLER" default:"false"`
	LogStack     bool   `envconfig:"LOG_STACK" default:"false"`

	Bind string `envconfig:"LISTEN" default:":9000"`

	BindMetrics string `envconfig:"LISTEN_METRICS" default:":9090"`

	AllowOrigins string `envconfig:"ALLOW_ORIGINS" default:"http://localhost:9000"`

	AllowHeaders string `envconfig:"ALLOW_HEADERS" default:"Content-Type,Authorization,User-Agent,Accept,Referer,Cookie,X-Request-Id,X-Log-Level"`

	EnabledPPROF bool   `envconfig:"ENABLED_PPROF" default:"true"`
	BindPPROF    string `envconfig:"BIND_PPROF" default:":6060"`

	once *sync.Once
}

var service *ServiceConfig

func Service() ServiceConfig {

	if service != nil {
		return *service
	}
	service = &ServiceConfig{once: new(sync.Once)}
	if err := envconfig.Process("", service); err != nil {
		panic(err)
	}
	return service.log()
}

func (cfg ServiceConfig) Logger() (logger zerolog.Logger) {

	level := zerolog.InfoLevel
	if newLevel, err := zerolog.ParseLevel(cfg.LogLevel); err == nil {
		level = newLevel
	}
	var out io.Writer = os.Stdout
	if cfg.LogFormat != formatJSON {
		out = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.StampMicro}
	}
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	ctxLog := zerolog.New(out).Level(level).With().Timestamp().Stack()
	if cfg.ReportCaller {
		ctxLog = ctxLog.Caller()
	}
	return ctxLog.Logger()
}

func (cfg *ServiceConfig) log() ServiceConfig {
	cfg.once.Do(
		func() {
			log.Logger = cfg.Logger()
			_, _ = maxprocs.Set(maxprocs.Logger(func(format string, v ...interface{}) {
				log.Info().Int("numCPU", runtime.NumCPU()).Msgf(format, v...)
			}))
			log.Info().
				Str("config", "service").
				Interface("values", cfg).
				Send()
		})
	return *cfg
}
