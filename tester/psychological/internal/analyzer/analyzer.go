package analyzer

import (
	"tester/psychological/tester/psychological/internal/survey/custom/blood"
	"tester/psychological/tester/psychological/internal/survey/custom/mbti"
)

func GetAnalyzeResult(surveyType string, surveyInput string) string {
	parseInputs := parseInput(surveyInput)

	return analyze(surveyType, parseInputs)
}

func analyze(surveyType string, surveyAnswer []string) string {
	var analysisResult string

	switch surveyType {
	case "blood":
		analysisResult = blood.GetBloodType(surveyAnswer)
	case "mbti":
		analysisResult = mbti.GetType(surveyAnswer)
	default:
		analysisResult = "Wrong Type"
	}

	//TODO: 추후 swtich 문이 아닌 Map 형식으로 변환예정
	return analysisResult
}

func parseInput(surveyInput string) []string {
	surveyData := make([]string, len(surveyInput))

	for i, r := range surveyInput {
		surveyData[i] = string(r)
	}

	return surveyData
}
