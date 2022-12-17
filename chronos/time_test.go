package bunker

import (
	"testing"
	"time"
)

func TestBunker(t *testing.T) {
	t.Run("testAddWeekDay", func(t *testing.T) {
		expected := "2022-10-12"
		today, _ := time.Parse("2006-01-02", "2022-10-03")
		actualUnsetHoliday := AddWeekDay(7, &today, nil).Format("2006-01-02")
		if actualUnsetHoliday != expected {
			t.Errorf("invalid addWeekDay\n\tExpected : %v\n\tActual : %v", expected, actualUnsetHoliday)
		}

		// Holidays in Indonesia
		// The Prophet Muhammad's Birthday,
		exampleHoliday, _ := time.Parse("2006-01-02", "2022-10-08")
		
		actualUnsetTodayAndHolidays := AddWeekDay(7, &today, []time.Time{
			exampleHoliday,
		}).Format("2006-01-02")
		if actualUnsetTodayAndHolidays != expected {
			t.Errorf("invalid addWeekDay\n\tExpected : %v\n\tActual : %v", expected, actualUnsetTodayAndHolidays)
		}
	})

}
