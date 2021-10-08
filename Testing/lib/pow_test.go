package lib

import "testing"

func TestPow(t *testing.T) {

	t.Run("two power three", func(t *testing.T) {

		v := Pow(2, 3)
		if v != 8 {

			t.Errorf("Expected %d,Got %d", 8, v)
		}
	})

	t.Run("two power four", func(t *testing.T) {

		v := Pow(2, 4)
		if v != 16 {

			t.Errorf("Expected %d,Got %d", 16, v)
		}
	})
}
