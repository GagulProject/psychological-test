package blood

import (
	"sort"
)

type BloodType string

const (
	A       BloodType = "A"
	B       BloodType = "B"
	AB      BloodType = "AB"
	O       BloodType = "O"
	UNKNOWN BloodType = ""
)

// GetBloodType 두 문자 를 정렬하여 BloodType 정보를 리턴한다.
func GetBloodType(chromosome []string) string {
	sort.Strings(chromosome)

	if chromosome[0] == string(O) && chromosome[1] == string(O) {
		return string(O)
	}
	if chromosome[0] == string(A) {
		if chromosome[1] == string(B) {
			return string(AB)
		}
		if chromosome[1] == string(A) || chromosome[1] == string(O) {
			return string(A)
		}
	}
	if chromosome[0] == string(B) {
		if chromosome[1] == string(B) || chromosome[1] == string(O) {
			return string(B)
		}
	}

	return string(UNKNOWN)
}
