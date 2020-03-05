package images

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Download(url string) ([]byte, error) {

	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
