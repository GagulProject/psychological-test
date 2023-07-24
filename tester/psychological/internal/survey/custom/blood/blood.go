package blood

import (
	"errors"
	"sort"
)

// GetBloodType 두 문자 를 정렬하여 BloodType 정보를 리턴한다.
func GetBloodType(chromosome []string) (string, error) {
	sort.Strings(chromosome)

	if chromosome[0] == "O" && chromosome[1] == "O" {
		return "O", nil
	}
	if chromosome[0] == "A" {
		if chromosome[1] == "B" {
			return "AB", nil
		}
		if chromosome[1] == "A " || chromosome[1] == "O" {
			return "A", nil
		}
	}
	if chromosome[0] == "B" {
		if chromosome[1] == "B " || chromosome[1] == "O" {
			return "B", nil
		}
	}

	return "", errors.New("can't find blood type")
}

type BloodType string

const (
	A  BloodType = "A"
	B  BloodType = "B"
	AB BloodType = "AB"
	O  BloodType = "O"
)
