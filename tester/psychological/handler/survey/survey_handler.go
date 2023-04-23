package survey

import (
	"fmt"
	"tester/psychological/tester/psychological/view/html"
)

func GetSurveyResult(surveyType string, surveyInput string) {
	// TODO: parse surveyInput
	p := html.ResultParams{
		Title:   "설문 결과",
		Message: fmt.Sprintf("당신은 %s 타입입니다!", surveyInput),
	}
	html.Result(p)
}
