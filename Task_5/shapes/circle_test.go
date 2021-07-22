package shapes

import (
	"math"
	"testing"
)
func TestCircle_Area(t *testing.T) {
	tests := []struct {
		name string
		c    Circle
		want float64
	}{
		{name: "TestWithPositiveRadius",
			c:    Circle{Radius: 1},
			want: math.Pi},

		{name: "TestWithZeroRadius",
			c:    Circle{Radius: 0},
			want: 0},

		{name: "TestWithNegativeRadius",
			c:    Circle{Radius: -1},
			want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				Radius: tt.c.Radius,
			}
			if got, _ := c.Area(); got != tt.want {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestCircle_Perimeter(t *testing.T) {
	tests := []struct {
		name string
		c    Circle
		want float64
	}{
		{name: "TestWithPositiveParameters",
			c:    Circle{Radius: 1},
			want: 2 * math.Pi,
		},
		{name: "TestWithZeroRadius",
			c:    Circle{Radius: 0},
			want: 0},

		{name: "TestWithNegativeRadius",
			c:    Circle{Radius: -1},
			want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				Radius: tt.c.Radius,
			}
			if got, _ := c.Perimeter(); got != tt.want {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}
