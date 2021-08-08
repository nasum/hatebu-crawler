package lib

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hokaccha/go-prettyjson"
)

type BookMark struct {
	Title     string   `json:"title"`
	URL       string   `json:"url"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
}

type BookMarkList []BookMark

func (bml *BookMarkList) Json() {
	var buf bytes.Buffer
	data, _ := json.Marshal(bml)
	err := json.Indent(&buf, []byte(data), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}

func (bml *BookMarkList) ShowJson() error {
	j, err := prettyjson.Marshal(bml)

	if err != nil {
		return err
	}
	fmt.Println(string(j))
	return nil
}
