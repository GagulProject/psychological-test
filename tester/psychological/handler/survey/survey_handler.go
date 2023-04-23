package survey

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"tester/psychological/tester/psychological/view/html"
)

func GetSurveyResult(surveyType string, surveyInput string) events.APIGatewayProxyResponse {
	// TODO: parse surveyInput
	p := html.ResultParams{
		Title:   "설문 결과",
		Message: fmt.Sprintf("당신은 %s 타입입니다!", surveyInput),
	}

	output, err := html.Result(p)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}
	}
	return events.APIGatewayProxyResponse{Body: *output, StatusCode: 200}
}
