package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/feeds"
	"log"
	"os"
	"time"
)

type jsfAttachment struct {
	URL      string `json:"url"`
	MIMEType string `json:"mime_type"`
	valid    bool
}

type jsfItem struct {
	ID            string          `json:"id"`
	URL           string          `json:"url"`
	Title         string          `json:"title"`
	ContentHTML   string          `json:"content_html"`
	DatePublished string          `json:"date_published"`
	DateModified  string          `json:"date_modified"`
	Tags          []string        `json:"tags"`
	Attachments   []jsfAttachment `json:"attachments"`
}

type jsfMain struct {
	Version     string    `json:"version"`
	Title       string    `json:"title"`
	HomePageURL string    `json:"home_page_url"`
	Items       []jsfItem `json:"items"`
}

func fromJsfItem(gi *feeds.Item, ji jsfItem) {
	gi.Title = ji.Title
	gi.Link = &feeds.Link{Href: ji.URL}
	gi.Created, _ = time.Parse(time.RFC3339, ji.DatePublished)
	gi.Updated, _ = time.Parse(time.RFC3339, ji.DateModified)
	gi.Id = ji.URL
	gi.Description = ji.ContentHTML
}

func main() {
	hostRawURL := flag.String("host", "http://ratan.blog", "Website URL")
	atomFlag := flag.Bool("atom", false, "Output Atom Feed")
	rssFlag := flag.Bool("rss", false, "Output RSS Feed")
	help := flag.Bool("help", false, "Print command-line options")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	var jsf jsfMain
	d := json.NewDecoder(os.Stdin)
	decodeErr := d.Decode(&jsf)
	if decodeErr != nil {
		log.Fatal(decodeErr)
	}

	var gf feeds.Feed
	gf.Title = jsf.Title
	gf.Link = &feeds.Link{Href: *hostRawURL}
	gf.Created = time.Now()
	gf.Items = make([]*feeds.Item, len(jsf.Items))
	for i, ji := range jsf.Items {
		gf.Items[i] = new(feeds.Item)
		fromJsfItem(gf.Items[i], ji)
	}

	if *rssFlag {
		rss, rssErr := gf.ToRss()
		if rssErr != nil {
			log.Fatal(rssErr)
		}
		fmt.Println(rss)
	}

	if *atomFlag {
		atom, atomErr := gf.ToAtom()
		if atomErr != nil {
			log.Fatal(atomErr)
		}
		fmt.Println(atom)
	}
}
