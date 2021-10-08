package lib

import (
	"testing"

	"github.com/raji802/gotraining/Testing/lib"
)

func TestPow(t *testing.T) {

	tt := []struct {
		name    string
		x, y, v int
	}{
		{"two power three", 2, 3, 8},
		{"two power four", 2, 4, 16},
		{"two power five", 2, 5, 32},
	}

	for _, t1 := range tt {

		t.Run(t1.name, func(t *testing.T) {

			v := lib.Pow(t1.x, t1.y)
			if v != t1.v {

				t.Errorf("Expected %d,Got %d", t1.v, v)
			}

		})

	}

}
