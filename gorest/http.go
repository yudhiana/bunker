package bunker

import (
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

	"github.com/yudhiana/bunker"

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

const (
	Auth   string = "Authorization"
	Bearer string = "Bearer "
)

type Requester struct {
	path    string
	BaseUrl string
	Method  string

	token     string
	Header    map[string][]string
	basicAuth map[string]string

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

func New(host string) *Requester {
	return &Requester{
		BaseUrl:   host,
		Header:    make(map[string][]string),
		FormData:  make(url.Values),
		QueryData: make(url.Values),
		basicAuth: make(map[string]string),
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

func (base *Requester) SetBasicAuth(username, password string) *Requester {
	base.basicAuth[username] = password
	return base
}

func (base *Requester) SetToken(token string) *Requester {
	base.token = token
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

func (base *Requester) initClient() *Requester {
	jar, errCookie := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if errCookie != nil {
		base.Errors = append(base.Errors, errCookie)
		return base
	}
	client := &http.Client{
		Jar:     jar,
		Timeout: base.TimeOut,
	}
	base.Client = client
	return base
}

func (base *Requester) SetTimeout(timeOut time.Duration) *Requester {
	base.TimeOut = timeOut
	return base
}

func (base *Requester) initRequestClient() bool {
	if base.initRequest().HaveError() {
		return false
	}

	if bunker.IsEmptyString(base.Request.Header.Get("Content-Type")) {
		base.Request.Header.Add("Content-Type", Json)
	}

	if bunker.IsEmptyString(base.Method) {
		base.Errors = append(base.Errors, errors.New("the request method has not been selected"))
		return false
	}

	if base.initClient().HaveError() {
		return false
	}

	return true
}

func (base *Requester) Do() *Requester {
	defer func() {
		if base.HaveError() {
			bunker.PrintErr(fmt.Sprint(base.Errors))
		}
	}()

	startTime := time.Now()
	if !base.initRequestClient() {
		return base
	}

	switch base.Method {
	case GET, HEAD, DELETE, OPTIONS, POST, PUT, PATCH:
		response, errRequestClient := base.Client.Do(base.Request)
		if errRequestClient != nil {
			base.Errors = append(base.Errors, errRequestClient)
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

func (base *Requester) initRequest() *Requester {
	request, errRequest := http.NewRequest(base.Method, base.BaseUrl, base.body)
	if errRequest != nil {
		base.Errors = append(base.Errors, errRequest)
		return base
	}
	if base.Ctx != nil {
		request = request.WithContext(base.Ctx)
	}
	if base.Header != nil {
		request.Header = base.Header
	}
	if base.basicAuth != nil {
		for user, pass := range base.basicAuth {
			request.SetBasicAuth(user, pass)
		}
	}
	if !bunker.IsEmptyString(base.token) {
		request.Header.Add(Auth, Bearer+base.token)
	}

	reqUrl := request.URL.Query()
	for param, values := range base.QueryData {
		for _, value := range values {
			reqUrl.Add(param, value)
		}
	}
	request.URL.RawQuery = reqUrl.Encode()
	if !bunker.IsEmptyString(base.path) {
		if !bunker.IsEmptyString(request.URL.Path) {
			request.URL.Path += base.path
		} else {
			request.URL.Path = base.path
		}
	}
	base.Request = request
	return base
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
		if !base.HaveError() {
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

func (base *Requester) AddPath(path string) *Requester {
	base.path = path
	return base
}

func (base *Requester) Query(query interface{}) *Requester {
	switch reflect.ValueOf(query).Kind() {
	case reflect.String:
		if input, isString := query.(string); isString {
			base.queryString(input)
		}
	case reflect.Map:
		switch input := query.(type) {
		case map[string]string:
			base.queryMap(input)
		default:
			base.Errors = append(base.Errors, fmt.Errorf("unsupported type of %T", reflect.TypeOf(query)))
		}
	default:
		base.Errors = append(base.Errors, fmt.Errorf("unsupported type of %T", reflect.TypeOf(query)))
	}
	return base
}

func (base *Requester) queryString(input string) *Requester {
	if values, errParse := url.ParseQuery(input); errParse != nil {
		base.Errors = append(base.Errors, errParse)
		return base
	} else {
		for param := range values {
			base.QueryData.Add(param, values.Get(param))
		}
	}
	return base
}

func (base *Requester) queryMap(input map[string]string) *Requester {
	for k, v := range input {
		base.QueryData.Add(k, v)
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
			base.Request.URL,
			base.headerToString(),
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

func (base *Requester) headerToString() (result string) {
	if base.Header != nil {
		header, _ := json.Marshal(base.Header)
		return string(header)
	}
	return
}
