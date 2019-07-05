package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/Junhyeok99/SimpleTranslate/server"
	"github.com/Junhyeok99/SimpleTranslate/trans"
	lang "golang.org/x/text/language"
	"gopkg.in/gookit/color.v1"
)

func main() {
	useServer := flag.Bool("s", false, "Use local server to Translate")
	isFile := flag.Bool("f", false, "Text Source from File")
	location := flag.String("l", "example.txt", "Location of file")
	text := flag.String("t", "", "Translating text")
	language := flag.String("lang", "ko", "Target language to Translate")
	
	flag.Parse()
	if *useServer {
		server.StartServer()
	}

	if !*isFile && *text == "" {
		flag.Usage()
		return
	}
	//Handling Flags

	tag, err := lang.Parse(*language)
	if err != nil {
		log.Fatalf("Invalid Language Code %s: %v", *language, err)
	}

	if *isFile {
		b, e := ioutil.ReadFile(*location)
		if e != nil {
			log.Fatalf("Cannot open File %s: %v\n", *location, e)
		}
		*text = string(b)
		*location = "File"
	} else {
		*location = "Command Line"
	}
	texts := strings.Split(*text, "\n")

	color.New(color.FgCyan, color.OpBold).Printf("Text from %s:\n", *location)
	for i, _ := range texts {
		fmt.Println(texts[i])
	}

	color.New(color.FgCyan, color.OpBold).Printf("\nTranslated with language [%s]:\n", *language)
	translations := trans.TextTranslation(texts, tag)
	for i, _ := range translations {
		fmt.Println(translations[i].Text)
	}

	return
}
