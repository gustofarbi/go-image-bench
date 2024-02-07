package main

import (
	"github.com/davidbyttow/govips/v2/vips"
	"os"
	"os/exec"
	"regexp"
	"testing"
)

func BenchmarkVipsHeader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vipsHeaderTest()
	}
}

func BenchmarkGovips(b *testing.B) {
	for i := 0; i < b.N; i++ {
		govipsSizeTest()
	}
}

func govipsSizeTest() {
	vips.LoggingSettings(func(string, vips.LogLevel, string) {}, vips.LogLevelDebug)
	f, err := os.Open("testdata/1820MB-image.tif")
	panicOnError(err)
	im, err := vips.NewImageFromReader(f)
	panicOnError(err)
	im.Close()
}

func vipsHeaderTest() {
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
}
