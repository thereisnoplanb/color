package color

import "fmt"

// Creates a Color from RGB components (red, green, blue) with full opacity (alpha = 255).
//
// # Parameters
//
//	red uint8
//
// The red component of the color, as a value between 0 and 255.
//
//	green uint8
//
// The green component of the color, as a value between 0 and 255.
//
//	blue uint8
//
// The blue component of the color, as a value between 0 and 255.
//
// # Returns
//
//	color Color
//
// A Color value representing the specified RGB color with full opacity.
func FromRGB(red, green, blue uint8) (color Color) {
	return Color(uint32(0xFF)<<24 | uint32(red)<<16 | uint32(green)<<8 | uint32(blue))
}

// Creates a Color from ARGB components (alpha, red, green, blue).
//
// # Parameters
//
//	alpha uint8
//
// The alpha component of the color, as a value between 0 and 255.
//
//	red uint8
//
// The red component of the color, as a value between 0 and 255.
//
//	green uint8
//
// The green component of the color, as a value between 0 and 255.
//
//	blue uint8
//
// The blue component of the color, as a value between 0 and 255.
//
// # Returns
//
//	color Color
//
// A Color value representing the specified ARGB color.
func FromARGB(alpha, red, green, blue uint8) (color Color) {
	return Color(uint32(alpha)<<24 | uint32(red)<<16 | uint32(green)<<8 | uint32(blue))
}

// Creates a Color from RGBA components (red, green, blue, alpha).
//
// # Parameters
//
//	red uint8
//
// The red component of the color, as a value between 0 and 255.
//
//	green uint8
//
// The green component of the color, as a value between 0 and 255.
//
//	blue uint8
//
// The blue component of the color, as a value between 0 and 255.
//
//	alpha uint8
//
// The alpha component of the color, as a value between 0 and 255.
//
// # Returns
//
//	color Color
//
// A Color value representing the specified RGBA color.
func FromRGBA(red, green, blue, alpha uint8) (color Color) {
	return Color(uint32(alpha)<<24 | uint32(red)<<16 | uint32(green)<<8 | uint32(blue))
}

// Creates a Color from HSL components (hue, saturation, lightness).
//
// # Parameters
//
//	hue float64
//
// The hue component of the color in the HSL color space, as a value between 0 and 360.
// 0.0 represents red, 120.0 represents green, 240.0 represents blue, 360.0 represents red again.
//
//	saturation float64
//
// The saturation component of the color in the HSL color space, as a value between 0 and 1.
// 0.0 represents a shade of gray and 1.0 represents a fully saturated color.
//
//	lightness float64
//
// The lightness component of the color in the HSL color space, as a value between 0 and 1.
// 0.0 represents black, 1.0 represents white, and 0.5 represents the pure color.
//
// # Returns
//
//	color Color
//
// A Color value representing the specified HSL color.
//
//	err error
//
// An error if the input values are not in a valid range.
func FromHSL(hue, saturation, lightness float64) (color Color, err error) {
	err = validateHSL(hue, saturation, lightness)
	if err != nil {
		return 0, err
	}
	red, green, blue := HSLToRGB(hue, saturation, lightness)
	return FromRGB(red, green, blue), nil
}

// Creates a Color from HSL components (hue, saturation, lightness).
// Panics if the input values are not in a valid range.
//
// # Parameters
//
//	hue float64
//
// The hue component of the color in the HSL color space, as a value between 0 and 360.
// 0.0 represents red, 120.0 represents green, 240.0 represents blue, 360.0 represents red again.
//
//	saturation float64
//
// The saturation component of the color in the HSL color space, as a value between 0 and 1.
// 0.0 represents a shade of gray and 1.0 represents a fully saturated color.
//
//	lightness float64
//
// The lightness component of the color in the HSL color space, as a value between 0 and 1.
// 0.0 represents black, 1.0 represents white, and 0.5 represents the pure color.
//
// # Returns
//
//	color Color
//
// A Color value representing the specified HSL color.
func MustHSL(hue, saturation, lightness float64) Color {
	color, err := FromHSL(hue, saturation, lightness)
	if err != nil {
		panic(err)
	}
	return color
}

// Creates a Color from HSLA components (hue, saturation, lightness, alpha).
//
// # Parameters
//
//	hue float64
//
// The hue component of the color in the HSL color space, as a value between 0 and 360.
// 0.0 represents red, 120.0 represents green, 240.0 represents blue, 360.0 represents red again.
//
//	saturation float64
//
// The saturation component of the color in the HSL color space, as a value between 0 and 1.
// 0.0 represents a shade of gray and 1.0 represents a fully saturated color.
//
//	lightness float64
//
// The lightness component of the color in the HSL color space, as a value between 0 and 1.
// 0.0 represents black, 1.0 represents white, and 0.5 represents the pure color.
//
//	alpha uint8
//
// The alpha component of the color, as a value between 0 and 255.
//
// # Returns
//
//	color Color
//
// A Color value representing the specified HSLA color.
//
//	err error
//
// An error if the input values are not in a valid range.
func FromHSLA(hue, saturation, lightness float64, alpha uint8) (color Color, err error) {
	err = validateHSL(hue, saturation, lightness)
	if err != nil {
		return 0, err
	}
	red, green, blue := HSLToRGB(hue, saturation, lightness)
	return FromARGB(alpha, red, green, blue), nil
}

// Creates a Color from HSLA components (hue, saturation, lightness, alpha).
//
// # Parameters
//
//	hue float64
//
// The hue component of the color in the HSL color space, as a value between 0 and 360.
// 0.0 represents red, 120.0 represents green, 240.0 represents blue, 360.0 represents red again.
//
//	saturation float64
//
// The saturation component of the color in the HSL color space, as a value between 0 and 1.
// 0.0 represents a shade of gray and 1.0 represents a fully saturated color.
//
//	lightness float64
//
// The lightness component of the color in the HSL color space, as a value between 0 and 1.
// 0.0 represents black, 1.0 represents white, and 0.5 represents the pure color.
//
//	alpha uint8
//
// The alpha component of the color, as a value between 0 and 255.
//
// # Returns
//
//	color Color
//
// A Color value representing the specified HSLA color.
//
//	err error
//
// An error if the input values are not in a valid range.
func MustHSLA(hue, saturation, lightness float64, alpha uint8) (color Color) {
	err := validateHSL(hue, saturation, lightness)
	if err != nil {
		panic(err)
	}
	red, green, blue := HSLToRGB(hue, saturation, lightness)
	return FromARGB(alpha, red, green, blue)
}

func validateHSL(hue, saturation, lightness float64) (err error) {
	if hue < 0 || hue >= 360 {
		return fmt.Errorf("hue out of range: %f (expected 0 - 360)", hue)
	}
	if saturation < 0 || saturation > 1 {
		return fmt.Errorf("saturation out of range: %f (expected 0 - 1)", saturation)
	}
	if lightness < 0 || lightness > 1 {
		return fmt.Errorf("lightness out of range: %f (expected 0 - 1)", lightness)
	}
	return nil
}

func (this Color) unpack() (alpha, red, green, blue uint8) {
	alpha = uint8(this >> 24)
	red = uint8(this >> 16)
	green = uint8(this >> 8)
	blue = uint8(this)
	return
}
