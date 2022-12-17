package errors_test

import (
	appErr "bunker/errors"
	"fmt"
	"net/http"
	"testing"
)

func TestNewError(t *testing.T) {

	errBadRequest := appErr.New(appErr.StatusBadRequest)

	if errBadRequest.Code != appErr.StatusBadRequest {
		t.Errorf("Invalid status code [%v] not equal bad requests", errBadRequest.Code)
	}

	if errBadRequest.HttpStatusCode != http.StatusBadRequest {
		t.Errorf("Invalid http status code [%v] not equal http status bad requests", errBadRequest.HttpStatusCode)
	}

	newError := fmt.Errorf("%s", appErr.BadRequest.ErrorCode)
	errBadRequest.SetError(newError)
	if errBadRequest.Error == nil {
		t.Error("error must not nil value")
	}

	if errBadRequest.Error != newError {
		t.Error("invalid error")
	}

	if errBadRequest.ErrorCode != appErr.BadRequest.ErrorCode {
		t.Error("invalid error code")
	}

	msgError := appErr.BadRequest.Message
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
