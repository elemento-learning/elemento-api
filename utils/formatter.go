package utils

import "strings"

func FormatTitleFromFirebase(title string) string {
	// remove the spacing on the title
	return strings.ReplaceAll(title, " ", "-")
}
