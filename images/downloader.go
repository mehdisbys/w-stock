package images

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Download() ([]byte, error) {
	url := "http://i.imgur.com/m1UIjW1.jpg"

	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
