package iteration

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	testStr := "a"
	repeatCount := 5
	repeated := Repeat(testStr, repeatCount)
	expected := strings.Repeat(testStr, repeatCount)

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
