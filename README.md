# gomod

Install with:

```go
go get github.com/aunsira/gomod
```

# Usage

Create a project and supply `ProjectKey` by `gomod.NewClient(ProjectKey)` and use actions object from the
`github.com/aunsira/gomod/actions` package to perform API operations.

```go
package main

import (
	"fmt"
	"log"

	gomod "github.com/aunsira/gomod"
	"github.com/aunsira/gomod/actions"
)

const (
	ProjectKey = <Project-Key>
)

func main() {
  createMod()
	getMod()
}

func getMod() {
	c := setProjectKey()

	var resp map[string]interface{}

	getModerations := &actions.GetModerations{
		ID: ModerationID,
	}

	if err := c.Call(&resp, getModerations); err != nil {
		log.Fatal(err)
	}

	data := resp["data"].(map[string]interface{})
	attrs := data["attributes"].(map[string]interface{})
	log.Println(attrs["answer"])
}

func createMod() {
	c := setProjectKey()
	moderationData, postModeration := &gomod.PostModeration{}, &actions.PostModeration{
		Data:           Data,
		PostbackURL:    PostbackURL,
		PostbackMethod: PostbackMethod,
		CustomID:       CustomID,
	}

	if err := c.Call(moderationData, postModeration); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Moderation: %#v\n", moderationData)
	log.Println("Resp data:  " + moderationData.Data.ID)
}

func setProjectKey() *gomod.Client {
	c, err := gomod.NewClient(ProjectKey)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
```
