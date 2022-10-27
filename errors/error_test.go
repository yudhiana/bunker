package errors_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/yudhiana/bunker/errors"
)

func TestNew(t *testing.T) {

	errBunker := errors.New(errors.StatusBadRequest)

	if errBunker.Code != errors.StatusBadRequest {
		t.Errorf("Invalid status code [%v] not equal bad requests", errBunker.Code)
	}

	if errBunker.HttpStatusCode != http.StatusBadRequest {
		t.Errorf("Invalid http status code [%v] not equal http status bad requests", errBunker.HttpStatusCode)
	}

	newError := fmt.Errorf("%s", errors.BadRequest.ErrorCode)
	errBunker.SetError(newError)
	if errBunker.Error == nil {
		t.Error("error must not nil value")
	}

	if errBunker.Error != newError {
		t.Error("invalid error")
	}

	if errBunker.ErrorCode != errors.BadRequest.ErrorCode {
		t.Error("invalid error code")
	}

	msgError := errors.BadRequest.Message
	errBunker.SetMessage(msgError)
	if errBunker.Message != msgError {
		t.Error("invalid message error")
	}

}
