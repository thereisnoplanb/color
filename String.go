package color

import "fmt"

// Returns a string representation of the color in the standard format "AARRGGBB".
//
// # Returns
//
//	text string
//
// A string representation of the color in the format "AARRGGBB", where:
//
//  - AA is the alpha component,
//  - RR is the red component,
//  - GG is the green component,
//  - BB is the blue component,
// all represented as two-digit hexadecimal values.
func (this Color) String() (text string) {
	return fmt.Sprintf("%08X", uint32(this))
}
