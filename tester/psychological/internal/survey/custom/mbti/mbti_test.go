package mbti

import (
	"github.com/stretchr/testify/assert"
	"tester/psychological/tester/psychological/internal/survey/database/bidirection_multiple_choice_summation/mbti"
	"testing"
)

func TestGetType(t *testing.T) {
	answers := make([]string, len(mbti.GetOptions()))
	for i := range mbti.GetOptions() {
		answers[i] = "Y"
	}
	mbtiType := GetType(answers)
	assert.Equal(t, "ESFJ", mbtiType)
}
