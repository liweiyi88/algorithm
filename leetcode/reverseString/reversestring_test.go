package reversestring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	assert := assert.New(t)

	s := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(s)

	fmt.Println(string(s))

	assert.Equal([]byte{'h', 'a', 'n', 'n', 'a', 'H'}, s)
}
