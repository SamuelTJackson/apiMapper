package db

import (
	"github.com/SamuelTJackson/apiMapper/staticErrors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var urlMapper map[string]string
var requestFiles map[string]string
var paths map[string][]string

func init() {
	urlMapper = make(map[string]string)
	requestFiles = make(map[string]string)
	paths = make(map[string][]string)

}

func GetURLForID(id string) (string, error) {
	if url, ok := urlMapper[id]; ok {
		return url, nil
	}
	return "", staticErrors.IDDoesNotExists
}
func ReadDataIntoMemory() {
	log.Printf("read database into memory")
	files, err := ioutil.ReadDir("./mappings/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		completeFileName := f.Name()
		log.Println("Reading file:", completeFileName)
		var fileExtension = filepath.Ext(completeFileName)
		log.Println("file extension:", fileExtension)
		var fileName = f.Name()[0 : len(completeFileName)-len(fileExtension)]
		log.Println("file name:", fileName)
		if fileExtension == ".jsonM" {
			file, err := os.Open("./mappings/" + completeFileName)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			b, err := ioutil.ReadAll(file)
			requestFiles[fileName] = string(b)

		} else {
			if fileExtension == ".mapping" {
				file, err := os.Open("./mappings/" + f.Name())
				if err != nil {
					log.Fatal(err)
				}
				defer file.Close()
				b, err := ioutil.ReadAll(file)
				urlMapper[fileName] = string(b)
			}
		}

	}
}
