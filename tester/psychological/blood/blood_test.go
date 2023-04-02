package blood

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBloodType(t *testing.T) {
	bloodType, err := GetBloodType("A", "B")
	assert.Nil(t, err)
	assert.Equal(t, AB, bloodType)

	bloodType, err = GetBloodType("B", "A")
	assert.Nil(t, err)
	assert.Equal(t, AB, bloodType)

	bloodType, err = GetBloodType("A", "O")
	assert.Nil(t, err)
	assert.Equal(t, A, bloodType)

	bloodType, err = GetBloodType("O", "A")
	assert.Nil(t, err)
	assert.Equal(t, A, bloodType)

	bloodType, err = GetBloodType("O", "O")
	assert.Equal(t, O, bloodType)

	bloodType, err = GetBloodType("G", "A")
	assert.NotNil(t, err)
	assert.EqualError(t, err, "can't find blood type")
}
