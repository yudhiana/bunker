package bunker

import (
	"fmt"
	"net/http"
	"testing"
)

func TestBunkerError(t *testing.T) {

	errBadRequest := New(StatusBadRequest)

	if errBadRequest.Code != StatusBadRequest {
		t.Errorf("Invalid status code [%v] not equal bad requests", errBadRequest.Code)
	}

	if errBadRequest.HttpStatusCode != http.StatusBadRequest {
		t.Errorf("Invalid http status code [%v] not equal http status bad requests", errBadRequest.HttpStatusCode)
	}

	newError := fmt.Errorf("%s", BadRequest.ErrorCode)
	errBadRequest.SetError(newError)
	if errBadRequest.Error == nil {
		t.Error("error must not nil value")
	}

	if errBadRequest.Error != newError {
		t.Error("invalid error")
	}

	if errBadRequest.ErrorCode != BadRequest.ErrorCode {
		t.Error("invalid error code")
	}

	msgError := BadRequest.Message
	errBadRequest.SetMessage(msgError)
	if errBadRequest.Message != msgError {
		t.Error("invalid message error")
	}

	newMsgError := "new error"
	errBadRequest.SetMessage(newMsgError)
	if errBadRequest.Message != newMsgError {
		t.Error("invalid message")
	}

}
