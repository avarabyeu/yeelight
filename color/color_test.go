package color

import (
	"testing"
)

// TestHexStringToRgbInt passes valid and invalid values to the methods and make assertions.
func TestHexStringToRgbInt(t *testing.T) {
	_, err := HexStringToRgbInt("invalid format")
	if err == nil {
		t.Errorf("Invalid format was passed, but it returned no error.")
	}

	value, err := HexStringToRgbInt("#112233")
	if err != nil {
		t.Errorf("Valid format was passed, but error occured.")
	}

	// 256 * 256 * 17 (R) + 256 * 34 (G) + 51 (B) = 1122867
	if value != 1122867 {
		t.Errorf("Calculated int value is incorrect")
	}
}
