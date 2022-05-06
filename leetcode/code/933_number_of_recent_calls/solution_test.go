package _33_number_of_recent_calls

import "testing"

func TestRecentCounter_Ping(t *testing.T) {
	counter := Constructor()
	if 1 != counter.Ping(1) {
		t.Fatalf("failed")
	}
	if 2 != counter.Ping(100) {
		t.Fatalf("failed")
	}
	if 3 != counter.Ping(3001) {
		t.Fatalf("failed")
	}
	if 3 != counter.Ping(3002) {
		t.Fatalf("failed")
	}
}

func TestRecentCounter_Ping1(t *testing.T) {
	counter := Constructor()
	if 1 != counter.Ping(642) {
		t.Fatalf("failed")
	}
	if 2 != counter.Ping(1849) {
		t.Fatalf("failed")
	}
	if 1 != counter.Ping(4921) {
		t.Fatalf("failed")
	}
	if 2 != counter.Ping(5936) {
		t.Fatalf("failed")
	}
	if 3 != counter.Ping(5957) {
		t.Fatalf("failed")
	}
}
