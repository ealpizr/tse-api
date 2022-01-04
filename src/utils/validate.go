package utils

import (
	"regexp"
	"strings"
)

func IsIDValid(id string) bool {
	idCheck := regexp.MustCompile("^[0-9]{9}$")
	return idCheck.MatchString(id)
}

func IsFullNameValid(fullName string) bool {
	s := strings.Split(fullName, " ")
	for _, v := range s {
		if len(v) < 3 {
			return false
		}
	}
	return len(s) >= 2
}
