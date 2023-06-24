package color

import (
	"testing"
)

func TestRGB(t *testing.T) {
	// var toHex float64 = 255

	cases := map[string]struct {
		R float64
		G float64
		B float64

		want string
	}{
		"Red": {
			R:    255,
			G:    0,
			B:    0,
			want: "ff0000",
		},
		"Green": {
			R:    0,
			G:    255,
			B:    0,
			want: "00ff00",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			rgb := RGB{R: tc.R, G: tc.G, B: tc.B}

			got := rgb.ToHex()
			if tc.want != got {
				t.Errorf("hex result didn't match, want '%v', got '%v'", tc.want, got)
			}
		})
	}
}

func TestHSL(t *testing.T) {
	cases := map[string]struct {
		H float64
		S float64
		L float64

		want string
	}{
		"Grey (808080)": {
			H:    0,
			S:    0,
			L:    0.5,
			want: "808080",
		},
		// "Thistle (DECCF5)": {
		// 	H:    266.0 / 360, // Hue is a degree value
		// 	S:    67.0 / 100,  // Saturation is a percentage
		// 	L:    88.0 / 100,  // Lightness is a percentage
		// 	want: "deccf5",
		// },
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			hsl := HSL{H: tc.H, S: tc.S, L: tc.L}

			got := hsl.ToRGB().ToHex()
			if tc.want != got {
				t.Errorf("hex result didn't match, want '%v', got '%v'", tc.want, got)
			}
		})
	}
}
