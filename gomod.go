/*
Package gomod provides GO binding for Posmoni REST API.
Full REST API documentation is available at https://datawow.readme.io/v1.0/reference.

Usage

Create a client with gomod.NewClient, along with supply your project key. After that, use
client.Call with actions object from the https://godoc.org/github.com/aunsira/gomod/actions
package to perform API calls. The first parameter to client.Call lets you supply a struct
object from index that listed below to unmarshal the result.

Example

	c, err := gomod.NewClient(ProjectKey)
	if err != nil {
		log.Fatal(err)
	}

	Moderation, getModeration := &gomod.GetModeration{}, &actions.GetModeration{
		ID: "5a52fb556e11571f570c1530",
	}

	if err := c.Call(Moderation, getModeration); err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", Moderation)

We also provide Get "any type" Image endpoint API. You only supply project key and Image
ID (or Customer ID) for reference.

Example

	c, err := gomod.NewClient(ProjectKey)
	if err != nil {
		log.Fatal(err)
	}
	resp := make(map[string]interface{})

	getImage := &actions.GetImage{
		ID: "5a52fb556e11571f570c1530",
	}

	if err := c.Call(&resp, getImage); err != nil {
		log.Fatal(err)
	}

	data := resp["data"].(map[string]interface{})
	meta := resp["meta"].(map[string]interface{})
	image := data["image"].(map[string]interface{})
	log.Println("Image ID: " + image["id"])
	log.Println("Image Status: " + image["status"])
	log.Println("Response code: " + meta["code"])

*/
package gomod
