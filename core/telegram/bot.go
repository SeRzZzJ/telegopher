package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"telegopher/core/network"
)

type AllowedUpdates []string

type Telegram struct {
	//The telegram client
	BaseUrl   string
	BaseRoute string
	Token     string
	ApiCaller *network.ApiCaller
}

func NewTelegram(token string) *Telegram {
	const baseUrl string = "https://api.telegram.org"
	const baseRoute string = "/bot"
	return &Telegram{
		BaseUrl:   baseUrl,
		BaseRoute: baseRoute,
		Token:     token,
		ApiCaller: network.NewApiCaller(baseUrl + baseRoute + token),
	}
}

func unmarshalResponse(response *http.Response, err error, unmarshal interface{}) {
	var body []byte
	var tgError Error
	if err != nil {
		panic(err)
	}

	body, err = io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, unmarshal)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &tgError)
	if err != nil {
		panic(err)
	}

	if !tgError.Ok {
		panic(errors.New(strings.Join([]string{"Telegram:", strconv.Itoa(tgError.ErrorCode), tgError.Description}, " ")))
	}
}

func (telegram *Telegram) GetUpdates(
	offset int,
	limit int,
	timeout int,
	allowed_updates AllowedUpdates,
) *[]Update {
	var (
		response  *http.Response
		err       error
		unmurshal struct {
			Result []Update `json:"result"`
		}
	)

	response, err = telegram.ApiCaller.GetParamsCallApi("getUpdates", map[string]string{
		"offset":          strconv.Itoa(offset),
		"limit":           strconv.Itoa(limit),
		"timeout":         strconv.Itoa(timeout),
		"allowed_updates": "[\"" + strings.Join(allowed_updates, "\",\"") + "\"]",
	})

	unmarshalResponse(response, err, unmurshal)

	return &unmurshal.Result
}

func (telegram *Telegram) SetWebhook(url string, setWebhookOpts *SetWebhookOpts) bool {
	var (
		data     map[string][]string = make(map[string][]string)
		response *http.Response
		err      error
		res      struct {
			Result bool `json:"result"`
		}
	)

	data["url"] = []string{url}

	if setWebhookOpts.IpAddress != nil {
		data["ip_address"] = []string{*setWebhookOpts.IpAddress}
	}

	if setWebhookOpts.MaxConnections != nil {
		data["max_connections"] = []string{strconv.Itoa(*setWebhookOpts.MaxConnections)}
	}

	if setWebhookOpts.AllowedUpdates != nil {
		data["allowed_updates"] = []string{strings.Join(*setWebhookOpts.AllowedUpdates, ",")}
	}

	if setWebhookOpts.DropPendingUpdates != nil {
		data["drop_pending_updates"] = []string{strconv.FormatBool(*setWebhookOpts.DropPendingUpdates)}
	}

	if setWebhookOpts.SecretToken != nil {
		data["secret_token"] = []string{*setWebhookOpts.SecretToken}
	}

	if setWebhookOpts.Certificate != nil {
		switch setWebhookOpts.Certificate.TypeInputFile {
		case typeFile:
			response, err = telegram.ApiCaller.PostCallApiFile("setWebhook", data, "certificate", setWebhookOpts.Certificate.File)
			unmarshalResponse(response, err, &res)
			return res.Result
		default:
			data["certificate"] = []string{setWebhookOpts.Certificate.File}
		}
	}
	response, err = telegram.ApiCaller.PostCallApiData("setWebhook", data)
	unmarshalResponse(response, err, &res)
	return res.Result
}

type SetWebhookOpts struct {
	Certificate        *InputFile
	IpAddress          *string
	MaxConnections     *int
	AllowedUpdates     *AllowedUpdates
	DropPendingUpdates *bool
	SecretToken        *string
}

func (telegram *Telegram) DeleteWebhook(dropPendingUpdates bool) bool {
	var (
		data     map[string][]string = make(map[string][]string)
		response *http.Response
		err      error
		res      struct {
			Result bool `json:"result"`
		}
	)
	data["drop_pending_updates"] = []string{strconv.FormatBool(dropPendingUpdates)}
	response, err = telegram.ApiCaller.PostCallApiData("deleteWebhook", data)
	unmarshalResponse(response, err, &res)
	return res.Result
}

func (telegram *Telegram) GetWebhookInfo() *WebhookInfo {
	var (
		response *http.Response
		err      error
		res      struct {
			Result WebhookInfo `json:"result"`
		}
	)
	response, err = telegram.ApiCaller.GetNonParamsCallApi("getWebhookInfo")
	unmarshalResponse(response, err, &res)
	return &res.Result
}

func (telegram *Telegram) GetMe() *User {
	var (
		response *http.Response
		err      error
		res      struct {
			Result User `json:"result"`
		}
	)
	response, err = telegram.ApiCaller.GetNonParamsCallApi("getMe")
	unmarshalResponse(response, err, &res)
	return &res.Result
}

func (telegram *Telegram) LogOut() bool {
	var (
		response *http.Response
		err      error
		res      struct {
			Result bool `json:"result"`
		}
	)
	response, err = telegram.ApiCaller.GetNonParamsCallApi("logOut")
	unmarshalResponse(response, err, &res)
	return res.Result
}

func (telegram *Telegram) Close() bool {
	var (
		response *http.Response
		err      error
		res      struct {
			Result bool `json:"result"`
		}
	)
	response, err = telegram.ApiCaller.GetNonParamsCallApi("close")
	unmarshalResponse(response, err, &res)
	return res.Result
}

func (telegram *Telegram) SendMessage(chat_id int, text string, option struct{}) *Message {
	var (
		response *http.Response
		err      error
		body     []byte
		res      Message
	)
	response, err = telegram.ApiCaller.PostCallApiData("sendMessage",
		map[string][]string{"chat_id": {strconv.Itoa(chat_id)},
			"text": {text}})
	if err != nil {
		panic(err)
	}

	body, err = io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("error:", err)
	}
	return &res
}
