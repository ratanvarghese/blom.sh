package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type jsfItem struct {
	URL   string   `json:"url"`
	Title string   `json:"title"`
	Tags  []string `json:"tags"`
}

func tagSort(itemList []jsfItem) (map[string][]jsfItem, []string) {
	res := make(map[string][]jsfItem)
	tagList := make([]string, 0)
	for _, ji := range itemList {
		for _, tag := range ji.Tags {
			if len(res[tag]) == 0 && len(tag) > 0 {
				tagList = append(tagList, tag)
			}
			res[tag] = append(res[tag], ji)
		}
	}
	sort.Strings(tagList)
	return res, tagList
}

func tagsPageLines(itemList []jsfItem) []string {
	outputLines := make([]string, 0)
	tagMap, tagList := tagSort(itemList)
	for _, tag := range tagList {
		outputLines = append(outputLines, fmt.Sprintf("### %v", strings.Title(tag)))
		for _, ji := range tagMap[tag] {
			outputLines = append(outputLines, fmt.Sprintf("+ [%v](%v)", ji.Title, ji.URL))
		}
		outputLines = append(outputLines, "")
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
	lines := tagsPageLines(items)
	for _, line := range lines {
		fmt.Println(line)
	}
}
