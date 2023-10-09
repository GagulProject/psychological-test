package analyzer

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"tester/psychological/tester/psychological/internal/survey/database/bidirection_multiple_choice_summation/mbti"
	"testing"
)

func TestGetAnalyseResult(T *testing.T) {
	surveyType := "mbti"
	surveyInputLength := len(mbti.GetOptions())
	surveyInput := strings.Repeat("Y", surveyInputLength)
	result := GetAnalyzeResult(surveyType, surveyInput)
	assert.Equal(T, "ESFJ", result)

	surveyType = "blood"
	surveyInput = "AB"
	result = GetAnalyzeResult(surveyType, surveyInput)
	assert.Equal(T, "AB", result)

	surveyType = "akgj"
	result = GetAnalyzeResult(surveyType, surveyInput)
	assert.Equal(T, "Wrong Type", result)
}

func TestParseInput(T *testing.T) {
	surveyAnswer := parseInput("1234")
	assert.Equal(T, []string{"1", "2", "3", "4"}, surveyAnswer)

	surveyAnswer = parseInput("wioeyr29")
	assert.Equal(T, []string{"w", "i", "o", "e", "y", "r", "2", "9"}, surveyAnswer)
}
