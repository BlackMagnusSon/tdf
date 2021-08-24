package timeconverter

import (
	"strconv"
	"time"
)

type ut struct {
	timestamps string
	unixOrUdf  bool // "unix" if true, "udf" is false
}

func (p ut) converter() int64 {
	if p.unixOrUdf == true {
		timetemplate := "20060102T150405"
		tm, err := time.Parse(p.timestamps, timetemplate[:len(p.timestamps)])
		if err != nil {
			panic(err)
		} else {
			return tm
		}
	} else {
		i, err := strconv.ParseInt("1405544146", 10, 64)
		if err != nil {
			panic(err)
		}
		tm := time.Unix(i, 0)
		return tm
	}
	return 0
}
