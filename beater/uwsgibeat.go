package beater

import (
    "fmt"
    "time"
    "net/url"

    "github.com/elastic/beats/libbeat/beat"
    "github.com/elastic/beats/libbeat/common"
    "github.com/elastic/beats/libbeat/logp"
    "github.com/elastic/beats/libbeat/publisher"

    "github.com/Localvox/uwsgibeat/config"
    "github.com/Localvox/uwsgibeat/parser"
)

const selector = "uwsgibeat"

type Uwsgibeat struct {
    // UbConfig holds configurations of Uwsgibeat parsed by libbeat.
    UbConfig config.Config
    done   chan struct{}
    events publisher.Client
    url    *url.URL
    period time.Duration
}


// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
    config := config.DefaultConfig
    if err := cfg.Unpack(&config); err != nil {
        return nil, fmt.Errorf("Error reading config file: %v", err)
    }

    ub := &Uwsgibeat{
        done: make(chan struct{}),
        UbConfig: config,
    }

    var u string
    var err error

    if ub.UbConfig.Input.URL != "" {
        u = ub.UbConfig.Input.URL
    } else {
        u = "127.0.0.1:1717"
    }
    ub.url, err = url.Parse(u)
    if err != nil {
        return nil, fmt.Errorf("Invalid uWSGI stats server address: %v", err)
    }

    if ub.UbConfig.Input.Period != 0 {
        ub.period = ub.UbConfig.Input.Period
    } else {
        ub.period = 1 * time.Second
    }

    logp.Debug(selector, "Init uwsgibeat")
    logp.Debug(selector, "Watch %v", ub.url)
    logp.Debug(selector, "Period %v", ub.period)

    return ub, nil
}

func (ub *Uwsgibeat) Run(b *beat.Beat) error {
    logp.Info("uwsgibeat is running! Hit CTRL-C to stop it.")

    ub.events = b.Publisher.Connect()
    p := parser.NewStatsParser()
    ticker := time.NewTicker(ub.period)
    defer ticker.Stop()

    for {
        select {
        case <-ub.done:
            return nil
        case <-ticker.C:
        }

        start := time.Now()

        s, err := p.Parse(*ub.url)
        if err != nil {
            logp.Err("Fail to read uWSGI stats: %v", err)
            goto GotoNext
        }
        ub.events.PublishEvent(common.MapStr{
            "@timestamp": common.Time(time.Now()),
            "type":        "uwsgi",
            "uwsgi":       s,
        })

    GotoNext:
        end := time.Now()
        duration := end.Sub(start)
        if duration.Nanoseconds() > ub.period.Nanoseconds() {
            logp.Warn("Ignoring tick(s) due to processing taking longer than one period")
        }
    }
}

func (ub *Uwsgibeat) Stop() {
    logp.Debug(selector, "Stop uwsgibeat")
    ub.events.Close()
    close(ub.done)
}
