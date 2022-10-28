package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var out = os.Stdout

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(out, "Error: %v\n", err)
		os.Exit(1)
	}
}

var (
	packagePath string
	srcPath     string
)

func init() {
	log.SetFlags(log.Lshortfile)
	flag.StringVar(&packagePath, "p", ".", "go package path")
	flag.StringVar(&srcPath, "s", "$GOPATH/src", "go source Path")
}

func run() error {
	flag.Parse()
	u, err := user.Current()
	if err != nil {
		return err
	}

	ac, err := filepath.Abs(strings.Replace(os.ExpandEnv(packagePath), "~", u.HomeDir, 1))
	if err != nil {
		return err
	}
	as, err := filepath.Abs(strings.Replace(os.ExpandEnv(srcPath), "~", u.HomeDir, 1))
	if err != nil {
		return err
	}
	packageName := strings.Replace(ac, as+"/", "", 1)
	if strings.HasPrefix(packageName, "/") {
		return fmt.Errorf("invalid input, go package Path %s, gosrc path %s", ac, as)
	}
	fmt.Fprint(out, packageName)
	return nil
}
