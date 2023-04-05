package main

import (
	_ "github.com/lib/pq"
	"myapiproject/internal/app"
)

const configsDir = "configs"

func main() {
	app.Run(configsDir)
}
