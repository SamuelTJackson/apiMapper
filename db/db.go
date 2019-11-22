package db

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var urlMapper map[string]string
var requestFiles map[string]string

func init() {
	urlMapper = make(map[string]string)
	requestFiles = make(map[string]string)

}
func ReadDataIntoMemory() {
	log.Printf("read database into memory")
	files, err := ioutil.ReadDir("./mappings/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		log.Println(f.Name())
		var fileExtension = filepath.Ext(f.Name())
		var fileName = f.Name()[0 : len(f.Name())-len(fileExtension)]
		if fileExtension == "json" {
			file, err := os.Open("./mappings/" + f.Name())
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			b, err := ioutil.ReadAll(file)
			requestFiles[fileName] = string(b)

		} else {
			if fileExtension == "mapping" {
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
