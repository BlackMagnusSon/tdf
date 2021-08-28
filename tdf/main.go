package main

import (
	"flag"
	"fmt"

	tc "tdf/timeconverter"
)

func main() {
	useUdf := flag.Bool("udf", false, "convert udf from udf or not")
	useTime := flag.String("t", "20200115T", "time")
	flag.Parse()
	var timeToTest = tc.Ut{Timestamps: *useTime, UnixOrUdf: *useUdf}
	var k = timeToTest.Converter()
	if *useUdf {
		fmt.Println(k.Unix())

	} else {
		fmt.Printf("%s\n", k.Format("20060102T150405"))

	}
}
