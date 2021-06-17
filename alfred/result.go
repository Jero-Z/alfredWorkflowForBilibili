package alfred

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Icon struct {
	Type string `json:"type,omitempty"`
	Path string `json:"path,‘omitempty"`
}

type Mod struct {
	Variables map[string]string `json:"variables,omitempty"`
	Valid     bool              `json:"valid,omitempty"`
	Arg       string            `json:"arg,omitempty"`
	Subtitle  string            `json:"subtitle,omitempty"`
	Icon      *Icon             `json:"icon,omitempty"`
}

type Text struct {
	Copy      string `json:"copy,omitempty"`
	LargeType string `json:"largetype,omitempty"`
}

type Item struct {
	Variables    map[string]string `json:"variables,omitempty"`
	Valid        bool              `json:"valid,omitempty"`
	UID          string            `json:"uid,omitempty"`
	Title        string            `json:"title"`
	Subtitle     string            `json:"subtitle,omitempty"`
	Arg          string            `json:"arg,omitempty"`
	Autocomplete string            `json:"autocomplete,omitempty"`
	Type         string            `json:"type,omitempty"`
	Match        string            `json:"match,omitempty"`
	QuickLookUrl string            `json:"quicklookurl,omitempty"`
	Mods         map[string]Mod    `json:"mods,omitempty"`
	Text         *Text             `json:"text,omitempty"`
	Icon         *Icon             `json:"icon,omitempty"`
}
type Result struct {
	Items []Item `json:"items"'`
}

func NewResult() *Result {
	i := make([]Item, 0)
	return &Result{
		Items: i,
	}
}

func (r *Result) Append(item Item) {
	r.Items = append(r.Items, item)
}
func (r *Result) Run() {
	b := new(bytes.Buffer)

	if err := json.NewEncoder(b).Encode(r); err != nil {
		log.Println(err)
		panic(err)
	}
	fmt.Print(b.String())
	os.Exit(0)
}
