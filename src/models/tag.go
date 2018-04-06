package models

import (
	"encoding/json"
	"fmt"
)

type Tag struct {
	Name   string
	Status string
}

func NewTag(tagName string) Tag {
	return Tag{Name: tagName, Status: "created"}
}

func (t *Tag) AsPayLoadFormat() interface{} {

	tagJSON := []byte(fmt.Sprintf(`{
    "tagInst": {
      "attributes": {
        "name": " %s",
        "status": "created"
      },
      "children": []
    }
  }`, t.Name))

	var tagData interface{}

	json.Unmarshal(tagJSON, tagData)

	return tagData

}
