package network

import (
	"bytes"
	"encoding/json"
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

	// fmt.Println(response.Body)
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
	fileName string,
	path string,
) (*json.Decoder, error) {
	var file *os.File
	file, _ = os.Open(path)
	defer file.Close()
	var body *bytes.Buffer = &bytes.Buffer{}
	var writer *multipart.Writer = multipart.NewWriter(body)
	for key, value := range data {
		writer.WriteField(key, value[0])
	}
	part, _ := writer.CreateFormFile("file", "file.txt")
	io.Copy(part, file)
	writer.Close()
	var response *http.Response
	var err error
	response, err = http.Post(method, writer.FormDataContentType(), body)
	return json.NewDecoder(response.Body), err
}

// func postDataRequest(url string, data map[string][]string) (interface{}, error) {
// 	var response *http.Response
// 	var err error
// 	var params Url.Values = Url.Values{}
// 	var counter int = 0
// 	for key, value := range data {
// 		params.Add(key, value[counter])
// 		counter++
// 	}
// 	response, err = http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
// 	return response, err
// }
