package main

import (
	"flag"
	"os"
	"io"
	"bufio"
	"github.com/juju/ratelimit"
)

func main() {
	bps := flag.Float64("kbps", 0, "kilo bytes per second")
	flag.Parse()
	
	reader := bufio.NewReader(os.Stdin)
	var writer io.Writer
	if (*bps != 0) {
		rl := ratelimit.NewBucketWithRate(*bps * 1024, 1024 * 1024)
		writer = ratelimit.Writer(bufio.NewWriter(os.Stdout), rl)
	} else {
		writer = bufio.NewWriter(os.Stdout)
	}

	_, err := io.Copy(writer, reader)
	if err != nil {
		panic(err)
	}
}
