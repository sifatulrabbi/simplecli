package services

import (
	"fmt"
	"os"
	"regexp"

	"github.com/sifatulrabbi/simplecli/internal/constants"
	"github.com/sifatulrabbi/simplecli/internal/utils"
)

// Minify HTML function
func MinifyHTML(path string, outPath string) {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Print("Invalid file path.\n")
		panic(constants.INCORRECT_ARGS)
	}
	html := string(content)
	minified := regexp.MustCompile(`\n`).ReplaceAllString(html, " ")
	minified = regexp.MustCompile(`  `).ReplaceAllString(minified, "")
	minified = regexp.MustCompile(`; `).ReplaceAllString(minified, ";")
	minified = regexp.MustCompile(`: `).ReplaceAllString(minified, ":")
	minified = regexp.MustCompile(` <`).ReplaceAllString(minified, "<")
	minified = regexp.MustCompile(` >`).ReplaceAllString(minified, ">")
	minified = regexp.MustCompile(` />`).ReplaceAllString(minified, "/>")
	minified = regexp.MustCompile(`> <`).ReplaceAllString(minified, "><")

	var writePath string
	if outPath != "" {
		writePath = outPath
	} else {
		writePath = utils.GenOutputPath(path)
	}
	if err := os.WriteFile(
		writePath,
		[]byte(minified),
		constants.FILE_MODE,
	); err != nil {
		panic(constants.UNABLE_ERROR)
	}
}
