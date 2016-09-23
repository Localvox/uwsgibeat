// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type UWSGIConfig struct {
    Period  time.Duration `config:"period"`

    URL     string        `config:"url"`
}

type Config struct {
    Input UWSGIConfig
}

var DefaultConfig = Config{
    UWSGIConfig {
        Period: 10 * time.Second,

        URL: "tcp://127.0.0.1:1717",
    },
}
