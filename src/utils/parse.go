package utils

import (
	"regexp"
	"strings"

	"github.com/chromedp/cdproto/cdp"
	"github.com/ealpizr/tse-api/src/models"
)

func ParseFullName(fullName string) (firstName string, fLastName string, sLastName string) {
	s := strings.Split(fullName, " ")

	// Support for middle names
	if len(s) > 3 {
		sLastName = s[len(s)-1]
		fLastName = s[len(s)-2]
		firstName = strings.Join(s[:len(s)-2], " ")
		return
	}

	firstName = s[0]
	fLastName = s[1]

	// Second last name is optional
	if len(s) == 3 {
		sLastName = s[2]
	}
	return
}

func ParseFoundResults(nodes []*cdp.Node) []models.Record {
	r := []models.Record{}
	nameRegex := regexp.MustCompile("[ÑA-Z]+[ ÑA-Z]+")
	idRegex := regexp.MustCompile("[0-9]{9}")
	for _, n := range nodes {
		for _, t := range n.Children {
			id := idRegex.FindString(t.NodeValue)
			fullName := nameRegex.FindString(t.NodeValue)
			// Skip invisible records
			if id != "" && fullName != "" {
				r = append(r, models.Record{ID: id, FullName: strings.TrimSpace(fullName)})
			}
		}
	}
	return r
}
