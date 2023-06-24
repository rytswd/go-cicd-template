package color

import (
	"testing"
)

func TestRGB(t *testing.T) {
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
