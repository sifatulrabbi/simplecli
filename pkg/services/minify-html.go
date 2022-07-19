package services

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/sifatulrabbi/simplecli/pkg/constants"
)

const FILE_PERMISSION = 0644

// Minify HTML function
func MinifyHTML(path string, oPath string) {
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
	if oPath != "" {
		writePath = oPath
	} else {
		oPath := strings.Split(path, "/")
		sFilename := strings.Split(oPath[len(oPath)-1], ".")
		writePath = "./" + strings.Join(sFilename[:len(sFilename)-1], "") + ".minified" + "." + sFilename[len(sFilename)-1]
	}
	if err := ioutil.WriteFile(writePath, []byte(minified), FILE_PERMISSION); err != nil {
		panic(constants.UNABLE_ERROR)
	}
}
