package retry

import (
	"fmt"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	cnt := 0
	err := fmt.Errorf("test error every time")
	f := func() error {
		if cnt < 2 {
			cnt++
			return err
		} else {
			cnt++
			return nil
		}
	}
	errFn := Retry(3, 1*time.Second, f)
	if errFn != nil {
		t.Error(errFn)
	} else {
		t.Log("pass")
	}
}
