package analysis

import (
	"testing"
)

func TestanalysisAlgo(s testing.F, b testing.B) string {
	surveyType := "blood"
	surveyAnswer := []string{"1", "2", "3"}
	analysisAlgo(surveyType, surveyAnswer)

	assert.Nil(s.T())
	assert.NotNil(s.T(), *output)
}

func TestsplitString(T testing.T) []string {
	surveyData := make([]string, len(surveyInput))

	for i, r := range surveyInput {
		surveyData[i] = string(r)
	}

	return surveyData
}
