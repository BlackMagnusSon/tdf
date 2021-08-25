package timeconverter

import (
	"strconv"
	"time"
)

//Ut is структура для времени
type Ut struct {
	Timestamps string
	UnixOrUdf  bool // "unix" if true, "udf" is false
}

//Converter is конвертирует время
func (p Ut) Converter() time.Time {
	loc, _ := time.LoadLocation("UTC")
	if p.UnixOrUdf == true {
		timetemplate := "20060102T150405"
		tm, err := time.Parse(timetemplate[:len(p.Timestamps)], p.Timestamps)
		if err != nil {
			panic(err)
		}
		return tm.In(loc)
	} else {
		i, err := strconv.ParseInt(p.Timestamps, 10, 64)
		if err != nil {
			panic(err)
		}
		tm := time.Unix(i, 0)
		return tm.In(loc)
	}

}
