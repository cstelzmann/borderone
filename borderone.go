package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

func main() {
	var exeRegex = regexp.MustCompile(`borderone-([^\-]+)-(\d+x\d+)`)
	var filename = filepath.Base(os.Args[0])
	var imagefile = os.Args[1]

	if exeRegex.MatchString(filename) {
		var matches = exeRegex.FindStringSubmatch(filename)

		var color = matches[1]
		var size = matches[2]

		exec.Command("magick", "convert", imagefile, "-bordercolor", color, "-border", size, imagefile).Run()
	}
}
