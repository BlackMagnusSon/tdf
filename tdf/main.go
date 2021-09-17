package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	tc "github.com/BlackMagnusSon/tdf/timeconverter"
)

func main() {
	useTime := flag.String("t", "20200115T", "time")
	flag.Parse()

	IsItRange := strings.Contains(*useTime, "-")

	if IsItRange {
		fromAndTo := strings.Split(*useTime, "-")
		for i := 0; i < len(fromAndTo); i++ {
			tm := fromAndTo[i]
			fromAndTo[i] = tc.TimeConvert(tm)

		}
		stValue, err := strconv.ParseInt(fromAndTo[0], 10, 64)
		ndValue, err := strconv.ParseInt(fromAndTo[1], 10, 64)
		if err != nil {
			panic(err)
		}
		if stValue < ndValue {
			fmt.Printf("from=%d&to=%d\n", stValue, ndValue)
		} else {
			fmt.Printf("from=%d&to=%d\n", ndValue, stValue)
		}

	} else {
		fmt.Print(tc.TimeConvert(*useTime) + "\n")
	}

}
