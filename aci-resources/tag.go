package main

type Tag struct {
  Name string
  Status string
}

func NewTag(tagName string) *Tag {
  return &Tag{Name: tagName, Status: "created", ObjectName: "tagInst"}
}
