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
	diary := flag.Bool("diary", false, "Use lexically ordered diary format, overrides 'short'")
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

		grYear := t.Year()
		grYearDay := t.YearDay()
		var out string
		if *diary {
			dayCode := tqtime.DayCode(tqtime.Day(grYear, grYearDay))
			monthLetter := tqtime.MonthLetter(tqtime.Month(grYear, grYearDay))
			year := tqtime.Year(grYear, grYearDay)
			out = fmt.Sprintf("%d-%s%02s", year, monthLetter, dayCode)
		} else if *short {
			out = tqtime.ShortDate(grYear, grYearDay)
		} else {
			out = tqtime.LongDate(grYear, grYearDay)
		}
		fmt.Println(out)
	}
}
