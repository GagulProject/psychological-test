package analysis

import (
	"tester/psychological/tester/psychological/internal/survey/custom/blood"
	"tester/psychological/tester/psychological/internal/survey/custom/mbti"
)

func GetAnalysisResult(surveyType string, surveyInput string) string {
	return analysisAlgo(surveyType, splitString(surveyInput))
}

func analysisAlgo(surveyType string, surveyAnswer []string) string {
	var analysisResult string

	switch surveyType {
	case "blood":
		analysisResult = blood.GetBloodType(surveyAnswer)
	case "mbti":
		analysisResult = mbti.GetMBTIType(surveyAnswer)
	default:
		analysisResult = "Wrong Type"
	}
	/*
		typeFunc := map[string]interface{}{
			"blood": blood.GetBloodType(surveyAnswer),
			"mbti":  mbti.GetMBTIType(surveyAnswer),
		}
	*/
	return analysisResult
}

func splitString(surveyInput string) []string {
	surveyData := make([]string, len(surveyInput))

	for i, r := range surveyInput {
		surveyData[i] = string(r)
	}

	return surveyData
}
