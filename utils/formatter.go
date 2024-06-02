package utils

import (
	"strconv"
	"strings"
)

func FormatTitleFromFirebase(title string) string {
	// remove the spacing on the title
	return strings.ReplaceAll(title, " ", "-")
}

func StringToUint(s string) (uint, error) {
	num, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(num), nil
}
