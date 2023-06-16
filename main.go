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
		if edgarIsInstalled() {
			if err := statsite.OutputMarkdownPages(); err != nil {
				log.Fatal(err)
			}
			if err := statsite.GenerateMarkdownIndexPages(); err != nil {
				log.Fatal(err)
			}
			if err := statsite.OutputMarkdownHomePage(); err != nil {
				log.Fatal(err)
			}
			cmd := exec.Command("edgar", os.Args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Env = append(os.Environ(), "EDGAR_RECURSIVE=true")
			//cmd.Dir = statsite.StatsDirectory
			cmd.Run()
		} else {
			if err := statsite.OutputPages(); err != nil {
				log.Fatal(err)
			}
			if err := statsite.GenerateIndexPages(); err != nil {
				log.Fatal(err)
			}
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
