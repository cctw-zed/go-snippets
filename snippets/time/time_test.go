package time

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	nativeRegisterTime := time.Unix(int64(1633063803), 0)
	eidRegisterTime := time.Unix(int64(1633063807), 0)
	duration := eidRegisterTime.Sub(nativeRegisterTime)
	if duration > 5*time.Second || duration < -5*time.Second {
		t.Errorf("!!")
	}
}
