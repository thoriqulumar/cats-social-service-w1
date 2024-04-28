package main

import (
	"github.com/thoriqulumar/cats-social-service-w1/configs"
)

func main() {
	// ...
	// Load the configuration file
	configs.LoadConfig()
	// ...

	initRouter()
}
