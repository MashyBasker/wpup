package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"mime/multipart"
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func SendFile(filePath string, WebhookUrl string) error {
	// open the file you want to upload
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("[ERR] Error opening file")
		return err
	}
	defer file.Close()

	// a buffer for storing the request body
	body := new(bytes.Buffer)

	// creating a multipart writer
	writer := multipart.NewWriter(body)

	// create new form file field
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		log.Fatal("[ERR] Could not create the form file")
		return err
	}

	// copy the file content to the form file field
	_, err = io.Copy(part, file)
	if err != nil {
		log.Fatal("[ERR] Could not copy file content")
		return err
	}

	// close multipart writer
	err = writer.Close()
	if err != nil {
		log.Fatal("[ERR] Could not close multipart")
		return err
	}

	// send the post request over http
	resp, err := http.Post(WebhookUrl, writer.FormDataContentType(), body)
	if err != nil {
		log.Fatal("[ERR Could not request")
		return err
	}
	defer resp.Body.Close()
	return nil
}


