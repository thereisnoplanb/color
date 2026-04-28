package color_test

import (
	"testing"

	"github.com/thereisnoplanb/color"
)

func TestHSLToRGB(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		hue        float64
		saturation float64
		lightness  float64
		wantRed    uint8
		wantGreen  uint8
		wantBlue   uint8
	}{
		{
			name:       "Test HSLToRGB with pure red color",
			hue:        0,
			saturation: 1,
			lightness:  0.5,
			wantRed:    255,
			wantGreen:  0,
			wantBlue:   0,
		},
		{
			name:       "Test HSLToRGB with pure green color",
			hue:        120,
			saturation: 1,
			lightness:  0.5,
			wantRed:    0,
			wantGreen:  255,
			wantBlue:   0,
		},
		{
			name:       "Test HSLToRGB with pure blue color",
			hue:        240,
			saturation: 1,
			lightness:  0.5,
			wantRed:    0,
			wantGreen:  0,
			wantBlue:   255,
		},
		{
			name:       "Test HSLToRGB with yellow color",
			hue:        60,
			saturation: 1,
			lightness:  0.5,
			wantRed:    255,
			wantGreen:  255,
			wantBlue:   0,
		},
		{
			name:       "Test HSLToRGB with cyan color",
			hue:        180,
			saturation: 1,
			lightness:  0.5,
			wantRed:    0,
			wantGreen:  255,
			wantBlue:   255,
		},
		{
			name:       "Test HSLToRGB with magenta color",
			hue:        300,
			saturation: 1,
			lightness:  0.5,
			wantRed:    255,
			wantGreen:  0,
			wantBlue:   255,
		},
		{
			name:       "Test HSLToRGB with gray color",
			hue:        0,
			saturation: 0,
			lightness:  0.5,
			wantRed:    128,
			wantGreen:  128,
			wantBlue:   128,
		},
		{
			name:       "Test HSLToRGB with black color",
			hue:        0,
			saturation: 0,
			lightness:  0,
			wantRed:    0,
			wantGreen:  0,
			wantBlue:   0,
		},
		{
			name:       "Test HSLToRGB with white color",
			hue:        0,
			saturation: 0,
			lightness:  1,
			wantRed:    255,
			wantGreen:  255,
			wantBlue:   255,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			red, green, blue := color.HSLToRGB(tt.hue, tt.saturation, tt.lightness)
			// TODO: update the condition below to compare got with tt.want.
			if red != tt.wantRed {
				t.Errorf("HSLToRGB() R = %v, want R = %v", red, tt.wantRed)
			}
			if green != tt.wantGreen {
				t.Errorf("HSLToRGB() G = %v, want G = %v", green, tt.wantGreen)
			}
			if blue != tt.wantBlue {
				t.Errorf("HSLToRGB() B = %v, want B = %v", blue, tt.wantBlue)
			}
		})
	}
}

func TestRGBToHSL(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		red            uint8
		green          uint8
		blue           uint8
		wantHue        float64
		wantSaturation float64
		wantLightness  float64
	}{
		{
			name:           "Test RGBToHSL with pure red color",
			red:            255,
			green:          0,
			blue:           0,
			wantHue:        0,
			wantSaturation: 1,
			wantLightness:  0.5,
		},
		{
			name:           "Test RGBToHSL with pure green color",
			red:            0,
			green:          255,
			blue:           0,
			wantHue:        120,
			wantSaturation: 1,
			wantLightness:  0.5,
		},
		{
			name:           "Test RGBToHSL with pure blue color",
			red:            0,
			green:          0,
			blue:           255,
			wantHue:        240,
			wantSaturation: 1,
			wantLightness:  0.5,
		},
		{
			name:           "Test RGBToHSL with yellow color",
			red:            255,
			green:          255,
			blue:           0,
			wantHue:        60,
			wantSaturation: 1,
			wantLightness:  0.5,
		},
		{
			name:           "Test RGBToHSL with cyan color",
			red:            0,
			green:          255,
			blue:           255,
			wantHue:        180,
			wantSaturation: 1,
			wantLightness:  0.5,
		},
		{
			name:           "Test RGBToHSL with magenta color",
			red:            255,
			green:          0,
			blue:           255,
			wantHue:        300,
			wantSaturation: 1,
			wantLightness:  0.5,
		},
		{
			name:           "Test RGBToHSL with gray color",
			red:            128,
			green:          128,
			blue:           128,
			wantHue:        0,
			wantSaturation: 0,
			wantLightness:  128.0 / 255.0,
		},
		{
			name:           "Test RGBToHSL with black color",
			red:            0,
			green:          0,
			blue:           0,
			wantHue:        0,
			wantSaturation: 0,
			wantLightness:  0,
		},
		{
			name:           "Test RGBToHSL with white color",
			red:            255,
			green:          255,
			blue:           255,
			wantHue:        0,
			wantSaturation: 0,
			wantLightness:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hue, saturation, lightness := color.RGBToHSL(tt.red, tt.green, tt.blue)
			// TODO: update the condition below to compare got with tt.want.
			if hue != tt.wantHue {
				t.Errorf("RGBToHSL() H = %v, want H = %v", hue, tt.wantHue)
			}
			if saturation != tt.wantSaturation {
				t.Errorf("RGBToHSL() S = %v, want S = %v", saturation, tt.wantSaturation)
			}
			if lightness != tt.wantLightness {
				t.Errorf("RGBToHSL() L = %v, want L = %v", lightness, tt.wantLightness)
			}
		})
	}
}
