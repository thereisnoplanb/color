package color_test

import (
	"testing"

	"github.com/thereisnoplanb/color"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		layout    color.Format
		value     string
		wantColor color.Color
		wantErr   bool
	}{
		{
			name:      "ARGB valid",
			layout:    color.ARGB,
			value:     "FF112233",
			wantColor: color.FromARGB(255, 17, 34, 51),
		},
		{
			name:    "ARGB invalid length",
			layout:  color.ARGB,
			value:   "FF1122",
			wantErr: true,
		},
		{
			name:      "RGB valid",
			layout:    color.RGB,
			value:     "112233",
			wantColor: color.FromRGB(17, 34, 51),
		},
		{
			name:    "RGB invalid length",
			layout:  color.RGB,
			value:   "11223344",
			wantErr: true,
		},
		{
			name:      "RGBA valid",
			layout:    color.RGBA,
			value:     "11223344",
			wantColor: color.FromARGB(68, 17, 34, 51),
		},
		{
			name:    "RGBA invalid length",
			layout:  color.RGBA,
			value:   "112233",
			wantErr: true,
		},
		{
			name:      "HARGB valid",
			layout:    color.HARGB,
			value:     "#FF112233",
			wantColor: color.FromARGB(255, 17, 34, 51),
		},
		{
			name:    "HARGB invalid length",
			layout:  color.HARGB,
			value:   "#FF1122",
			wantErr: true,
		},
		{
			name:    "HARGB missing #",
			layout:  color.HARGB,
			value:   "FF112233",
			wantErr: true,
		},
		{
			name:      "Unknown format",
			layout:    "999",
			value:     "112233",
			wantColor: color.FromRGB(17, 34, 51),
		},
		{
			name:    "Unknown format 2",
			layout:  "999",
			value:   "notanumber",
			wantErr: true,
		},
		{
			name:      "INT valid",
			layout:    color.INT,
			value:     "4278190080", // 0xFF000000
			wantColor: color.FromARGB(255, 0, 0, 0),
		},
		{
			name:    "INT invalid format",
			layout:  color.INT,
			value:   "0xFF000000",
			wantErr: true,
		},
		{
			name:    "INT invalid number",
			layout:  color.INT,
			value:   "notanumber",
			wantErr: true,
		},
		{
			name:      "CSSRGB valid",
			layout:    color.CSSRGB,
			value:     "rgb(17, 34, 51)",
			wantColor: color.FromRGB(17, 34, 51),
		},
		{
			name:    "CSSRGB invalid format",
			layout:  color.CSSRGB,
			value:   "17, 34, 51",
			wantErr: true,
		},
		{
			name:      "CSSRGB invalid value",
			layout:    color.CSSRGB,
			value:     "rgb(1017, 34, 51)",
			wantColor: color.FromRGB(17, 34, 51),
		},
		{
			name:      "CSSRGBA valid",
			layout:    color.CSSRGBA,
			value:     "rgba(17, 34, 51, 50%)",
			wantColor: color.FromRGBA(17, 34, 51, 128),
		},
		{
			name:      "CSSRGBA valid",
			layout:    color.CSSRGBA,
			value:     "rgba(17, 34, 51, 0.5)",
			wantColor: color.FromRGBA(17, 34, 51, 128),
		},
		{
			name:    "CSSRGBA invalid format",
			layout:  color.CSSRGBA,
			value:   "17, 34, 51, 0.5",
			wantErr: true,
		},
		{
			name:      "CSS4RGB valid",
			layout:    color.CSS4RGB,
			value:     "rgb(17 34 51)",
			wantColor: color.FromRGB(17, 34, 51),
		},
		{
			name:    "CSS4RGB invalid format",
			layout:  color.CSS4RGB,
			value:   "17 34 51",
			wantErr: true,
		},
		{
			name:      "CSS4RGBA valid",
			layout:    color.CSS4RGBA,
			value:     "rgba(17 34 51 / 50%)",
			wantColor: color.FromRGBA(17, 34, 51, 128),
		},
		{
			name:      "CSS4RGBA valid",
			layout:    color.CSS4RGBA,
			value:     "rgba(17 34 51 / 0.5)",
			wantColor: color.FromRGBA(17, 34, 51, 128),
		},
		{
			name:    "CSS4RGBA invalid format",
			layout:  color.CSS4RGBA,
			value:   "17 34 51 / 50%",
			wantErr: true,
		},
		{
			name:      "HSL valid",
			layout:    color.HSL,
			value:     "hsl(120, 100%, 50%)",
			wantColor: color.MustHSL(120, 1, 0.5),
		},
		{
			name:    "HSL invalid format",
			layout:  color.HSL,
			value:   "120, 100%, 50%",
			wantErr: true,
		},
		{
			name:      "HSLA valid",
			layout:    color.HSLA,
			value:     "hsla(120, 100%, 50%, 0.5)",
			wantColor: color.MustHSLA(120, 1, 0.5, 127),
		},
		{
			name:    "HSLA invalid format",
			layout:  color.HSLA,
			value:   "120, 100%, 50%, 0.5",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			color, err := color.Parse(tt.layout, tt.value)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Parse() failed: %v", err)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Parse() succeeded unexpectedly")
			}
			if color != tt.wantColor {
				t.Errorf("Parse() = %v, want %v", color, tt.wantColor)
			}
		})
	}
}
