package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mechiko/utility"
	"go.uber.org/zap"
)

var fileExe string
var dir string

// если home true то папка создается локально
// var home = flag.Bool("home", false, "")
var file = flag.String("file", "", "json")
var dbA3 = flag.String("db", "", "db")
var order = flag.String("order", "", "order number")
var serial = flag.Int("serial", 7, "lenght serial TG")
var gtin = flag.String("gtin", "", "gtin")

func init() {
	flag.Parse()
	exe, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get executable path: %v\n", err)
		os.Exit(1)
	}
	fileExe = exe
	dir, err = filepath.Abs(filepath.Dir(fileExe))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get absolute path: %v\n", err)
		os.Exit(1)
	}
	if err := os.Chdir(dir); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to change directory: %v\n", err)
		os.Exit(1)
	}
}

func errMessageExit(loger *zap.SugaredLogger, title string, err error) {
	if loger != nil {
		loger.Errorf("%s %v", title, err)
	}
	if err != nil {
		utility.MessageBox(title, err.Error())
	} else {
		utility.MessageBox(title, "")
	}

	os.Exit(1)
}
