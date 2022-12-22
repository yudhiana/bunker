package bunker

import (
	"log"
	"testing"
)

func TestBunkerHttp(t *testing.T) {
	req := New("http://api.ipify.org").Get().SetDebug(true).Query(map[string]string{
		"data": `{"id":1}`,
	}).SetHeader("x-app-origin", "xman", "wolverine").Do()

	if req.HaveError() {
		log.Println("error", req.Errors)
	}
}
