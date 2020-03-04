package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"sainsburys-stock/domain"
	"sainsburys-stock/images"
	"sainsburys-stock/output"
	"sainsburys-stock/upload"
)

func main() {

	file, err := images.Download()
	if err != nil {
		log.Print(err)
	}

	err = upload.Uploader("123456.jpg", file)
	if err != nil {
		log.Print(err)
	}

	results, err := ReadData("productNames.csv")
	if err != nil {
		log.Fatal(err)
	}

	writer, err := output.NewWriterCSV("out.csv", results)

	if err != nil {
		log.Fatal(err)
	}

	writer.Output()
}

func ReadData(filename string) (chan domain.Result, error) {
	results := make(chan domain.Result)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)

	go func() {
		defer close(results)
		for fileScanner.Scan() {
			time.Sleep(time.Millisecond * 2000)
			product := strings.Split(string(fileScanner.Bytes()), ",")
			res, err := makeRequest(product[0])
			if err != nil {
				log.Println(err)
				continue
			}
			results <- *res
		}
	}()

	return results, nil
}

func makeRequest(product string) (*domain.Result, error) {

	url := "https://help.sainsburys.co.uk/help/_stocklookup"

	payload := strings.NewReader(fmt.Sprintf("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"query\"\r\n\r\n %s \r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"store\"\r\n\r\n0500\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--", product))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
	req.Header.Add("User-Agent", "PostmanRuntime/7.20.1")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "help.sainsburys.co.uk")
	req.Header.Add("Content-Type", "multipart/form-data; boundary=--------------------------363133940998647992413471")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Cookie", "vsid=eyJpdiI6InIwdFVxcEExMW1zYU1rQ1BVRkluUFE9PSIsInZhbHVlIjoiQnNwalhwREVqMHBqNzFwR2htRklcL3RKSGR5RzdrZUVnUEV6cmRaeE1HemRERXNGMXB2cHpBMHdkS3U5eGtwTVlYVlhaelFJRFlrWUVqUHlKMGFnQkFnPT0iLCJtYWMiOiI5MGM5MmViOTM0ZTM0MDRhMjQxYWVkNzA3NzljYTdkN2NkZTgyNDIwNzc5ZGM0MzliZDRjZDNhZjM1ZWJjNjQxIn0%3D")
	req.Header.Add("Content-Length", "295")
	req.Header.Add("Connection", "keep-alive")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	sa := domain.StoreAvailability{}

	err := json.Unmarshal(body, &sa)
	if err != nil {
		return nil, err
	}

	for _, item := range sa.Items {

		if item.Description == product {
			log.Println(product)
			return &domain.Result{
				Product:   product,
				Available: item.Store.Stock.OnHand,
				NotFound:  false,
			}, nil
		}
	}

	return &domain.Result{
		Product:  product,
		NotFound: true,
	}, nil
}
