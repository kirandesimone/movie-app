package helpers

import (
	"io"
	"log"
	"net/http"
)

func ValidateImg(uri string) string {
	if uri == "" {
		return "notFound"
	} else {
		res, err := http.Get(uri)
		if err != nil {
			log.Print("could not resolve")
		}

		content, e := io.ReadAll(res.Body)
		res.Body.Close()
		if e != nil {
			log.Fatal(e)
		}
		c := string(content)
		if c == "Not Found" {
			return "notFound"
		}
	}
	return uri
}
