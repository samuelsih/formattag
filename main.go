package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/momaek/formattag/align"
)

var version string

func main() {
	var (
		showVersion    bool
		writeToConsole bool
	)

	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&writeToConsole, "C", false, "Write result to console")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if showVersion {
		fmt.Println("Version:", version)
		return
	}

	dirfs := os.DirFS(".")

	fs.WalkDir(dirfs, ".", func(p string, info fs.DirEntry, err error) error {
		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}

		align.Init(p)

		b, errWalk := align.Do()
		if errWalk != nil {
			log.Fatalf("Align failed for %s - %v", p, errWalk)
		}

		if writeToConsole {
			fmt.Println(string(b))
			return nil
		}

		errWalk = os.WriteFile(p, b, 0)
		if errWalk != nil {
			log.Fatalf("Cannot write to file %s - %v", p, errWalk)
		}

		return nil
	})
}
