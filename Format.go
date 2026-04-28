package color

import (
	"fmt"
	"strconv"
)

type Format string

const INT Format = "INT"
const HARGB Format = "#AARRGGBB"
const ARGB Format = "AARRGGBB"
const HRGB Format = "#RRGGBB"
const RGB Format = "RRGGBB"
const HRGBA Format = "#RRGGBBAA"
const RGBA Format = "RRGGBBAA"
const HSL Format = "hsl(H, S%, L%)"
const HSLA Format = "hsla(H, S%, L%, A)"
const CSSRGB Format = "rgb(R, G, B)"
const CSSRGBA Format = "rgba(R, G, B, A)"
const CSS4RGB Format = "rgb(R G B)"
const CSS4RGBA Format = "rgba(R G B / A)"

// Returns a string representation of the color in the specified format.
// If the format is not recognized, it defaults to the standard string representation.
//
// # Parameters
//
//	format Format
//
// The desired output format for the color string. Supported formats include:
//   - INT: A decimal integer representing the color (e.g., "4278190079").
//   - HARGB: A hex string in the format "#AARRGGBB" (e.g., "#FF0000FF").
//   - ARGB: A hex string in the format "AARRGGBB" (e.g., "FF0000FF").
//   - HRGB: A hex string in the format "#RRGGBB" (e.g., "#FF0000").
//   - RGB: A hex string in the format "RRGGBB" (e.g., "FF0000").
//   - HRGBA: A hex string in the format "#RRGGBBAA" (e.g., "#FF0000FF").
//   - RGBA: A hex string in the format "RRGGBBAA" (e.g., "FF0000FF").
//   - HSL: A string in the format "hsl(H, S%, L%)" (e.g., "hsl(0, 100%, 50%)").
//   - HSLA: A string in the format "hsla(H, S%, L%, A)" (e.g., "hsla(0, 100%, 50%, 1)").
//   - CSSRGB: A string in the format "rgb(R, G, B)" (e.g., "rgb(255, 0, 0)").
//   - CSSRGBA: A string in the format "rgba(R, G, B, A)" (e.g., "rgba(255, 0, 0, 1)").
//   - CSS4RGB: A string in the format "rgb(R G B)" (e.g., "rgb(255 0 0)").
//   - CSS4RGBA: A string in the format "rgba(R G B / A)" (e.g., "rgba(255 0 0 / 1)").
//
// ​
//
// # Returns
//
// A string representation of the color in the specified format.
func (this Color) Format(format Format) string {
	switch format {
	case INT:
		return fmt.Sprintf("%d", uint32(this))
	case HARGB:
		return fmt.Sprintf("#%08X", uint32(this))
	case ARGB:
		return fmt.Sprintf("%08X", uint32(this))
	case HRGB:
		return fmt.Sprintf("#%06X", uint32(this)&0x00FFFFFF)
	case RGB:
		return fmt.Sprintf("%06X", uint32(this)&0x00FFFFFF)
	case HRGBA:
		return fmt.Sprintf("#%08X", (uint32(this)&0x00FFFFFF<<8)|(uint32(this)&0xFF000000>>24))
	case RGBA:
		return fmt.Sprintf("%08X", (uint32(this)&0x00FFFFFF<<8)|(uint32(this)&0xFF000000>>24))
	case HSL:
		_, red, green, blue := this.unpack()
		hue, saturation, lightness := RGBToHSL(red, green, blue)
		return fmt.Sprintf("hsl(%s, %s%%, %s%%)", strconv.FormatFloat(hue, 'f', -1, 64), strconv.FormatFloat(saturation*100, 'f', -1, 64), strconv.FormatFloat(lightness*100, 'f', -1, 64))
	case HSLA:
		alpha, red, green, blue := this.unpack()
		hue, saturation, lightness := RGBToHSL(red, green, blue)
		return fmt.Sprintf("hsla(%s, %s%%, %s%%, %s)", strconv.FormatFloat(hue, 'f', -1, 64), strconv.FormatFloat(saturation*100, 'f', -1, 64), strconv.FormatFloat(lightness*100, 'f', -1, 64), strconv.FormatFloat(float64(alpha)/255, 'f', -1, 64))
	case CSSRGB:
		_, red, green, blue := this.unpack()
		return fmt.Sprintf("rgb(%d, %d, %d)", red, green, blue)
	case CSSRGBA:
		alpha, red, green, blue := this.unpack()
		return fmt.Sprintf("rgba(%d, %d, %d, %.2g)", red, green, blue, float64(alpha)/255)
	case CSS4RGB:
		_, red, green, blue := this.unpack()
		return fmt.Sprintf("rgb(%d %d %d)", red, green, blue)
	case CSS4RGBA:
		alpha, red, green, blue := this.unpack()
		return fmt.Sprintf("rgba(%d %d %d / %.2g)", red, green, blue, float64(alpha)/255)
	default:
		return this.String()
	}
}
