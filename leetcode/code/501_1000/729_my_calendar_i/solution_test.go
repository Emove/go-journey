package _29_my_calendar_i

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	calendar := Constructor()
	if !calendar.Book(10, 20) {
		t.Errorf("failed, want=%v, got=%v", true, false)
	}
	if calendar.Book(15, 25) {
		t.Errorf("failed, want=%v, got=%v", false, true)
	}
	if !calendar.Book(20, 30) {
		t.Errorf("failed, want=%v, got=%v", true, false)
	}

}
