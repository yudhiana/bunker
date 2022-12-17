package bunker

import (
	"testing"
)

func BenchmarkAddWeekDay(b *testing.B) {
	AddWeekDay(7, nil, nil)
}
