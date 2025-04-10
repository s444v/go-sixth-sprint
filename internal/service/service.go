package service

import (
	"fmt"
	"regexp"
)

func stringDetector(s string) string {
	if check := regexp.MustCompile(`.-`); check.MatchString(s) {
		fmt.Println(true)
	}
	fmt.Println(false)
}
