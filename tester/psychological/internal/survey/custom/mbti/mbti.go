package mbti

import (
	"errors"
)

// GetBloodType 두 문자 를 정렬하여 BloodType 정보를 리턴한다.
func GetMBTIType(surveyInput []string) (string, error) {

	return "", errors.New("can't find blood type")
}

type MBTIType string

const (
	INFJ MBTIType = "A"
	B    MBTIType = "B"
	AB   MBTIType = "AB"
	O    MBTIType = "O"
)
