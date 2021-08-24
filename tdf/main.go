package main

import timeconverter "tdf/timeconverter"
import "fmt"

func main() {
	var tested int64 = 0 
	var time_To_Test =  timeconverter.ut{timestamps:"19999999",unixOrUdf:false}
	var tested = time_To_Test.converter()
	
	fmt.Printf("tTime is: %s\n", k.Format("15:04:05.999")
}
