package analysis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAnalysisResult(T *testing.T) {
	surveyType := "mbti"
	surveyInput := "1234"
	result := GetAnalysisResult(surveyType, surveyInput)
	assert.Equal(T, "INFJ", result)

	surveyType = "blood"
	surveyInput = "AB"
	result = GetAnalysisResult(surveyType, surveyInput)
	assert.Equal(T, "AB", result)

	surveyType = "akgj"
	result = GetAnalysisResult(surveyType, surveyInput)
	assert.Equal(T, "Wrong Type", result)

}

func TestanalysisAlgo(T *testing.T) {
	surveyType := "mbti"
	surveyInput := []string{"A", "B", "C"}
	result := analysisAlgo(surveyType, surveyInput)
	assert.Equal(T, "INFJ", result)

	surveyType = "blood"
	surveyInput = []string{"A", "B"}
	result = analysisAlgo(surveyType, surveyInput)
	assert.Equal(T, "AB", result)

	surveyType = "akgj"
	result = analysisAlgo(surveyType, surveyInput)
	assert.Equal(T, "Wrong Type", result)
}

func TestsplitString(T *testing.T) {
	surveyAnswer := splitString("1234")
	assert.Equal(T, []string{"1", "2", "3", "4"}, surveyAnswer)

	surveyAnswer = splitString("wioeyr29")
	assert.Equal(T, []string{"w", "i", "o", "e", "y", "r", "2", "9"}, surveyAnswer)
}
