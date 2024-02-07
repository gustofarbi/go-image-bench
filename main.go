package main

import (
	"fmt"
	"github.com/davidbyttow/govips/v2/vips"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

func main() {
	vipsHeader()
	govipsSize()
}

func govipsSize() {
	vips.LoggingSettings(func(string, vips.LogLevel, string) {}, vips.LogLevelDebug)
	start := time.Now()
	f, err := os.Open("testdata/1820MB-image.tif")
	panicOnError(err)
	im, err := vips.NewImageFromReader(f)
	panicOnError(err)
	defer im.Close()

	w, h := im.Width(), im.Height()
	fmt.Printf("govips %s:   width: %d, height: %d\n", time.Since(start), w, h)
}

func vipsHeader() {
	start := time.Now()
	f, err := os.Open("testdata/1820MB-image.tif")
	panicOnError(err)

	cmd := exec.Command("vipsheader", "stdin")
	cmd.Stdin = f

	re := regexp.MustCompile(`(\d+)x(\d+)`)
	b, err := cmd.Output()
	panicOnError(err)

	matches := re.FindAllStringSubmatch(string(b), -1)

	if len(matches[0]) != 3 {
		panic("no matches")
	}

	w, err := strconv.Atoi(matches[0][1])
	panicOnError(err)
	h, err := strconv.Atoi(matches[0][2])
	panicOnError(err)

	fmt.Printf("vipsheader %s:   width: %d, height: %d\n", time.Since(start), w, h)
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
