package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	tc "github.com/BlackMagnusSon/tdf/timeconverter"
)

func main() {
	useTime := flag.String("t", "20200115T", "time")
	flag.Parse()
	useUdf := strings.Contains(*useTime, "T")
	var timeToTest = tc.Ut{Timestamps: *useTime, UnixOrUdf: useUdf}
	tm := make(chan time.Time)
	go timeToTest.Converter(tm)
	k := <-tm
	if useUdf {
		fmt.Println(k.Unix())

	} else {
		fmt.Printf("%s\n", k.Format("20060102T150405"))

	}
}
