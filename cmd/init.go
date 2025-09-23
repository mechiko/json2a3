package main

import (
	"flag"
	"os"

	"github.com/mechiko/utility"
	"go.uber.org/zap"
)

var file = flag.String("file", "", "json")
var dbA3 = flag.String("db", "", "db")
var order = flag.String("order", "", "order number")
var serial = flag.Int("serial", 7, "lenght serial TG")
var gtin = flag.String("gtin", "", "gtin")

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
