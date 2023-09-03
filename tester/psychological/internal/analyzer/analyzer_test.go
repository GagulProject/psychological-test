package analyzer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAnalyseResult(T *testing.T) {
	surveyType := "mbti"
	surveyInput := "1234"
	result := GetAnalyzeResult(surveyType, surveyInput)
	assert.Equal(T, "INFJ", result)

	surveyType = "blood"
	surveyInput = "AB"
	result = GetAnalyzeResult(surveyType, surveyInput)
	assert.Equal(T, "AB", result)

	surveyType = "akgj"
	result = GetAnalyzeResult(surveyType, surveyInput)
	assert.Equal(T, "Wrong Type", result)
}

func TestAnalye(T *testing.T) {
	surveyType := "mbti"
	surveyInput := []string{"A", "B", "C"}
	result := Analyze(surveyType, surveyInput)
	assert.Equal(T, "INFJ", result)

	surveyType = "blood"
	surveyInput = []string{"A", "B"}
	result = Analyze(surveyType, surveyInput)
	assert.Equal(T, "AB", result)

	surveyType = "akgj"
	result = Analyze(surveyType, surveyInput)
	assert.Equal(T, "Wrong Type", result)
}

func TestparseInput(T *testing.T) {
	surveyAnswer := parseInput("1234")
	assert.Equal(T, []string{"1", "2", "3", "4"}, surveyAnswer)

	surveyAnswer = parseInput("wioeyr29")
	assert.Equal(T, []string{"w", "i", "o", "e", "y", "r", "2", "9"}, surveyAnswer)
}
