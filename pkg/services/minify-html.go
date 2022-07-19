package services

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/sifatulrabbi/simplecli/pkg/constants"
	"github.com/sifatulrabbi/simplecli/pkg/utils"
)

// Minify HTML function
func MinifyHTML(path string, outPath string) {
	content, err := ioutil.ReadFile(path)
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
	if err := ioutil.WriteFile(
		writePath,
		[]byte(minified),
		constants.FILE_MODE,
	); err != nil {
		panic(constants.UNABLE_ERROR)
	}
}
