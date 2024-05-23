package ascii

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func ContSpace(slice [][]string) int {
	m := 0
	for _, v := range slice {
		for i, n := range v {
			if n == "      " {
				if i == 7 {
					m++
				}
			} else {
				break
			}
		}
	}
	return m
}

func AddSpaces(slice [][]string) string {
	lenght := 0
	space := 0
	text := " "
	for i := 0; i < len(slice); i++ {
		lenght += len(slice[i][1])
	}
	width := GetWidthConsol()
	if ContSpace(slice) != 0 {
		space = (width - lenght) / ContSpace(slice)
	}
	return strings.Repeat(text, space)
}

func AlignText(text string, align string) {
	asciiArt := strings.Split(text, "\n")
	width := GetWidthConsol()

	for _, line := range asciiArt {

		lineLength := len(line)
		padding := (width - lineLength) / 2
		paddingrghit := (width - lineLength)

		leftPaddingStr := strings.Repeat(" ", padding)
		RghitPaddingStr := strings.Repeat(" ", paddingrghit)
		if align == "center" {
			fmt.Printf("%s%s\n", leftPaddingStr, line)
		}
		if align == "left" {
			fmt.Printf("%s", line)
		}
		if align == "right" {
			fmt.Printf("%s%s", RghitPaddingStr, line)
		}

	}
}

func GetWidthConsol() int {
	fd := int(os.Stdout.Fd())
	width, _, err := term.GetSize(fd)
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		os.Exit(0)
	}
	return width
}
