package main

import (
	"bookingapp/helper"
	"unicode"
)

func isValidName(s string) bool {
	// this is how we access the exported constant from the helper package
	const ExportedConstant = helper.ExportedConstant
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
