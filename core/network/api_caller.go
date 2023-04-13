package network

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	Url "net/url"
	"os"
	"strings"
)

type ApiCaller struct {
	requestUrl string
}

func NewApiCaller(requestUrl string) *ApiCaller {
	return &ApiCaller{requestUrl: requestUrl}
}

func (apiCaller *ApiCaller) GetNonParamsCallApi(method string) (*http.Response, error) {
	return getRequest(apiCaller.requestUrl + "/" + method)
}

func (apiCaller *ApiCaller) GetParamsCallApi(method string, params map[string]string) (*http.Response, error) {
	var request string = ""
	for key, value := range params {
		if len(request) == 0 {
			request += "?"
		}
		if len(request) > 2 {
			request += "&"
		}
		request += key + "=" + value
	}
	return getRequest(apiCaller.requestUrl + "/" + method + request)
}

func getRequest(url string) (*http.Response, error) {
	var response *http.Response
	var err error
	response, err = http.Get(url)
	return response, err
}

func (apiCaller *ApiCaller) PostCallApiData(method string, data map[string][]string) (*http.Response, error) {
	var response *http.Response
	var err error
	var params Url.Values = Url.Values{}
	for key, value := range data {
		params.Add(key, value[0])

	}
	response, err = http.Post(apiCaller.requestUrl+"/"+method, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
	return response, err
}

func (apiCaller *ApiCaller) PostCallApiFile(
	method string,
	data map[string][]string,
	fieldFileName string,
	path string,
) (*http.Response, error) {
	var file *os.File
	file, _ = os.Open(path)
	defer file.Close()
	var body *bytes.Buffer = &bytes.Buffer{}
	var writer *multipart.Writer = multipart.NewWriter(body)
	for key, value := range data {
		writer.WriteField(key, value[0])
	}
	part, _ := writer.CreateFormFile(fieldFileName, file.Name())
	io.Copy(part, file)
	writer.Close()
	var response *http.Response
	var err error
	response, err = http.Post(method, writer.FormDataContentType(), body)
	return response, err
}
