package bunker

import (
	"bunker"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"

	"golang.org/x/net/publicsuffix"
)

// HTTP methods we support
const (
	POST    string = http.MethodPost
	GET     string = http.MethodGet
	HEAD    string = http.MethodHead
	PUT     string = http.MethodPut
	DELETE  string = http.MethodDelete
	PATCH   string = http.MethodPatch
	OPTIONS string = http.MethodOptions
)

// ShortContentTypes defines some short content types.
const (
	Json       string = "application/json"
	UrlEncoded string = "application/x-www-form-urlencoded"
	Form       string = "application/x-www-form-urlencoded"
	FormData   string = "application/x-www-form-urlencoded"
)

type Requester struct {
	URL    string
	Host   string
	Method string

	Header map[string][]string

	FormData  url.Values
	QueryData url.Values

	Client   *http.Client
	Response *http.Response
	Request  *http.Request

	TimeOut time.Duration

	Errors []error

	Ctx context.Context

	Debug bool

	timeRequest time.Duration
	timeIn      time.Time

	body io.Reader
}

func New(url string) *Requester {
	return &Requester{
		URL: url,
	}
}

// Header := map[string][]string{
// 			"Accept-Encoding": {"gzip, deflate"},
// 			"Accept-Language": {"en-us"},
// 			"Foo": {"Bar", "two"},
// 		}

func (base *Requester) HaveError() bool {
	return len(base.Errors) != 0
}

func (base *Requester) SetHeader(param string, values ...string) *Requester {
	base.Header[param] = values
	return base
}

func (base *Requester) SetHeaders(headers interface{}) *Requester {
	switch kind := reflect.ValueOf(headers); kind.Kind() {
	case reflect.Map:
		switch typeOf := kind.Interface().(type) {
		case map[string][]string:
			base.setMapStringHeaders(typeOf)
		case map[string]interface{}:
			base.setMapAnyHeaders(typeOf)
		default:
			base.Errors = append(base.Errors, fmt.Errorf("unsupported type of %T", reflect.TypeOf(typeOf)))
		}
	default:
		base.Errors = append(base.Errors, fmt.Errorf("unsupported type of %T", reflect.TypeOf(kind)))
	}
	return base
}

func (base *Requester) setMapStringHeaders(headers map[string][]string) *Requester {
	base.Header = headers
	return base
}

func (base *Requester) setMapAnyHeaders(headers map[string]interface{}) *Requester {
	for k, v := range headers {
		switch value := v.(type) {
		case []string:
			base.Header[k] = value
		default:
			base.Errors = append(base.Errors, fmt.Errorf("unsupported type of %T", reflect.TypeOf(value)))
		}
	}
	return base
}

func (base *Requester) SetContentType(contentType string) *Requester {
	switch bunker.IsEmptyString(contentType) {
	case true:
		base.Header["Content-Type"] = append(base.Header["Content-Type"], contentType)
	default:
		base.Header["Content-Type"] = append(base.Header["Content-Type"], Json)
	}
	return base
}

func (base *Requester) initClient() *http.Client {
	jar, _ := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	return &http.Client{
		Jar:     jar,
		Timeout: base.TimeOut,
	}
}

func (base *Requester) SetTimeout(timeOut time.Duration) *Requester {
	base.TimeOut = timeOut
	return base
}

func (base *Requester) Do() *Requester {
	startTime := time.Now()
	switch base.Method {
	case GET, HEAD, DELETE, OPTIONS:
		request, errGet := base.initRequest(nil)
		if errGet != nil {
			base.Errors = append(base.Errors, errGet)
			return base
		}
		response, errClient := base.initClient().Do(request)
		if errClient != nil {
			base.Errors = append(base.Errors, errClient)
			return base
		}
		base.Response = response
	case POST, PUT, PATCH:
		request, errGet := base.initRequest(base.body)
		if errGet != nil {
			base.Errors = append(base.Errors, errGet)
			return base
		}
		response, errClient := base.initClient().Do(request)
		if errClient != nil {
			base.Errors = append(base.Errors, errClient)
			return base
		}
		base.Response = response
	default:
		base.Errors = append(base.Errors, fmt.Errorf("unsupported method of %s", base.Method))
	}
	base.timeRequest = time.Since(startTime)
	base.timeIn = time.Now()
	base.printDebug()
	return base
}

func (base *Requester) initRequest(body io.Reader) (request *http.Request, err error) {
	request, err = http.NewRequest(GET, base.URL, body)
	base.Request = request
	if err != nil {
		return
	}
	if base.Ctx != nil {
		return request.WithContext(base.Ctx), nil
	}
	if base.Header != nil {
		request.Header = base.Header
	}
	return
}

func (base *Requester) Get() *Requester {
	base.Method = GET
	return base
}

func (base *Requester) Post() *Requester {
	base.Method = POST
	return base
}

func (base *Requester) Put() *Requester {
	base.Method = PUT
	return base
}

func (base *Requester) Patch() *Requester {
	base.Method = PATCH
	return base
}

func (base *Requester) Delete() *Requester {
	base.Method = DELETE
	return base
}

func (base *Requester) Head() *Requester {
	base.Method = HEAD
	return base
}

func (base *Requester) Options() *Requester {
	base.Method = OPTIONS
	return base
}

func (base *Requester) SetDebug(debug bool) *Requester {
	base.Debug = debug
	return base
}

func (base *Requester) printDebug() {
	if base.Debug || os.Getenv("debug") == "true" {
		debugMessage := base.debug()
		if base.HaveError() {
			bunker.PrintErr(fmt.Sprint(base.Errors))
		} else {
			bunker.LogInfo(debugMessage)
		}
	}
}

func (base *Requester) readAll(input io.ReadCloser) []byte {
	if input != nil {
		result, err := io.ReadAll(input)
		if err != nil {
			base.Errors = append(base.Errors, err)
			return nil
		}
		return result
	}
	return nil
}

func (base *Requester) readBodyRequest() []byte {
	if base.Request != nil && base.Request.Body != nil {
		result, err := io.ReadAll(base.Request.Body)
		if err != nil {
			base.Errors = append(base.Errors, err)
			return nil
		}
		base.Request.Body = io.NopCloser(bytes.NewBuffer(result))
		return result
	}
	return nil
}

func (base *Requester) SetPayload(body interface{}) *Requester {
	switch reflect.ValueOf(body).Kind() {
	case reflect.Map:
		base.setBodyMap(body)
	case reflect.Struct:
		base.setBodyStruct(body)
	case reflect.String:
		base.setBodyString(body)
	default:
		base.Errors = append(base.Errors, fmt.Errorf("unsupported type of %T", reflect.TypeOf(body)))
	}

	return base
}

func (base *Requester) setBodyMap(input interface{}) {
	switch data := input.(type) {
	case map[string]interface{}, map[string]string:
		base.setBodyMapString(data)
	default:
		base.Errors = append(base.Errors, fmt.Errorf("unsupported type of %T", reflect.TypeOf(data)))
	}
}

func (base *Requester) setBodyMapString(input interface{}) *Requester {
	dataMarshal, errMarshal := json.Marshal(input)
	if errMarshal != nil {
		base.Errors = append(base.Errors, errMarshal)
		return base
	}

	base.body = strings.NewReader(string(dataMarshal))
	return base
}

func (base *Requester) setBodyStruct(input interface{}) *Requester {
	dataMarshal, errMarshal := json.Marshal(input)
	if errMarshal != nil {
		base.Errors = append(base.Errors, errMarshal)
		return base
	}

	base.body = strings.NewReader(string(dataMarshal))
	return base
}

func (base *Requester) setBodyString(input interface{}) *Requester {
	switch data := input.(type) {
	case string:
		base.body = strings.NewReader(data)
	default:
		base.Errors = append(base.Errors, fmt.Errorf("unsupported type of %T", reflect.TypeOf(data)))
	}
	return base
}

func (base *Requester) debug() string {
	if base.Request != nil {
		return fmt.Sprintf(
			`
	REQUEST
	==========================================================
	%s / %s / %s
	URL             : %s
	HEADERS         : %v
	BODY REQUEST    : 
	%v

	RESPONSE
	==========================================================
	STATUS          : %s
	RECEIVED AT     : %v
	RESPONSE TIME   : %v
	
	BODY RESPONSE   :
	%v

	==========================================================
	`,
			time.Now().Format("2006/01/02 15:04:05"),
			base.Method,
			base.Request.Proto,
			base.URL,
			base.Header,
			string(base.readBodyRequest()),
			base.Response.Status,
			base.timeIn.Format(time.RFC1123),
			base.timeRequest.Seconds(),
			string(base.readAll(base.Response.Body)),
		)
	} else {
		base.Errors = append(base.Errors, errors.New("can't debug before request"))
		return ""
	}
}
