package translate

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type client struct {
	location        string
	endpoint        string
	subscriptionKey string
}

func NewClient(key, location string, endpoints ...string) *client {
	endpoint := "https://api.cognitive.microsofttranslator.com/"
	for _, v := range endpoints {
		endpoint = v
		break
	}
	return &client{
		location:        location,
		endpoint:        endpoint,
		subscriptionKey: key,
	}
}

func (c *client) To(req *Request) ([]Response, error) {
	uri := c.endpoint + "/translate?api-version=3.0"

	u, _ := url.Parse(uri)
	q := u.Query()
	q.Add("from", req.Language)
	for _, v := range req.ToLanguage {
		q.Add("to", v)
	}
	u.RawQuery = q.Encode()
	body := req.Data
	b, _ := json.Marshal(body)
	tReq, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	tReq.Header.Add("Ocp-Apim-Subscription-Key", c.subscriptionKey)
	tReq.Header.Add("Ocp-Apim-Subscription-Region", c.location)
	tReq.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(tReq)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Decode the JSON response
	var result []Response
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}
