package main

import (
	"cloud.google.com/go/translate"
	"context"
	"flag"
	"fmt"
	lang "golang.org/x/text/language"
	"io/ioutil"
	"log"
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
	fmt.Printf("Current language: %s\n", *language)

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

	fmt.Printf("From %s: %s\n", *location, *text)

	c := context.Background()
	client, err := translate.NewClient(c)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	var t []string
	t = append(t, *text)
	translations, err := client.Translate(c, t, tag, nil)
	if err != nil {
		log.Fatalf("Failed to Translate Text: %v", err)
	}
	fmt.Printf("Translated: %s\n", translations[0].Text)

	return
}
