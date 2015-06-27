package crawler

import (
	"./../xml_parser"
	"fmt"
	"github.com/franela/goreq"
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
		Aplimit: 500,
		Format:  "xml",
	}

	res, errR := goreq.Request{
		Uri:         "http://zh.moegirl.org/api.php",
		QueryString: first,
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
			Aplimit:    500,
			Format:     "xml",
			Apcontinue: CurrentResult.Next,
		}

		fmt.Println("The next is " + next.Apcontinue)
		time.Sleep(500 * time.Millisecond)
		res, errR = goreq.Request{
			Uri:         "http://zh.moegirl.org/api.php",
			QueryString: next,
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
