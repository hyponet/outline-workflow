package main

import (
	"github.com/hyponet/outline-workflow/cmd/apps"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if err := apps.RootCmd.Execute(); err != nil {
		panic(err)
	}
}
