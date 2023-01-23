package ioutil

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
)

func ParseIPBlockList(r io.Reader) (map[string]string, error) {

	ipMap := make(map[string]string, 0)

	// read csv values using csv.Reader
	csvReader := csv.NewReader(r)
	csvReader.Comma = ';'

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		ipMap[rec[0]] = rec[2]
	}
	return ipMap, nil
}

func DownloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
