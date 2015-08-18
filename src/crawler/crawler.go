package crawler

import (
	"fmt"
	"github.com/franela/goreq"
	"github.com/wicast/moe-words-library/src/xml_parser"
	"time"
)

type Item struct {
	Action     string
	List       string
	Aplimit    int
	Format     string
	Apcontinue string
}

type ResultSet struct {
	Next  string
	Pages []xml_parser.Page
}

func QueryAll() []ResultSet {
	first := Item{
		Action:  "query",
		List:    "allpages",
		Aplimit: 100,
		Format:  "xml",
	}

	res, errR := goreq.Request{
		Uri:         "http://zh.moegirl.org/api.php",
		QueryString: first,
		UserAgent:   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.155 Safari/537.36",
	}.Do()

	if errR != nil {
		panic(errR)
	}

	xml_data, err := res.Body.ToString()
	if err != nil {
		panic(err)
	}

	raw_data := xml_parser.Parse(xml_data)

	CurrentResult := ResultSet{
		raw_data.Continue_query.Apcontinue,
		raw_data.Query.Pages,
	}

	Result := make([]ResultSet, 0)
	Result = append(Result, CurrentResult)

	for CurrentResult.Next != "" {
		next := Item{
			Action:     "query",
			List:       "allpages",
			Aplimit:    100,
			Format:     "xml",
			Apcontinue: CurrentResult.Next,
		}

		fmt.Println("The next is " + next.Apcontinue)
		time.Sleep(500 * time.Millisecond)
		res, errR = goreq.Request{
			Uri:         "http://zh.moegirl.org/api.php",
			QueryString: next,
			UserAgent:   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.155 Safari/537.36",
		}.Do()
		if errR != nil {
			panic(errR)
		}

		xml_data, err = res.Body.ToString()
		if err != nil {
			panic(err)
		}

		raw_data = xml_parser.Parse(xml_data)

		CurrentResult = ResultSet{
			raw_data.Continue_query.Apcontinue,
			raw_data.Query.Pages,
		}

		Result = append(Result, CurrentResult)
	}

	return Result
}
