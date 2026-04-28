package color

import (
	"database/sql/driver"
	"fmt"
)

// Implements the driver.Value interface.
func (this Color) Value() (driver.Value, error) {
	return this.Format(HRGBA), nil
}

// Implements the sql.Scanner interface.
func (this *Color) Scan(src any) error {
	switch value := src.(type) {
	case string:
		parsedColor, err := Parse(HRGBA, value)
		if err != nil {
			return err
		}
		*this = parsedColor
	case []byte:
		parsedColor, err := Parse(HRGBA, string(value))
		if err != nil {
			return err
		}
		*this = parsedColor
	default:
		return fmt.Errorf("unsupported type: %T", src)
	}
	return nil
}
