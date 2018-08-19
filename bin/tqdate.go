package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ratanvarghese/tqtime"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	inputFormat := flag.String("inputformat", time.UnixDate, "Reference date in Gregorian input format")
	input := flag.String("input", "", "Gregorian input date, use stdin if omitted")
	help := flag.Bool("help", false, "Print command-line options")
	short := flag.Bool("short", false, "Use short output format")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	var inputReader *bufio.Reader
	if *input == "" {
		inputReader = bufio.NewReader(os.Stdin)
	} else {
		inputReader = bufio.NewReader(strings.NewReader((*input) + "\n"))
	}

	for {
		line, lineErr := inputReader.ReadString('\n')
		if lineErr != nil {
			break
		}
		t, parseErr := time.Parse(*inputFormat, strings.TrimSpace(line))
		if parseErr != nil {
			log.Fatal(parseErr.Error())
			break
		}

		var out string
		if *short {
			out = tqtime.ShortDate(t.Year(), t.YearDay())
		} else {
			out = tqtime.LongDate(t.Year(), t.YearDay())
		}
		fmt.Println(out)
	}
}
