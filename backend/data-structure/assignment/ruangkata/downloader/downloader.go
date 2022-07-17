package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Download(url, outputFile string) error {
	fmt.Println("Downloading english dictionary...")

	data, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer data.Body.Close()

	_, err = os.Stat(outputFile)
	var fOutput *os.File
	if err == nil {
		fmt.Println("File already exists, skipping download:", outputFile)
		return nil
	}
	fOutput, err = os.Create(outputFile)
	if err != nil {
		return err
	}
	fmt.Println("Downloaded to:", outputFile)
	io.Copy(fOutput, data.Body)
	return nil
}
