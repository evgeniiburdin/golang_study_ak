package main

import (
	"strings"
	"testing"
)

func Test_generateActivationKey(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateActivationKey()
			keyParts := strings.Split(got, "-")
			if len(keyParts) != 4 {
				t.Errorf("generateActivationKey() got = %v. key must suit format XXXX-XXXX-XXXX-XXXX", got)
			}
			for _, keyPart := range keyParts {
				for _, b := range keyPart {
					if !(b >= 'A' && b <= 'Z' || b >= '0' && b <= '9') {
						t.Errorf(`generateActivationKey() got = %v. "%v": key must consist either of ASCII chars or nums`, got, string(b))
					}
				}
			}

		})
	}
}
