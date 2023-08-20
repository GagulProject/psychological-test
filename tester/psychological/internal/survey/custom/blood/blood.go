package blood

import (
	"sort"
)

// GetBloodType 두 문자 를 정렬하여 BloodType 정보를 리턴한다.
func GetBloodType(chromosome []string) string {
	sort.Strings(chromosome)

	if chromosome[0] == "O" && chromosome[1] == "O" {
		return "O"
	}
	if chromosome[0] == "A" {
		if chromosome[1] == "B" {
			return "AB"
		}
		if chromosome[1] == "A " || chromosome[1] == "O" {
			return "A"
		}
	}
	if chromosome[0] == "B" {
		if chromosome[1] == "B " || chromosome[1] == "O" {
			return "B"
		}
	}

	return ""
}

type BloodType string

const (
	A  BloodType = "A"
	B  BloodType = "B"
	AB BloodType = "AB"
	O  BloodType = "O"
)
