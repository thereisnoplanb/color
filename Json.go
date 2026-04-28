package color

import "encoding/json"

// MarshalJSON implements the json.Marshaler interface.
// It encodes the Color as a hex string in the format "#RRGGBBAA".
func (this Color) MarshalJSON() (data []byte, err error) {
	return json.Marshal(this.Format(HRGBA))
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// It decodes a hex string in the format "#RRGGBBAA" into a Color.
func (this *Color) UnmarshalJSON(data []byte) (err error) {
	var colorStr string
	if err := json.Unmarshal(data, &colorStr); err != nil {
		return err
	}
	color, err := Parse(HRGBA, colorStr)
	if err != nil {
		return err
	}

	*this = color
	return nil
}
