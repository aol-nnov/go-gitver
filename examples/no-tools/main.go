//go:generate go run -mod=vendor git.rootprojects.org/root/go-gitver --fail

package main

import (
	"flag"
	"fmt"
)

var (
	GitRev       = "0000000"
	GitVersion   = "v0.0.0-pre0+0000000"
	GitTimestamp = "0000-00-00T00:00:00+0000"
)

func main() {
	showVersion := flag.Bool("version", false, "Print version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Println(GitRev)
		fmt.Println(GitVersion)
		fmt.Println(GitTimestamp)
		return
	}

	fmt.Println("Hello, World!")
}
