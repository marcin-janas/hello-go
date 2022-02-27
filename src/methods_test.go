package main

import (
	"testing"
)

func TestArea(t *testing.T) {
	tests := []struct {
		name string
		rect rectangle
		want int
	}{
		{
			"empty rectangle",
			rectangle{"empty", 0, 1},
			0,
		},
		{
			"special rectangle",
			rectangle{"square", 2, 2},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rect.area(); got != tt.want {
				t.Errorf("%s: want[%v] != got[%v]", tt.name, tt.want, got)
			}
		})
	}
}
