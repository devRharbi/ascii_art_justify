package ascii

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	size  int
	align string
)

func CheckArgsAndFlag(arg []string) []string {
	flag.CommandLine.SetOutput(io.Discard)
	flag.Usage = func() {
		MessageErrors()
	}
	flag.StringVar(&align, "align", "", "Specify the text alignment (left, right, center)")
	flag.Parse()
	args := flag.Args()

	if align != "" {
		size = len(align)
		if size > 0 && (!strings.HasPrefix(arg[1], "--align=")) {
			MessageErrors()
			os.Exit(0)
		}
	}
	return args
}

func PrintResult(sl [][]string, s string) {
	result := PrintAsciArt(sl, s, align)

	if align == "center" || align == "left" || align == "right" {
		AlignText(result, align)
	} else {
		fmt.Print(result)
	}
	// if size == 0 {
	// 	fmt.Print(result)
	// }
}
