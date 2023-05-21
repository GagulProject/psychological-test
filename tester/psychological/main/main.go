package main

import (
	runtime "github.com/aws/aws-lambda-go/lambda"
	"tester/psychological/tester/psychological/handler/survey"
)

func main() {
	runtime.Start(survey.GetSurveyResult)
}
