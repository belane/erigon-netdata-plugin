package erigon

import (
	"errors"
	"time"

	"github.com/netdata/go.d.plugin/agent/module"
	"github.com/netdata/go.d.plugin/pkg/prometheus"
	"github.com/netdata/go.d.plugin/pkg/web"
)

func init() {
	creator := module.Creator{
		Create: func() module.Module { return New() },
	}

	module.Register("erigon", creator)
}

func New() *Erigon {
	config := Config{
		HTTP: web.HTTP{
			Request: web.Request{
				URL: "http://127.0.0.1:6060/debug/metrics/prometheus",
			},
			Client: web.Client{
				Timeout: web.Duration{Duration: time.Second},
			},
		},
	}

	return &Erigon{
		Config: config,
		charts: charts.Copy(),
	}
}

type (
	Config struct {
		web.HTTP `yaml:",inline"`
	}

	Erigon struct {
		module.Base
		Config `yaml:",inline"`

		prom   prometheus.Prometheus
		charts *Charts
	}
)

func (v Erigon) validateConfig() error {
	if v.URL == "" {
		return errors.New("URL is not set")
	}
	return nil
}

func (v *Erigon) initClient() error {
	client, err := web.NewHTTPClient(v.Client)
	if err != nil {
		return err
	}

	v.prom = prometheus.New(client, v.Request)
	return nil
}

func (v *Erigon) Init() bool {
	if err := v.validateConfig(); err != nil {
		v.Errorf("error on validating config: %v", err)
		return false
	}
	if err := v.initClient(); err != nil {
		v.Errorf("error on initializing client: %v", err)
		return false
	}
	return true
}

func (v *Erigon) Check() bool {
	return len(v.Collect()) > 0
}

func (v *Erigon) Charts() *Charts {
	return v.charts
}

func (v *Erigon) Collect() map[string]int64 {
	mx, err := v.collect()
	if err != nil {
		v.Error(err)
	}

	if len(mx) == 0 {
		return nil
	}
	return mx
}

func (Erigon) Cleanup() {}
