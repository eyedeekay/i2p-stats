package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/eyedeekay/i2p-stats/site"
)

var Docroot = docroot

func docroot() string {
	//home := os.Getenv("HOME")
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	i2p := filepath.Join(home, "i2p")
	eepsite := filepath.Join(i2p, "eepsite")
	docroot := filepath.Join(eepsite, "docroot")
	return docroot
}

// Get the user's home directory.
// Build the path to the i2p directory inside the home directory.
// Build the path to the eepsite directory inside the i2p directory.
// Build the path to the docroot directory inside the eepsite directory.
// Return the docroot path.

var runDir = flag.String("dir", ".", "directory to run from")

func main() {
	flag.Parse()
	if statsite, err := site.NewStatsSite(*runDir); err != nil {
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
		gitAddCmd := exec.Command("git", "add", statsite.StatsDirectory)
		gitAddCmd.Stdout = os.Stdout
		gitAddCmd.Stderr = os.Stderr
		gitAddCmd.Run()
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
