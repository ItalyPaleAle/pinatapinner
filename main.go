package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const endpoint = "https://api.pinata.cloud/pinning/pinFileToIPFS"

var (
	pinataAPIKey    string
	pinataSecretKey string
)

// Just like main, but supports "return <status_code>" to make things easier
func mainWithExitCode() int {
	// Check if the env vars are set
	pinataAPIKey = os.Getenv("PINATA_API_KEY")
	pinataSecretKey = os.Getenv("PINATA_SECRET_KEY")
	if pinataAPIKey == "" || pinataSecretKey == "" {
		fmt.Println("Environmental variables `PINATA_API_KEY` and `PINATA_SECRET_KEY` are required")
		return 1
	}

	// Ensure that we have a folder to upload
	if len(os.Args) < 2 {
		fmt.Println("Need to specify the path to upload")
		return 1
	}
	folder := os.Args[1]

	// TODO: CHECK THAT FOLDER EXISTS

	// Check if we have a name for the bundle
	name := "Uploaded via pinata-uploader"
	if len(os.Args) > 2 {
		name = os.Args[2]
	}

	// Upload the folder
	err := uploadFolder(folder, name)
	if err != nil {
		fmt.Println("Error while uploading folder:", err)
		return 2
	}

	return 0
}

// Uploads a folder
func uploadFolder(folder string, name string) error {
	// Build the request
	fu := NewFormUploader()

	files := []string{
		"index.html",
		"about.html",
		"feed.xml",
		"page/1.html",
		"page/2.html",
	}
	fu.AddFiles("file", "../drop", files...)

	// Send the request
	client := &http.Client{
		// Do not set a timeout, as the files might be large
		Timeout: 0,
	}
	headers := make(map[string]string, 2)
	headers["pinata_api_key"] = pinataAPIKey
	headers["pinata_secret_api_key"] = pinataSecretKey
	resp, err := fu.Post(client, endpoint, headers)
	if err != nil {
		return err
	}

	// Get the response
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Output the response (should be a JSON)
	fmt.Println(string(res))

	return nil
}

// Entry point
func main() {
	os.Exit(mainWithExitCode())
}
