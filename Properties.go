package color

// Gets the red component of this Color structure.
//
// # Returns
//
//	red uint8
//
// The red component of the color, as a value between 0 and 255.
func (this Color) Red() (red uint8) {
	return uint8(this >> 16)
}

// Gets the green component of this Color structure.
//
// # Returns
//
//	green uint8
//
// The green component of the color, as a value between 0 and 255.
func (this Color) Green() (green uint8) {
	return uint8(this >> 8)
}

// Gets the blue component of this Color structure.
//
// # Returns
//
//	blue uint8
//
// The blue component of the color, as a value between 0 and 255.
func (this Color) Blue() (blue uint8) {
	return uint8(this)
}

// Gets the alpha component of this Color structure.
//
// # Returns
//
//	alpha uint8
//
// The alpha component of the color, as a value between 0 and 255.
// 0 represents fully transparent and 255 represents fully opaque.
func (this Color) Alpha() (alpha uint8) {
	return uint8(this >> 24)
}

// Gets the hue-saturation-lightness (HSL) hue value, in degrees, for this Color structure.
//
// # Returns
//
//	hue float64
//
// The hue component of the color in the HSL color space, as a value between 0 and 360.
// 0.0 represents red, 120.0 represents green, 240.0 represents blue, 360.0 represents red again.
func (this Color) Hue() (hue float64) {
	_, r, g, b := this.unpack()
	h, _, _ := RGBToHSL(r, g, b)
	return h
}

// Gets the hue-saturation-lightness (HSL) saturation value for this Color structure.
//
// # Returns
//
//	saturation float64
//
// The saturation component of the color in the HSL color space, as a value between 0 and 1.
// 0.0 represents a shade of gray and 1.0 represents a fully saturated color.
func (this Color) Saturation() (saturation float64) {
	_, r, g, b := this.unpack()
	_, s, _ := RGBToHSL(r, g, b)
	return s
}

// Gets the hue-saturation-lightness (HSL) lightness value for this Color structure.
//
// # Returns
//
//	lightness float64
//
// The lightness component of the color in the HSL color space, as a value between 0 and 1.
// 0.0 represents black and 1.0 represents white.
func (this Color) Lightness() (lightness float64) {
	_, r, g, b := this.unpack()
	_, _, l := RGBToHSL(r, g, b)
	return l
}
