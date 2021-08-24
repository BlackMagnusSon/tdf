package timeconverter

import (
	"strconv"
	"time"
)

type ut struct {
	timestamps string
	unixOrUdf  bool // "unix" if true, "udf" is false
}

func (p ut) converter() time.Time {
	var tm time.Time = nil
	if p.unixOrUdf == true {
		timetemplate := "20060102T150405"
		tm, err := time.Parse(p.timestamps, timetemplate[:len(p.timestamps)])
		if err != nil {
			panic(err)
		}
	} else {
		i, err := strconv.ParseInt("1405544146", 10, 64)
		if err != nil {
			panic(err)
		}
		tm := time.Unix(i, 0)
	}
	return tm

}
