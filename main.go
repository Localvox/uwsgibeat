package main

import (
    "os"

	"github.com/elastic/beats/libbeat/beat"

    "github.com/Localvox/uwsgibeat/beater"
)

var Version = "1.0.0-beta1"
var Name = "uwsgibeat"

func main() {
	err := beat.Run(Name, Version, beater.New)
    if err != nil {
        os.Exit(1)
    }
}
