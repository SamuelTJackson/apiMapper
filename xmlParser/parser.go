package xmlParser

import (
	"encoding/xml"
	"fmt"
	"github.com/SamuelTJackson/apiMapper/staticErrors"
	"os"
	"reflect"
	"strings"
)

type Tag struct {
	XMLName    xml.Name
	Content    string
	XMLPath    string
	Attributes []xml.Attr
	Parent     *Tag
	Childes    []*Tag
}
type Attribute struct {
	Name  string
	Value string
}

func printXML(tag *Tag) {
	for _, child := range tag.Childes {
		if child.Content != "" {
			fmt.Println(child.XMLName, ":", child.Content)
			fmt.Println("xmlPath:", child.XMLPath)
		} else {
			fmt.Println(child.XMLName)
		}

		if len(child.Childes) > 0 {
			printXML(child)
		}
	}
}

func getImportantTags(tag *Tag, currentList []Tag) ([]Tag, error) {
	if tag == nil {
		return currentList, nil
	}
	for _, child := range tag.Childes {
		if child.Content != "" {
			currentList = append(currentList, *child)
		}
		currentList, _ = getImportantTags(child, currentList)
	}
	return currentList, nil
}
func GetImportantTags(tag *Tag) (list []Tag, err error) {
	return getImportantTags(tag, list)
}

func ParseURL(url string) ([]Tag, error) {
	if url == "" {
		return nil, staticErrors.EmptyRequest
	}
	// Open our xmlFile
	xmlFile, err := os.Open("test/test.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	var parents []*Tag
	var root Tag
	root.XMLPath = "/"
	parents = append(parents, &root)
	//var tags []Tag
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			fmt.Println("t is nil")
			break
		}
		fmt.Println(reflect.TypeOf(t).String())
		// Inspect the type of the token just read.
		switch se := t.(type) {
		case xml.StartElement:
			var tag Tag
			if len(parents) > 0 {
				parents[len(parents)-1].Childes = append(parents[len(parents)-1].Childes, &tag)
				tag.Parent = parents[len(parents)-1]
				tag.Attributes = se.Attr
				tag.XMLName = se.Name
				tag.XMLPath = parents[len(parents)-1].XMLPath + "/" + se.Name.Local
			} else {
				panic("parents can not be empty")
			}

			parents = append(parents, &tag)

		case xml.EndElement:
			parents = parents[:len(parents)-1]
		case xml.CharData:
			parents[len(parents)-1].Content = strings.TrimSpace(string(se))
		}

	}
	//printXML(&root)

	list, err := GetImportantTags(&root)
	for _, entry := range list {
		fmt.Println("Content:", entry.Content)
	}
	return list, err
}
