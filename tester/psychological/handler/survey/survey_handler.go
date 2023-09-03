package survey

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"tester/psychological/tester/psychological/internal/analyzer"
	"tester/psychological/tester/psychological/view/html"
)

func GetSurveyResult(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// TODO: parse surveyType
	surveyInput := request.QueryStringParameters["survey_input"]
	surveyType := request.QueryStringParameters["survey_type"]

	surveyResult := analyzer.GetAnalyzeResult(surveyType, surveyInput)

	p := html.ResultParams{
		Title:   "설문 결과",
		Message: fmt.Sprintf("당신은 %s 타입입니다!", surveyResult),
	}

	output, err := html.Result(p)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, err
	}
	return events.APIGatewayProxyResponse{Headers: map[string]string{"content-type": "text/html"}, Body: *output, StatusCode: 200}, nil
}
