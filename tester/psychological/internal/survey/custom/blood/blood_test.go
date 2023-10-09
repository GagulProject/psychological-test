package blood

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBloodType(t *testing.T) {
	bloodType := GetBloodType([]string{"A", "B"})
	assert.Equal(t, string(AB), bloodType)

	bloodType = GetBloodType([]string{"B", "A"})
	assert.Equal(t, string(AB), bloodType)

	bloodType = GetBloodType([]string{"A", "O"})
	assert.Equal(t, string(A), bloodType)

	bloodType = GetBloodType([]string{"O", "A"})
	assert.Equal(t, string(A), bloodType)

	bloodType = GetBloodType([]string{"O", "O"})
	assert.Equal(t, string(O), bloodType)

	bloodType = GetBloodType([]string{"G", "A"})
	assert.Equal(t, string(UNKNOWN), bloodType)
}
