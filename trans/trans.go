package trans

import (
	"context"

	"log"

	"cloud.google.com/go/translate"
	lang "golang.org/x/text/language"
)

func TextTranslation(s []string, t lang.Tag) []translate.Translation {
	c := context.Background()
	client, err := translate.NewClient(c)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	translations, err := client.Translate(c, s, t, nil)
	if err != nil {
		log.Fatalf("Failed to Translate Text: %v", err)
	}

	return translations
}