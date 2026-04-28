package color_test

import (
	"testing"

	"github.com/thereisnoplanb/color"
)

func TestColor_Format(t *testing.T) {
	tests := []struct {
		name   string // description of this test case
		color  color.Color
		format color.Format
		want   string
	}{
		{
			name:   "Test Format with INT format",
			color:  color.FromARGB(255, 255, 0, 0), // Red color
			format: color.INT,
			want:   "4294901760",
		},
		{
			name:   "Test Format with HRGB format",
			color:  color.FromARGB(255, 255, 0, 0), // Red color
			format: color.HRGB,
			want:   "#FF0000",
		},
		{
			name:   "Test Format with RGB format",
			color:  color.FromARGB(255, 255, 0, 0), // Red color
			format: color.RGB,
			want:   "FF0000",
		},
		{
			name:   "Test Format with HARGB format",
			color:  color.FromARGB(255, 255, 0, 0), // Red color
			format: color.HARGB,
			want:   "#FFFF0000",
		},
		{
			name:   "Test Format with ARGB format",
			color:  color.FromARGB(255, 255, 0, 0), // Red color
			format: color.ARGB,
			want:   "FFFF0000",
		},
		{
			name:   "Test Format with HRGBA format",
			color:  color.FromARGB(255, 255, 0, 0), // Red color
			format: color.HRGBA,
			want:   "#FF0000FF",
		},
		{
			name:   "Test Format with RGBA format",
			color:  color.FromARGB(255, 255, 0, 0), // Red color
			format: color.RGBA,
			want:   "FF0000FF",
		},
		{
			name:   "Test Format with HSL format",
			color:  color.FromARGB(255, 255, 0, 0), // Red color
			format: color.HSL,
			want:   "hsl(0, 100%, 50%)",
		},
		{
			name:   "Test Format with HSLA format",
			color:  color.FromARGB(255, 255, 0, 0), // Red color
			format: color.HSLA,
			want:   "hsla(0, 100%, 50%, 1)",
		},
		{
			name:   "Test Format with CSSRGB format",
			color:  color.FromARGB(255, 255, 0, 0), // Red color
			format: color.CSSRGB,
			want:   "rgb(255, 0, 0)",
		},
		{
			name:   "Test Format with CSSRGBA format",
			color:  color.FromARGB(255, 255, 0, 0), // Red color
			format: color.CSSRGBA,
			want:   "rgba(255, 0, 0, 1)",
		},
		{
			name:   "Test Format with CSS4RGB format",
			color:  color.FromARGB(255, 255, 0, 0), // Red color
			format: color.CSS4RGB,
			want:   "rgb(255 0 0)",
		},
		{
			name:   "Test Format with CSS4RGBA format",
			color:  color.FromARGB(255, 255, 0, 0), // Red color
			format: color.CSS4RGBA,
			want:   "rgba(255 0 0 / 1)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.color.Format(tt.format)
			if got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
