package romantointeger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInteger(t *testing.T) {
	assert.Equal(t, 3, romanToInt("III"))
	assert.Equal(t, 58, romanToInt("LVIII"))
	assert.Equal(t, 1994, romanToInt("MCMXCIV"))
}
