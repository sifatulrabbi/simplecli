package utils

import (
	"strings"
)

func GenOutputPath(path string) string {
	outPath := strings.Split(path, "/")
	sFilename := strings.Split(outPath[len(outPath)-1], ".")
	writePath := "./" + strings.Join(sFilename[:len(sFilename)-1], "") + ".minified" + "." + sFilename[len(sFilename)-1]
	return writePath
}
