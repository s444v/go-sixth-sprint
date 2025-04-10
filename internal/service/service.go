package service

import (
	"regexp"

	"github.com/s444v/go-sixth-sprint/pkg/morse"
)

func Detector(s string) string {
	if check := regexp.MustCompile(`^[ .-]+$`); check.MatchString(s) {
		return morse.ToText(s)
	}
	return morse.ToMorse(s)
}
