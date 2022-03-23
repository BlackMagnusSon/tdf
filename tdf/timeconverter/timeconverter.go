package timeconverter

import (
	"strconv"
	"strings"
	"time"
)

//Ut is структура для времени
type Ut struct {
	Timestamps string
	UnixOrUdf  bool // "unix" if true, "udf" is false
}

//Converter is конвертирует время
func (p Ut) Converter(k chan time.Time) {
	loc, _ := time.LoadLocation("UTC")
	if p.UnixOrUdf {
		timetemplate := "20060102T150405"
		tm, err := time.Parse(timetemplate[:len(p.Timestamps)], p.Timestamps)
		if err != nil {
			panic(err)
		}
		k <- tm.In(loc)
		return
	} else {
		i, err := strconv.ParseInt(p.Timestamps, 10, 64)
		if err != nil {
			panic(err)
		}
		tm := time.Unix(i, 0)
		k <- tm.In(loc)
		return
	}

}

//TimeConvert Converting time
func TimeConvert(useTime string) string {
	if useTime == "now" {
		return strconv.FormatInt(time.Now().Unix(), 10)
	}
	useUdf := strings.Contains(useTime, "T")
	var timeToTest = Ut{Timestamps: useTime, UnixOrUdf: useUdf}
	tm := make(chan time.Time)
	go timeToTest.Converter(tm)
	k := <-tm
	if useUdf {
		return strconv.FormatInt(k.Unix(), 10)

	} else {
		return (k.Format("20060102T150405"))

	}
}
