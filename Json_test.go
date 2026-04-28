package color_test

import (
	"encoding/json"
	"testing"

	"github.com/thereisnoplanb/color"
)

func TestColor_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		color   color.Color
		want    []byte
		wantErr bool
	}{
		{
			name:    "Test MarshalJSON with a Red color",
			color:   color.Red,
			want:    []byte("\"#FF0000FF\""),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := json.Marshal(tt.color)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("MarshalJSON() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("MarshalJSON() succeeded unexpectedly")
			}
			if string(got) != string(tt.want) {
				t.Errorf("MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_MarshalJSON_InStruct(t *testing.T) {
	type TestStruct struct {
		Name  string       `json:"name"`
		Color *color.Color `json:"color"`
	}
	getPointer := func(c color.Color) *color.Color {
		return &c
	}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		testStruct TestStruct
		want       []byte
		wantErr    bool
	}{
		{
			name: "Test MarshalJSON in struct with a Red color",
			testStruct: TestStruct{
				Name:  "Test",
				Color: getPointer(color.Red),
			},
			want:    []byte("{\"name\":\"Test\",\"color\":\"#FF0000FF\"}"),
			wantErr: false,
		},
		{
			name: "Test MarshalJSON in struct with a nil color",
			testStruct: TestStruct{
				Name:  "Test",
				Color: nil,
			},
			want:    []byte("{\"name\":\"Test\",\"color\":null}"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := json.Marshal(tt.testStruct)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("MarshalJSON() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("MarshalJSON() succeeded unexpectedly")
			}
			if string(got) != string(tt.want) {
				t.Errorf("MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data      []byte
		wantColor color.Color
		wantErr   bool
	}{
		{
			name:      "Test UnmarshalJSON with a valid color string",
			data:      []byte("\"#FF0000FF\""),
			wantColor: color.Red,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var color color.Color
			err := json.Unmarshal(tt.data, &color)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("UnmarshalJSON() failed: %v", err)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UnmarshalJSON() succeeded unexpectedly")
				return
			}
			if color != tt.wantColor {
				t.Errorf("UnmarshalJSON() = %v, want %v", color, tt.wantColor)
			}
		})
	}
}

func TestColor_UnmarshalJSON_InStruct(t *testing.T) {
	type TestStruct struct {
		Name  string       `json:"name"`
		Color *color.Color `json:"color"`
	}
	getPointer := func(c color.Color) *color.Color {
		return &c
	}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data       []byte
		wantStruct TestStruct
		wantErr    bool
	}{
		{
			name: "Test UnmarshalJSON in struct with a valid color string",
			data: []byte("{\"name\":\"Test\",\"color\":\"#FF0000FF\"}"),
			wantStruct: TestStruct{
				Name:  "Test",
				Color: getPointer(color.Red),
			},
			wantErr: false,
		},
		{
			name: "Test UnmarshalJSON in struct with a null color",
			data: []byte("{\"name\":\"Test\",\"color\":null}"),
			wantStruct: TestStruct{
				Name:  "Test",
				Color: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotStruct TestStruct
			err := json.Unmarshal(tt.data, &gotStruct)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("UnmarshalJSON() failed: %v", err)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UnmarshalJSON() succeeded unexpectedly")
				return
			}
			if gotStruct.Name != tt.wantStruct.Name {
				t.Errorf("UnmarshalJSON() Name = %v, want %v", gotStruct.Name, tt.wantStruct.Name)
			}
			if (gotStruct.Color == nil) != (tt.wantStruct.Color == nil) {
				t.Errorf("UnmarshalJSON() Color nil = %v, want nil = %v", gotStruct.Color == nil, tt.wantStruct.Color == nil)
			} else if gotStruct.Color != nil && *gotStruct.Color != *tt.wantStruct.Color {
				t.Errorf("UnmarshalJSON() Color = %v, want %v", *gotStruct.Color, *tt.wantStruct.Color)
			}
		})
	}
}

func TestColor_UnmarshalJSON_FromPointer(t *testing.T) {
	getPointer := func(c color.Color) *color.Color {
		return &c
	}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data      []byte
		wantColor *color.Color
		wantErr   bool
	}{
		{
			name:      "Test UnmarshalJSON in struct with a valid color string",
			data:      []byte("\"#FF0000FF\""),
			wantColor: getPointer(color.Red),
			wantErr:   false,
		},
		{
			name:      "Test UnmarshalJSON in struct with a null color",
			data:      []byte("null"),
			wantColor: nil,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotColor *color.Color
			err := json.Unmarshal(tt.data, &gotColor)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("UnmarshalJSON() failed: %v", err)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UnmarshalJSON() succeeded unexpectedly")
				return
			}
			if (gotColor == nil) != (tt.wantColor == nil) {
				t.Errorf("UnmarshalJSON() Color nil = %v, want nil = %v", gotColor == nil, tt.wantColor == nil)
			} else if gotColor != nil && *gotColor != *tt.wantColor {
				t.Errorf("UnmarshalJSON() Color = %v, want %v", *gotColor, *tt.wantColor)
			}
		})
	}
}
