package color

import (
	"fmt"
	"strconv"
	"strings"
)

var formatLength = map[Format]int{
	HARGB: 9,
	ARGB:  8,
	HRGB:  7,
	RGB:   6,
	HRGBA: 9,
	RGBA:  8,
}

// Parse takes a string representation of a color in the specified format and returns a Color object.
// If the input string is not in a valid format, it returns an error.
//
// # Parameters
//
//	layout Format
//
// The format of the input color string. Supported formats include:
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
//	value string
//
// The string representation of the color to be parsed.
//
// # Returns
//
//	color Color
//
// The parsed Color object.
//
//	err error
//
// An error if the input string is not in a valid format.
func Parse(layout Format, value string) (color Color, err error) {
	if lenght, ok := formatLength[layout]; ok {
		if len(value) != lenght {
			return 0, fmt.Errorf("invalid color length")
		}
	}
	switch layout {
	case INT:
		colorInt, err := strconv.ParseUint(value, 10, 32)
		return Color(colorInt), err
	case HRGB:
		if value[0] != '#' {
			return 0, fmt.Errorf("invalid value")
		}
		value = value[1:]
		fallthrough
	case RGB:
		colorInt, err := strconv.ParseUint(value, 16, 24)
		return Color(colorInt | 0xFF000000), err
	case HRGBA:
		if value[0] != '#' {
			return 0, fmt.Errorf("invalid value")
		}
		value = value[1:]
		fallthrough
	case RGBA:
		colorInt, err := strconv.ParseUint(value, 16, 32)
		return Color((colorInt >> 8) | (colorInt << 24)), err
	case HARGB:
		if value[0] != '#' {
			return 0, fmt.Errorf("invalid value")
		}
		value = value[1:]
		fallthrough
	case ARGB:
		colorInt, err := strconv.ParseUint(value, 16, 32)
		return Color(colorInt), err
	case HSL:
		if strings.HasPrefix(value, "hsl(") && strings.HasSuffix(value, ")") {
			value = value[4 : len(value)-1]
			parts := strings.Split(value, ",")
			if len(parts) != 3 {
				return 0, fmt.Errorf("invalid HSL format")
			}
			hue, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
			if err != nil {
				return 0, fmt.Errorf("invalid hue value: %v", err)
			}
			saturationStr := strings.TrimSpace(parts[1])
			if !strings.HasSuffix(saturationStr, "%") {
				return 0, fmt.Errorf("saturation must end with %%")
			}
			saturation, err := strconv.ParseFloat(strings.TrimSuffix(saturationStr, "%"), 64)
			if err != nil {
				return 0, fmt.Errorf("invalid saturation value: %v", err)
			}
			lightnessStr := strings.TrimSpace(parts[2])
			if !strings.HasSuffix(lightnessStr, "%") {
				return 0, fmt.Errorf("lightness must end with %%")
			}
			lightness, err := strconv.ParseFloat(strings.TrimSuffix(lightnessStr, "%"), 64)
			if err != nil {
				return 0, fmt.Errorf("invalid lightness value: %v", err)
			}
			return FromHSL(hue, saturation/100, lightness/100)
		}
		return 0, fmt.Errorf("invalid HSL format")
	case HSLA:
		if strings.HasPrefix(value, "hsla(") && strings.HasSuffix(value, ")") {
			value = value[5 : len(value)-1]
			parts := strings.Split(value, ",")
			if len(parts) != 4 {
				return 0, fmt.Errorf("invalid HSLA format")
			}
			hue, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
			if err != nil {
				return 0, fmt.Errorf("invalid hue value: %v", err)
			}
			saturationStr := strings.TrimSpace(parts[1])
			if !strings.HasSuffix(saturationStr, "%") {
				return 0, fmt.Errorf("saturation must end with %%")
			}
			saturation, err := strconv.ParseFloat(strings.TrimSuffix(saturationStr, "%"), 64)
			if err != nil {
				return 0, fmt.Errorf("invalid saturation value: %v", err)
			}
			lightnessStr := strings.TrimSpace(parts[2])
			if !strings.HasSuffix(lightnessStr, "%") {
				return 0, fmt.Errorf("lightness must end with %%")
			}
			lightness, err := strconv.ParseFloat(strings.TrimSuffix(lightnessStr, "%"), 64)
			if err != nil {
				return 0, fmt.Errorf("invalid lightness value: %v", err)
			}
			alpha, err := strconv.ParseFloat(strings.TrimSpace(parts[3]), 64)
			if err != nil {
				return 0, fmt.Errorf("invalid alpha value: %v", err)
			}
			return FromHSLA(hue, saturation/100, lightness/100, uint8(alpha*255))
		}
		return 0, fmt.Errorf("invalid HSLA format")
	case CSSRGB:
		if strings.HasPrefix(value, "rgb(") && strings.HasSuffix(value, ")") {
			value = value[4 : len(value)-1]
			parts := strings.Split(value, ",")
			if len(parts) != 3 {
				return 0, fmt.Errorf("invalid CSS RGB format")
			}
			red, err := strconv.ParseUint(strings.TrimSpace(parts[0]), 10, 8)
			if err != nil {
				return 0, fmt.Errorf("invalid red value: %v", err)
			}
			green, err := strconv.ParseUint(strings.TrimSpace(parts[1]), 10, 8)
			if err != nil {
				return 0, fmt.Errorf("invalid green value: %v", err)
			}
			blue, err := strconv.ParseUint(strings.TrimSpace(parts[2]), 10, 8)
			if err != nil {
				return 0, fmt.Errorf("invalid blue value: %v", err)
			}
			return FromRGB(uint8(red), uint8(green), uint8(blue)), nil
		}
		return 0, fmt.Errorf("invalid CSS RGB format")
	case CSSRGBA:
		if strings.HasPrefix(value, "rgba(") && strings.HasSuffix(value, ")") {
			value = value[5 : len(value)-1]
			parts := strings.Split(value, ",")
			if len(parts) != 4 {
				return 0, fmt.Errorf("invalid CSS RGBA format")
			}
			red, err := strconv.ParseUint(strings.TrimSpace(parts[0]), 10, 8)
			if err != nil {
				return 0, fmt.Errorf("invalid red value: %v", err)
			}
			green, err := strconv.ParseUint(strings.TrimSpace(parts[1]), 10, 8)
			if err != nil {
				return 0, fmt.Errorf("invalid green value: %v", err)
			}
			blue, err := strconv.ParseUint(strings.TrimSpace(parts[2]), 10, 8)
			if err != nil {
				return 0, fmt.Errorf("invalid blue value: %v", err)
			}
			alphaStr := strings.TrimSpace(parts[3])
			percent := false
			if strings.HasSuffix(alphaStr, "%") {
				percent = true
				alphaStr = alphaStr[:len(alphaStr)-1]
			}
			alpha, err := strconv.ParseFloat(alphaStr, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid alpha value: %v", err)
			}
			if percent {
				alpha = alpha / 100
			}
			if alpha < 0 || alpha > 1 {
				return 0, fmt.Errorf("alpha value must be between 0 and 1, or between 0%% and 100%%")
			}
			return FromARGB(uint8(alpha*255+0.5), uint8(red), uint8(green), uint8(blue)), nil
		}
		return 0, fmt.Errorf("invalid CSS RGBA format")
	case CSS4RGB:
		if strings.HasPrefix(value, "rgb(") && strings.HasSuffix(value, ")") {
			value = value[4 : len(value)-1]
			parts := strings.Fields(value)
			if len(parts) != 3 {
				return 0, fmt.Errorf("invalid CSS4 RGB format")
			}
			red, err := strconv.ParseUint(strings.TrimSpace(parts[0]), 10, 8)
			if err != nil {
				return 0, fmt.Errorf("invalid red value: %v", err)
			}
			green, err := strconv.ParseUint(strings.TrimSpace(parts[1]), 10, 8)
			if err != nil {
				return 0, fmt.Errorf("invalid green value: %v", err)
			}
			blue, err := strconv.ParseUint(strings.TrimSpace(parts[2]), 10, 8)
			if err != nil {
				return 0, fmt.Errorf("invalid blue value: %v", err)
			}
			return FromRGB(uint8(red), uint8(green), uint8(blue)), nil
		}
		return 0, fmt.Errorf("invalid CSS4 RGB format")
	case CSS4RGBA:
		if strings.HasPrefix(value, "rgba(") && strings.HasSuffix(value, ")") {
			value = value[5 : len(value)-1]
			parts := strings.Fields(value)
			if len(parts) != 5 || parts[3] != "/" {
				return 0, fmt.Errorf("invalid CSS4 RGBA format")
			}
			red, err := strconv.ParseUint(strings.TrimSpace(parts[0]), 10, 8)
			if err != nil {
				return 0, fmt.Errorf("invalid red value: %v", err)
			}
			green, err := strconv.ParseUint(strings.TrimSpace(parts[1]), 10, 8)
			if err != nil {
				return 0, fmt.Errorf("invalid green value: %v", err)
			}
			blue, err := strconv.ParseUint(strings.TrimSpace(parts[2]), 10, 8)
			if err != nil {
				return 0, fmt.Errorf("invalid blue value: %v", err)
			}
			alphaStr := strings.TrimSpace(parts[4])
			percent := false
			if strings.HasSuffix(alphaStr, "%") {
				percent = true
				alphaStr = alphaStr[:len(alphaStr)-1]
			}
			alpha, err := strconv.ParseFloat(alphaStr, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid alpha value: %v", err)
			}
			if percent {
				alpha = alpha / 100
			}
			if alpha < 0 || alpha > 1 {
				return 0, fmt.Errorf("alpha value must be between 0 and 1 or between 0%% and 100%%")
			}
			return FromARGB(uint8(alpha*255+0.5), uint8(red), uint8(green), uint8(blue)), nil
		}
		return 0, fmt.Errorf("invalid CSS4 RGBA format")
	default:
		color, err := strconv.ParseUint(value, 16, 32)
		return Color(color | 0xFF000000), err
	}
}

// Parse takes a string representation of a color in the specified format and returns a Color object.
// If the input string is not in a valid format, it panics.
//
// # Parameters
//
//	layout Format
//
// The format of the input color string. Supported formats include:
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
//	value string
//
// The string representation of the color to be parsed.
//
// # Returns
//
//	color Color
//
// The parsed Color object.
func MustParse(layout Format, value string) (color Color) {
	color, err := Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return color
}
