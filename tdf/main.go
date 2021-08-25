package main

import (
	"flag"
	"fmt"
	tc "tdf/timeconverter"
)

func main() {
	var source string
	fmt.Scan(&source)
	fmt.Printf("\n\n\n %s \n", source)

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
