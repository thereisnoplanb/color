package color_test

import (
	"testing"

	"github.com/thereisnoplanb/color"
)

func TestColor_Red(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		color   color.Color
		wantRed uint8
	}{
		{
			name:    "Test Red() with a Red color (pure red)",
			color:   color.Red,
			wantRed: 255,
		},
		{
			name:    "Test Red() with a Lime color (pure green)",
			color:   color.Lime,
			wantRed: 0,
		},
		{
			name:    "Test Red() with a Blue color (pure blue)",
			color:   color.Blue,
			wantRed: 0,
		},
		{
			name:    "Test Red() with a Maroon color",
			color:   color.Maroon,
			wantRed: 128,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			red := tt.color.Red()
			if red != tt.wantRed {
				t.Errorf("Red() = %v, want %v", red, tt.wantRed)
			}
		})
	}
}

func TestColor_Green(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		color     color.Color
		wantGreen uint8
	}{
		{
			name:      "Test Green() with a Red color (pure red)",
			color:     color.Red,
			wantGreen: 0,
		},
		{
			name:      "Test Green() with a Lime color (pure green)",
			color:     color.Lime,
			wantGreen: 255,
		},
		{
			name:      "Test Green() with a Blue color (pure blue)",
			color:     color.Blue,
			wantGreen: 0,
		},
		{
			name:      "Test Green() with a Green color",
			color:     color.Green,
			wantGreen: 128,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			green := tt.color.Green()
			if green != tt.wantGreen {
				t.Errorf("Green() = %v, want %v", green, tt.wantGreen)
			}
		})
	}
}

func TestColor_Blue(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		color    color.Color
		wantBlue uint8
	}{
		{
			name:     "Test Blue() with a Red color (pure red)",
			color:    color.Red,
			wantBlue: 0,
		},
		{
			name:     "Test Blue() with a Lime color (pure green)",
			color:    color.Lime,
			wantBlue: 0,
		},
		{
			name:     "Test Blue() with a Blue color (pure blue)",
			color:    color.Blue,
			wantBlue: 255,
		},
		{
			name:     "Test Blue() with a Navy color",
			color:    color.Navy,
			wantBlue: 128,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blue := tt.color.Blue()
			if blue != tt.wantBlue {
				t.Errorf("Blue() = %v, want %v", blue, tt.wantBlue)
			}
		})
	}
}

func TestColor_Alpha(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		color     color.Color
		wantAlpha uint8
	}{
		{
			name:      "Test Alpha() with a Transparent color",
			color:     color.Transparent,
			wantAlpha: 0,
		},
		{
			name:      "Test Alpha() with a White color",
			color:     color.White,
			wantAlpha: 255,
		},
		{
			name:      "Test Alpha() with a Black color",
			color:     color.Black,
			wantAlpha: 255,
		},
		{
			name:      "Test Alpha() with a Half-transparent color",
			color:     color.Color(0x80000000), // 50% transparent black
			wantAlpha: 128,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alpha := tt.color.Alpha()
			if alpha != tt.wantAlpha {
				t.Errorf("Alpha() = %v, want %v", alpha, tt.wantAlpha)
			}
		})
	}
}

func TestColor_Hue(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		color   color.Color
		wantHue float64
	}{
		{
			name:    "Test Hue() with a Red color (pure red)",
			color:   color.Red,
			wantHue: 0.0,
		},
		{
			name:    "Test Hue() with a Lime color (pure green)",
			color:   color.Lime,
			wantHue: 120.0,
		},
		{
			name:    "Test Hue() with a Blue color (pure blue)",
			color:   color.Blue,
			wantHue: 240.0,
		},
		{
			name:    "Test Hue() with a Yellow color",
			color:   color.Yellow,
			wantHue: 60.0,
		},
		{
			name:    "Test Hue() with a Cyan color",
			color:   color.Cyan,
			wantHue: 180.0,
		},
		{
			name:    "Test Hue() with a Magenta color",
			color:   color.Magenta,
			wantHue: 300.0,
		},
		{
			name:    "Test Hue() with a Black color",
			color:   color.Black,
			wantHue: 0.0,
		},
		{
			name:    "Test Hue() with a White color",
			color:   color.White,
			wantHue: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hue := tt.color.Hue()
			if hue != tt.wantHue {
				t.Errorf("Hue() = %v, want %v", hue, tt.wantHue)
			}
		})
	}
}

func TestColor_Saturation(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		color          color.Color
		wantSaturation float64
	}{
		{
			name:           "Test Saturation() with a Red color (pure red)",
			color:          color.Red,
			wantSaturation: 1.0,
		},
		{
			name:           "Test Saturation() with a Lime color (pure green)",
			color:          color.Lime,
			wantSaturation: 1.0,
		},
		{
			name:           "Test Saturation() with a Blue color (pure blue)",
			color:          color.Blue,
			wantSaturation: 1.0,
		},
		{
			name:           "Test Saturation() with a Yellow color",
			color:          color.Yellow,
			wantSaturation: 1.0,
		},
		{
			name:           "Test Saturation() with a Cyan color",
			color:          color.Cyan,
			wantSaturation: 1.0,
		},
		{
			name:           "Test Saturation() with a Magenta color",
			color:          color.Magenta,
			wantSaturation: 1.0,
		},
		{
			name:           "Test Saturation() with a Black color",
			color:          color.Black,
			wantSaturation: 0.0,
		},
		{
			name:           "Test Saturation() with a White color",
			color:          color.White,
			wantSaturation: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saturation := tt.color.Saturation()
			if saturation != tt.wantSaturation {
				t.Errorf("Saturation() = %v, want %v", saturation, tt.wantSaturation)
			}
		})
	}
}

func TestColor_Lightness(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		color         color.Color
		wantLightness float64
	}{
		{
			name:          "Test Lightness() with a Red color (pure red)",
			color:         color.Red,
			wantLightness: 0.5,
		},
		{
			name:          "Test Lightness() with a Lime color (pure green)",
			color:         color.Lime,
			wantLightness: 0.5,
		},
		{
			name:          "Test Lightness() with a Blue color (pure blue)",
			color:         color.Blue,
			wantLightness: 0.5,
		},
		{
			name:          "Test Lightness() with a Maroon color",
			color:         color.Maroon,
			wantLightness: 64.0 / 255.0,
		},
		{
			name:          "Test Lightness() with a Navy color",
			color:         color.Navy,
			wantLightness: 64.0 / 255.0,
		},
		{
			name:          "Test Lightness() with a Green color",
			color:         color.Green,
			wantLightness: 64.0 / 255.0,
		},
		{
			name:          "Test Lightness() with a Black color",
			color:         color.Black,
			wantLightness: 0.0,
		},
		{
			name:          "Test Lightness() with a White color",
			color:         color.White,
			wantLightness: 1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lightness := tt.color.Lightness()
			if lightness != tt.wantLightness {
				t.Errorf("Lightness() = %v, want %v", lightness, tt.wantLightness)
			}
		})
	}
}
