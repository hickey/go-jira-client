package gojira

import (
	"net/url"
    "encoding/xml"
	"fmt"
	"time"
)



type ActivityItem struct {
	Title    string    `xml:"title"json:"title"`
	Id       string    `xml:"id"json:"id"`
	Link     []Link    `xml:"link"json:"link"`
	Updated  time.Time `xml:"updated"json:"updated"`
	Author   Person    `xml:"author"json:"author"`
	Summary  Text      `xml:"summary"json:"summary"`
	Category Category  `xml:"category"json:"category"`
}

type ActivityFeed struct {
	XMLName  xml.Name        `xml:"http://www.w3.org/2005/Atom feed"json:"xml_name"`
	Title    string          `xml:"title"json:"title"`
	Id       string          `xml:"id"json:"id"`
	Link     []Link          `xml:"link"json:"link"`
	Updated  time.Time       `xml:"updated,attr"json:"updated"`
	Author   Person          `xml:"author"json:"author"`
	Entries  []*ActivityItem `xml:"entry"json:"entries"`
}
	
type Category struct {
	Term string `xml:"term,attr"json:"term"`
}

type Link struct {
	Rel  string `xml:"rel,attr,omitempty"json:"rel"`
	Href string `xml:"href,attr"json:"href"`
}

type Person struct {
	Name     string `xml:"name"json:"name"`
	URI      string `xml:"uri"json:"uri"`
	Email    string `xml:"email"json:"email"`
	InnerXML string `xml:",innerxml"json:"inner_xml"`
}

type Text struct {
	Type string `xml:"type,attr,omitempty"json:"type"`
	Body string `xml:",chardata"json:"body"`
}



func (j *Jira) UserActivity(user string) (ActivityFeed, error) {
	url := j.BaseUrl + j.ActivityPath + "?streams=" + url.QueryEscape("user IS " + user)

	return j.Activity(url)
}

func (j *Jira) Activity(url string) (ActivityFeed, error) {

	contents := j.buildAndExecRequest("GET", url)

	var activity ActivityFeed
	err := xml.Unmarshal(contents, &activity)
	if err != nil {
		fmt.Println("%s", err)
	}

	return activity, err
}

