package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/eyedeekay/i2p-stats/site"
)

func main() {
	if statsite, err := site.NewStatsSite("weather"); err != nil {
		log.Fatal(err)
	} else {
		if err := statsite.OutputPages(); err != nil {
			log.Fatal(err)
		}
		if err := statsite.GenerateIndexPages(); err != nil {
			log.Fatal(err)
		}

		if edgarIsInstalled() {
			statsite.OutputMarkdownHomePage()
			cmd := exec.Command("edgar", os.Args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			//TODO set EDGAR_RECURSIVE=true in the cmd environment
			cmd.Dir = statsite.StatsDirectory
			cmd.Run()
		} else {
			if err := statsite.OutputHomePage(); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func edgarIsInstalled() bool {
	_, err := exec.LookPath("edgar")
	if err != nil {
		gopath := os.Getenv("GOPATH")
		if gopath != "" {
			binPath := filepath.Join(gopath, "bin", "edgar")
			_, err = os.Stat(binPath)
			if err == nil {
				return true
			}
		}
		return false
	}
	return true
}
