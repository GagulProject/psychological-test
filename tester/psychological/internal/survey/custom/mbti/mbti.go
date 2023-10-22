package mbti

import (
	"tester/psychological/tester/psychological/internal/survey/collector/bidirection_multiple_choice_summation"
)

// GetBloodType 두 문자 를 정렬하여 BloodType 정보를 리턴한다.
func GetType(surveyInput []string) string {

	typ, err := bidirection_multiple_choice_summation.GetType("mbti", surveyInput)
	if err != nil {
		return ""
	}
	return typ
}
