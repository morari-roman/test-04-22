package request

import (
	"fmt"
	"io"
	"net/http"

	"github.com/morari-roman/test-04-22/asana/internal/asana"
)

const (
	baseUrl = "https://app.asana.com/api/1.0"
	token   = "2/1210044610546546/1210044739009575:79a8dd9042370a158955c9c5da52b291"
)

type AsanaData interface {
	DecodeData()
}

type Request struct {
	baseUrl   string
	suffix    string
	asanaData *asana.AsanaData
}

func New(sf string) *Request {
	return &Request{
		suffix:    sf,
		baseUrl:   baseUrl,
		asanaData: asana.New(),
	}
}

func (r *Request) MakeRequest() (*asana.AsanaData, error) {
	url := fmt.Sprintf("%s/%s", r.baseUrl, r.suffix)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	err = r.asanaData.DecodeData(body)

	return r.asanaData, err
}
