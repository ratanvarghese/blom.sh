package main

import (
	"encoding/json"
	"fmt"
	"github.com/ratanvarghese/tqtime"
	"log"
	"os"
	"time"
)

type jsfItem struct {
	URL           string `json:"url"`
	Title         string `json:"title"`
	DatePublished string `json:"date_published"`
}

func archiveSeperator(gt1 time.Time, gt2 time.Time) (bool, string) {
	g1Year := gt1.Year()
	g1YearDay := gt1.YearDay()
	tq1Year := tqtime.Year(g1Year, g1YearDay)
	tq1Mon := tqtime.Month(g1Year, g1YearDay)
	tq1Day := tqtime.Day(g1Year, g1YearDay)

	g2Year := gt2.Year()
	g2YearDay := gt2.YearDay()
	tq2Year := tqtime.Year(g2Year, g2YearDay)
	tq2Mon := tqtime.Month(g2Year, g2YearDay)
	tq2Day := tqtime.Day(g2Year, g2YearDay)

	isSpecialDay := (tq2Mon == tqtime.SpecialDay)

	var seperatorText string
	if tq2Day == tqtime.AldrinDay || tq2Mon == tqtime.Hippocrates { //A feature of this calendar which is annoying for archives
		seperatorText = fmt.Sprintf("### Hippocrates & Aldrin Day, %d AT", tq2Year)
		if (tq1Mon == tqtime.Hippocrates || tq1Day == tqtime.AldrinDay) && tq1Year == tq2Year {
			return false, seperatorText
		}
	} else if isSpecialDay {
		seperatorText = fmt.Sprintf("### %s, %d AT", tqtime.DayName(tq2Day), tq2Year)
	} else {
		seperatorText = fmt.Sprintf("### %s, %d AT", tq2Mon.String(), tq2Year)
	}
	needSeperation := (tq1Year != tq2Year) || (tq1Mon != tq2Mon) || (isSpecialDay && (tq1Day != tq2Day))
	return needSeperation, seperatorText
}

func archiveLines(itemList []jsfItem) []string {
	if len(itemList) < 1 {
		return nil
	}
	var t1 time.Time //intentionally starting at zero value, always a different year than first article.
	outputLines := make([]string, 0)
	itemCount := len(itemList)
	for i, ji := range itemList {
		t2, _ := time.Parse(time.RFC3339, ji.DatePublished)
		if sep, sepText := archiveSeperator(t1, t2); sep {
			if i > 0 { //The start of a section is the end of the previous section, unless *no* previous section.
				outputLines = append(outputLines, "")
			}
			outputLines = append(outputLines, sepText)
		}
		shortPath := itemCount - i
		shortURL := "r3n.me"
		line := fmt.Sprintf("+ [%v](%v), short URL: [%v/%v](/%v)", ji.Title, ji.URL, shortURL, shortPath, shortPath)
		outputLines = append(outputLines, line)
		t1 = t2
	}
	return outputLines
}

func main() {
	var items []jsfItem
	d := json.NewDecoder(os.Stdin)
	err := d.Decode(&items)
	if err != nil {
		log.Fatal(err)
	}
	lines := archiveLines(items)
	for _, line := range lines {
		fmt.Println(line)
	}
}
