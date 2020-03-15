package blackbox_exporter

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/blackbox_exporter/config"
	"github.com/prometheus/blackbox_exporter/prober"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/promlog"
	"net/http"
	"net/url"
	"sync"
	"time"
)


type result struct {
	id          int64
	moduleName  string
	target      string
	debugOutput string
	success     bool
}

// resultHistory contains two history slices: `results` contains most recent `maxResults` results.
// After they expire out of `results`, failures will be saved in `preservedFailedResults`. This
// ensures that we are always able to see debug information about recent failures.
type resultHistory struct {
	mu                     sync.Mutex
	nextId                 int64
	results                []*result
	preservedFailedResults []*result
	maxResults             uint
}



func HttpProbe(url *url.URL, targetHack string) {

	//sc := &config.SafeConfig{
	//	C: &config.Config{},
	//}

	// w
	// r := http.Request{Method: "GET", URL: url}
	//conf := sc.C
	//promlogConfig := &promlog.Config{}
	//logger := promlog.New(promlogConfig)
	//rh := &resultHistory{maxResults: 100}
	//
	//prober.ProbeHTTP()



	r := &http.Request{Method: "GET", URL: url}
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(120))
	defer cancel()
	r = r.WithContext(ctx)

	params := r.URL.Query()
	target := params.Get("target")
	fmt.Println("scratch target: ", target)
	if target == "" {
		target = targetHack
	}

	sc := &config.SafeConfig{
		C: &config.Config{},
	}
	conf := sc.C

	registry := prometheus.NewRegistry()

	promlogConfig := &promlog.Config{}
	promlogConfig.Level = &promlog.AllowedLevel{}
	promlogConfig.Level.Set("debug")
	promlogConfig.Format = &promlog.AllowedFormat{}

	logger := promlog.New(promlogConfig)
	moduleName := "http_2x"
	sl := newScrapeLogger(logger, moduleName, target)

	prober.ProbeHTTP(ctx, target, conf.Modules[moduleName], registry, sl)

}

type scrapeLogger struct {
	next         log.Logger
	buffer       bytes.Buffer
	bufferLogger log.Logger
}

func (sl scrapeLogger) Log(keyvals ...interface{}) error {
	sl.bufferLogger.Log(keyvals...)
	kvs := make([]interface{}, len(keyvals))
	copy(kvs, keyvals)
	// Switch level to debug for application output.
	for i := 0; i < len(kvs); i += 2 {
		if kvs[i] == level.Key() {
			kvs[i+1] = level.DebugValue()
		}
	}
	return sl.next.Log(kvs...)
}

func newScrapeLogger(logger log.Logger, module string, target string) *scrapeLogger {
	logger = log.With(logger, "module", module, "target", target)
	sl := &scrapeLogger{
		next:   logger,
		buffer: bytes.Buffer{},
	}
	bl := log.NewLogfmtLogger(&sl.buffer)
	sl.bufferLogger = log.With(bl, "ts", log.DefaultTimestampUTC, "caller", log.Caller(6), "module", module, "target", target)
	return sl
}


