package color

import "math"

// Converts between HSL and RGB color spaces.
//
// # Parameters
//
//   hue float64
//
// The hue component of the color in the HSL color space, as a value between 0 and 360.
// 0.0 represents red, 120.0 represents green, 240.0 represents blue, 360.0 represents red again.
//
//   saturation float64
//
// The saturation component of the color in the HSL color space, as a value between 0 and 1.
// 0.0 represents a shade of gray and 1.0 represents a fully saturated color.
//
//   lightness float64
//
// The lightness component of the color in the HSL color space, as a value between 0 and 1.
// 0.0 represents black, 1.0 represents white, and 0.5 represents the pure color.
//
// # Returns
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
func HSLToRGB(hue, saturation, lightness float64) (red, green, blue uint8) {
	hue = math.Mod(hue, 360)
	if hue < 0 {
		hue += 360
	}

	hue /= 360

	if saturation == 0 {
		v := uint8(lightness*255 + 0.5)
		return v, v, v
	}

	var q float64
	if lightness < 0.5 {
		q = lightness * (1 + saturation)
	} else {
		q = lightness + saturation - lightness*saturation
	}
	p := 2*lightness - q

	hue2rgb := func(p, q, t float64) float64 {
		if t < 0 {
			t += 1
		}
		if t > 1 {
			t -= 1
		}
		switch {
		case t < 1.0/6.0:
			return p + (q-p)*6*t
		case t < 1.0/2.0:
			return q
		case t < 2.0/3.0:
			return p + (q-p)*(2.0/3.0-t)*6
		default:
			return p
		}
	}

	rF := hue2rgb(p, q, hue+1.0/3.0)
	gF := hue2rgb(p, q, hue)
	bF := hue2rgb(p, q, hue-1.0/3.0)

	return uint8(rF*255 + 0.5), uint8(gF*255 + 0.5), uint8(bF*255 + 0.5)
}

// Converts between RGB and HSL color spaces.
//
// # Parameters
//
//   red uint8
//
// The red component of the color, as a value between 0 and 255.
//
//   green uint8
//
// The green component of the color, as a value between 0 and 255.
//
//   blue uint8
//
// The blue component of the color, as a value between 0 and 255.
//
// # Returns
//
//   hue float64
//
// The hue component of the color in the HSL color space, as a value between 0 and 360.
//
//   saturation float64
//
// The saturation component of the color in the HSL color space, as a value between 0 and 1.
//
//   lightness float64
//
// The lightness component of the color in the HSL color space, as a value between 0 and 1.
func RGBToHSL(red, green, blue uint8) (hue, saturation, lightness float64) {
	rf := float64(red) / 255
	gf := float64(green) / 255
	bf := float64(blue) / 255

	max := math.Max(rf, math.Max(gf, bf))
	min := math.Min(rf, math.Min(gf, bf))
	delta := max - min

	lightness = (max + min) / 2

	if delta == 0 {
		saturation = 0
		hue = 0
		return
	}

	if lightness < 0.5 {
		saturation = delta / (max + min)
	} else {
		saturation = delta / (2 - max - min)
	}

	switch max {
	case rf:
		hue = (gf - bf) / delta
		if gf < bf {
			hue += 6
		}
	case gf:
		hue = (bf-rf)/delta + 2
	case bf:
		hue = (rf-gf)/delta + 4
	}

	hue *= 60

	return
}
