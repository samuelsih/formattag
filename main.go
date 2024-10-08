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

	targetdir := "."
	if len(os.Args) == 2 {
		targetdir = os.Args[1]
	}

	fsys := os.DirFS(targetdir)

	err := fs.WalkDir(fsys, ".", func(p string, info fs.DirEntry, err error) error {
		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}

		align.Init(p)

		b, errWalk := align.Do()
		if errWalk != nil {
			return fmt.Errorf("align failed for %s - %w", p, errWalk)
		}

		if writeToConsole {
			fmt.Println(string(b))
			return nil
		}

		errWalk = os.WriteFile(p, b, 0)
		if errWalk != nil {
			return fmt.Errorf("cannot write to file %s - %w", p, errWalk)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
