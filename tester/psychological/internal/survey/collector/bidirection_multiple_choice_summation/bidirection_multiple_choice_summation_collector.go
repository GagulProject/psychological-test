package bidirection_multiple_choice_summation

import (
	"errors"
	"tester/psychological/tester/psychological/internal/survey/database/bidirection_multiple_choice_summation/mbti"
	bmcs "tester/psychological/tester/psychological/internal/survey/model/bidirection_multiple_choice_summation"
)

const YesValue = "Y"
const NoValue = "N"

func GetType(dataName string, answer []string) (string, error) {
	options, err := getOptions(dataName)
	if err != nil {
		return "", err
	}

	answers := make(map[string]int)

	for i := range options {
		if answer[i] == YesValue {
			answers[options[i].YesValue]++
		} else {
			answers[options[i].NoValue]++
		}
	}

	values, err := getValues(dataName)
	if err != nil {
		return "", err
	}

	var typ string

	for i := range values {
		v1 := values[i][0]
		v2 := values[i][1]
		if answers[v1] > answers[v2] {
			typ += v1
		} else {
			typ += v2
		}
	}
	return typ, nil
}

func getOptions(dataName string) ([]bmcs.Option, error) {
	switch dataName {
	case "mbti":
		return mbti.GetOptions(), nil
	}
	return nil, errors.New("not found")
}

func getValues(dataName string) ([][]string, error) {
	switch dataName {
	case "mbti":
		return mbti.GetValues(), nil
	}

	return nil, errors.New("not found")
}
