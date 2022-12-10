package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYorN(t *testing.T) {
	assert.Equal(t, "Yes", YorN(true))
	assert.Equal(t, "No", YorN(false))
}
