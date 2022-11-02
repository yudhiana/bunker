package chronos

import "time"

type goTimesImpl interface {
	SetToday(today *time.Time) *gotimes
	SetHolidays(h []time.Time) *gotimes
	AddWeekDay(sla int) (t time.Time)
}
