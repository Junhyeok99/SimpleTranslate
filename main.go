package main

import (
	"cloud.google.com/go/translate"
	"context"
	"flag"
	"fmt"
	lang "golang.org/x/text/language"
	"gopkg.in/gookit/color.v1"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	isFile := flag.Bool("f", false, "Text Source from File")
	location := flag.String("l", "example.txt", "Location of file")
	text := flag.String("t", "", "Translating text")
	language := flag.String("lang", "ko", "Target language to Translate")

	flag.Parse()
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
	var t []string
	for i, _ := range texts {
		fmt.Println(texts[i])
		t = append(t, texts[i])
	}

	c := context.Background()
	client, err := translate.NewClient(c)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	translations, err := client.Translate(c, t, tag, nil)
	if err != nil {
		log.Fatalf("Failed to Translate Text: %v", err)
	}
	color.New(color.FgCyan, color.OpBold).Printf("\nTranslated with language [%s]:\n", *language)
	for i, _ := range translations {
		fmt.Println(translations[i].Text)
	}

	return
}
