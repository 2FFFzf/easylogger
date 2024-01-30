package easylogger

import (
	"testing"
)

func TestInitLogger(t *testing.T) {
	if got := InitLogger("./log.env"); got != nil {
		t.Errorf("Hello() = %q, want %v", got, nil)
	}
}
