package advent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GCD(t *testing.T) {
	assert.Equal(t, 5, GCD(10, 45))
	assert.Equal(t, 3, GCD(1701, 3768))
}
