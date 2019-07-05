package server

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Junhyeok99/SimpleTranslate/trans"
	lang "golang.org/x/text/language"
)

func printBoth(w http.ResponseWriter, s string) {
	fmt.Fprintf(w, s)
	fmt.Println(s)
}

func transHandler(w http.ResponseWriter, r *http.Request) {
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		printBoth(w, fmt.Sprintf("could not parse query: %v\n", err))
	} else if len(params) == 0 {
		printBoth(w, fmt.Sprintf("Please Enter query to parse\n"))
	} else {
		for key, _ := range params {
			if len(params[key]) >= 2 {
				printBoth(w, fmt.Sprintf("flag(%s) must have only parameter\n", key))
			} else {
				text := params["-t"][0]
				texts := strings.Split(text, "\n")
				printBoth(w, "Original Text: ")
				for i, _ := range texts {
					printBoth(w, texts[i])
				}

				printBoth(w, "\nTranslated Text: ")
				translations := trans.TextTranslation(texts, lang.Korean)
				for i, _ := range translations {
					printBoth(w, translations[i].Text)
				}
			}
		}
	}
}

func StartServer() {
	fmt.Println("Serving on http://localhost:1357/")
	http.Handle("/", http.FileServer(http.Dir("server/resource")))
	http.HandleFunc("/translate", transHandler)
	http.ListenAndServe(":1357", nil)
}