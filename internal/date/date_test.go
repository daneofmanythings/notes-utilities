package date

import (
	"testing"
	"time"
)

func TestGenerateDate(t *testing.T) {
	testCases := []struct {
		name     string
		args     []string
		expected string
	}{
		{"default", []string{}, time.Now()},
	}
}
