package gomod

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/aunsira/gomod/actions"
	"github.com/aunsira/gomod/config"
	a "github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

const (
	TestProjectKey = "A7sRxaQKxo2hRQzNwkk5Qqx4"
	TestImageID    = "5a44671ab3957c2ab5c33326"
	TestImageData  = "https://assets-cdn.github.com/images/modules/open_graph/github-mark.png"
)

func readFile(path string) []byte {
	result, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
	}
	return result
}

type mockPayload struct {
	ID       string
	CustomID string
}

func (*mockPayload) Endpoint() (string, string, string) {
	return config.PosmoniAPIURL, "method", "path"
}

func (m *mockPayload) Payload(endpoint, method, path string) (*http.Request, error) {
	return nil, errors.New("Mock error for payload testing")
}

func TestNewClient(t *testing.T) {
	c, err := NewClient(TestProjectKey)
	a.Nil(t, err)
	a.NotNil(t, c)
}

func TestNewClient_ErrorInvalidKey(t *testing.T) {
	c, err := NewClient("")
	a.NotNil(t, err)
	a.Nil(t, c)
}

func TestClient_CallGetClosedQuestion(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	cq, getCQ := &GetModeration{}, &actions.GetModeration{
		ID: TestImageID,
	}

	endpoint, _, path := getCQ.Endpoint()
	mockResp := readFile("./testdata/closed_question.json")
	gock.New(endpoint).
		Get(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(cq, getCQ)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, cq)
	a.Equal(t, getCQ.ID, cq.Data.ID)
}

func TestClient_CallGetListClosedQuestion(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	cq, getCQ := &GetModerations{}, &actions.GetModerations{
		ID: TestImageID,
	}

	endpoint, _, path := getCQ.Endpoint()
	mockResp := readFile("./testdata/closed_questions.json")
	gock.New(endpoint).
		Get(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(cq, getCQ)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, cq)
	a.Equal(t, getCQ.ID, cq.Data[0].ID)
}

func TestClient_CallPostClosedQuestion(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	pq, postCQ := &PostModeration{}, &actions.PostModeration{
		Data: TestImageData,
	}

	endpoint, _, path := postCQ.Endpoint()
	mockResp := readFile("./testdata/post_closed_question.json")
	gock.New(endpoint).
		Post(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(pq, postCQ)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, pq)
	a.Equal(t, postCQ.Data, pq.Data.Source)
}

func TestClient_InvalidCall(t *testing.T) {
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	closedQuestion, getImage := &GetModeration{}, &actions.GetModeration{}

	endpoint, _, path := getImage.Endpoint()
	gock.New(endpoint).
		Get(path).
		Reply(401)

	e := c.Call(closedQuestion, getImage)
	a.EqualError(t, e, e.Error())
}

func TestClient_InvalidPayload(t *testing.T) {
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	m := &mockPayload{}
	e := c.Call(nil, m)
	a.NotNil(t, e)
	a.EqualError(t, e, "Mock error for payload testing")
}
