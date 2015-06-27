package xml_parser

import (
	"encoding/xml"
)

type Entry struct {
	Continue_query allpages `xml:"query-continue>allpages"`
	Query          allpages          `xml:"query>allpages"`
}

type continue_query struct {
	NextPage allpages `xml:"allpages"`
}

type query struct {
	Allpages allpages `xml:"allpages"`
}

type allpages struct {
	Apcontinue string `xml:"apcontinue,attr"`
	Pages      []Page `xml:"p"`
}

type Page struct {
	Pageid int    `xml:"pageid,attr"`
	Title  string `xml:"title,attr"`
}

func Parse(s string) Entry {
	r := Entry{}
	err := xml.Unmarshal([]byte(s), &r)
	if err != nil {
		panic(err)
	}
	return r
}
