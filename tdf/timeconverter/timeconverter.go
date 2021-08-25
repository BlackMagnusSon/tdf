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
	if p.UnixOrUdf == true {
		timetemplate := "20060102T150405"
		tm, err := time.Parse(p.Timestamps, timetemplate[:len(p.Timestamps)])
		if err != nil {
			panic(err)
		}
		return tm
	} else {
		i, err := strconv.ParseInt("1405544146", 10, 64)
		if err != nil {
			panic(err)
		}
		tm := time.Unix(i, 0)
		return tm
	}

}
