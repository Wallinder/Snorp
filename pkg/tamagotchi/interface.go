package tamagotchi

import (
	"io"
	"net/http"
)

type Reciever interface {
	NewRequest(method string, uri string, body io.Reader) (*http.Response, error)
}

func Notify(reciever Reciever, method string, uri string, body io.Reader) (*http.Response, error) {
	reponse, err := reciever.NewRequest(method, uri, body)
	return reponse, err
}
