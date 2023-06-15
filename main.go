package main

import (
	"log"

	"github.com/eyedeekay/i2p-stats/site"
)

func main() {
	if site, err := site.NewStatsSite("weather"); err != nil {
		log.Fatal(err)
	} else {
		if err := site.OutputPages(); err != nil {
			log.Fatal(err)
		}
		if err := site.GenerateIndexPages(); err != nil {
			log.Fatal(err)
		}
		if err := site.OutputHomePage(); err != nil {
			log.Fatal(err)
		}
	}
}
