package bunker

import (
	"log"
	"testing"
)

func TestBunkerHttp(t *testing.T) {
	req := New("http://api.ipify.org").Get().SetDebug(true).SetHeaders(map[string][]string{
		"x-app-origin": {"xman"},
	}).Do()

	if req.HaveError() {
		log.Println("error", req.Errors)
	}
}
