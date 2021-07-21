package shapes

import "testing"

func TestRectangle_Area(t *testing.T) {
	
	tests := []struct {
		name   string
		r Rectangle
		want   float64
	}{
		{name: "TestWithPositiveParameters",
			r: Rectangle{Height: 1, Width: 1},
			want: 1,
		},
		{name: "TestWithZeroParameters",
			r: Rectangle{Height: 0, Width: 0},
			want: 0,
		},
		{name: "TestWithNegativeParameters",
			r: Rectangle{Height: -1, Width: -1},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				Height: tt.r.Height,
				Width:  tt.r.Width,
			}
			if got := r.Area(); got != tt.want {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Perimeter(t *testing.T) {

	tests := []struct {
		name   string
		r Rectangle
		want   float64
	}{
		{name: "TestWithPositiveParameters",
			r: Rectangle{Height: 1, Width: 1},
			want: 4,
		},
		{name: "TestWithZeroParameters",
			r: Rectangle{Height: 0, Width: 0},
			want: 0,
		},
		{name: "TestWithNegativeParameters",
			r: Rectangle{Height: -1, Width: -1},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				Height: tt.r.Height,
				Width:  tt.r.Width,
			}
			if got := r.Perimeter(); got != tt.want {
				t.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

