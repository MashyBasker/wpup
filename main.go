package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func main() {
	// read the directory path from command-line parameters
	if len(os.Args) < 2 {
		fmt.Println("Usage: wpup <directory path>")
		os.Exit(0)
	}
	wpDir := os.Args[1]

	// get the channel webhook URL from the environment variables
	val, present := os.LookupEnv("DISCORD_WEBHOOK")
	if !present {
		log.Fatal("[ERR] Webhook environment variable is not set")
	} 

	// get the list of all the files
	fileList := ListFiles(wpDir)
	
	// send the files to specified discord channel
	for _, e := range(fileList) {
		SendFile(e, val)
		pathSplit := strings.Split(e, "/")
		fmt.Printf("✅ %s has been sent\n", pathSplit[len(pathSplit)-1])
	}
	fmt.Println("Files have been sent successfully 🔥")
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
	pathSplit := strings.Split(filePath, "/")
	part, err := writer.CreateFormFile("file", pathSplit[len(pathSplit)-1])
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

func ListFiles(dirPath string) []string{
	// expland tilde path to absolute path
	user, err := user.Current()
	if err != nil {
		log.Fatal("[ERR] Could not get username", err)
	}
	homeDir := user.HomeDir
	if dirPath == "~" {
		dirPath = homeDir
	} else if strings.HasPrefix(dirPath, "~/") {
		dirPath = filepath.Join(homeDir, dirPath[2:])
	}

	// read contents of the dir into a list
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal("[ERR] Could not read directory")
	}

	var fileList []string

	// filter all the files into a slice
	for _, e := range entries {
		if !e.IsDir() {
			fileList = append(fileList, dirPath +"/"+ e.Name())
		}
	}
	return fileList
}