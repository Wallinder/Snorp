package state

import (
	"fmt"
	"io"
	"net/http"
)

type HttpRequest struct {
	Method string
	Uri    string
	Body   io.Reader
}

func (session *SessionState) SendRequest(r HttpRequest) (*http.Response, error) {
	url := session.Config.Bot.Api + r.Uri

	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		return nil, err
	}

	req.Header = session.GlobalHeaders

	res, err := session.Client.Do(req)
	if err != nil {
		return nil, err
	}

	statuscode := res.StatusCode

	if statuscode != 200 && statuscode != 201 {
		return res, fmt.Errorf("%d", statuscode)
	}
	return res, nil
}
