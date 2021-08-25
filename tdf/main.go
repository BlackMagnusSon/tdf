package main

import (
	"fmt"
	tc "tdf/timeconverter"
)

func main() {
	var timeToTest = tc.Ut{Timestamps: "19999999", UnixOrUdf: false}
	var k = timeToTest.Converter()

	fmt.Printf("Time is: %s\n", k.Format("15:04:05.999"))
	fmt.Printf("Time is: %s\n", k.Format("15:04:05.999"))
}
