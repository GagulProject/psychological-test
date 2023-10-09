package bidirection_multiple_choice_summation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetType(t *testing.T) {
	options, _ := getOptions("mbti")
	var answers []string
	for i, v := range options {
		answers[i] = v.YesValue
	}
	typ, err := GetType("mbti", answers)
	assert.NoError(t, err)
	assert.Equal(t, "INFJ", typ)
}
