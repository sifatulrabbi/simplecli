package services

import (
	"io/ioutil"
	"regexp"

	"github.com/sifatulrabbi/simplecli/pkg/constants"
	"github.com/sifatulrabbi/simplecli/pkg/utils"
)

func MinifyCSS(path string, outPath string) {
	// Read file
	content, err := ioutil.ReadFile(path)
	if err != nil {
		// If the file is not found
		panic(constants.INCORRECT_ARGS)
	}
	// Convert bytes into strings
	css := string(content)
	// Minify the strings
	minified := regexp.MustCompile(`(//)+.*`).ReplaceAllString(css, "")
	minified = regexp.MustCompile(`\n`).ReplaceAllString(minified, " ")
	minified = regexp.MustCompile("  ").ReplaceAllString(minified, "")
	minified = regexp.MustCompile("; ").ReplaceAllString(minified, ";")
	minified = regexp.MustCompile(": ").ReplaceAllString(minified, ":")
	minified = regexp.MustCompile(" } ").ReplaceAllString(minified, "}")
	minified = regexp.MustCompile(" { ").ReplaceAllString(minified, "{")
	minified = regexp.MustCompile(", ").ReplaceAllString(minified, ",")
	// Flush minified string in to the output file
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
