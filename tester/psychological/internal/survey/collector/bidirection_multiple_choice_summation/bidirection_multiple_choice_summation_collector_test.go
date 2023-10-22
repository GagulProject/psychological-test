package bidirection_multiple_choice_summation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetType(t *testing.T) {
	options, _ := getOptions("mbti")
	answers := make([]string, len(options))
	for i := range options {
		answers[i] = "N"
	}
	typ, err := GetType("mbti", answers)
	assert.NoError(t, err)
	assert.Equal(t, "INTP", typ)
}
