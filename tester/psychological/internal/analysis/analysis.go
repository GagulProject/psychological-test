package analysis

func GetAnalysisResult(surveyType string, surveyInput string) string {
	return analysisAlgo(surveyType, splitString(surveyInput))
}

func analysisAlgo(surveyType string, surveyAnswer []string) string {
	GetBloodType()

	analysisResult := "착한"

	return analysisResult
}

func splitString(surveyInput string) []string {
	surveyData := make([]string, len(surveyInput))

	for i, r := range surveyInput {
		surveyData[i] = string(r)
	}

	return surveyData
}
