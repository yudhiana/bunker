package bunker

import (
	"time"
)

type gotimes struct {
	today    time.Time
	holidays []time.Time
	weekend  map[int]bool
	nexttime time.Time
}

type GoTimes struct {
	Time time.Time
	*gotimes
}

func newTime() goTimesImpl {
	return &gotimes{
		today: time.Now(),
		weekend: map[int]bool{
			int(time.Saturday): true,
			int(time.Sunday):   true,
		},
	}
}

func (gt *gotimes) isWeekend(d time.Time) (w bool) {
	return gt.weekend[int(d.Weekday())]
}

func (gt *gotimes) SetToday(today *time.Time) *gotimes {
	if today != nil {
		gt.today = *today
	}
	return gt
}

func (gt *gotimes) SetHolidays(h []time.Time) *gotimes {
	gt.holidays = append(gt.holidays, h...)
	return gt
}

func (gt *gotimes) AddWeekDay(sla int) (t time.Time) {
	for i := 1; i <= sla; i++ {
		gt.nexttime = gt.addDate(0, 0, i)
		if gt.isWeekend(gt.nexttime) {
			sla++
		} else {
			if len(gt.holidays) > 0 {
				if gt.isOffDay(gt.nexttime) {
					sla++
				}
			}
			gt.nexttime = gt.addDate(0, 0, i)
		}
	}
	return gt.nexttime
}

func (gt *gotimes) isOffDay(d time.Time) (h bool) {
	for _, dayOff := range gt.holidays {
		if dayOff.Equal(d) {
			return true
		}
	}
	return
}

func (gt gotimes) addDate(year, month, day int) (t time.Time) {
	return gt.today.AddDate(year, month, day)
}

func AddWeekDay(day int, today *time.Time, holidays []time.Time) time.Time {
	return newTime().SetToday(today).SetHolidays(holidays).AddWeekDay(day)
}
