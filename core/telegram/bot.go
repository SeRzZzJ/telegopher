package telegram

import (
	"encoding/json"
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

func (telegram *Telegram) GetUpdates(offset int,
	limit int,
	timeout int,
	allowed_updates AllowedUpdates) *[]Update {
	var (
		response *http.Response
		err      error
		body     []byte
		res      Response
	)

	response, err = telegram.ApiCaller.GetParamsCallApi("getUpdates", map[string]string{"offset": strconv.Itoa(offset),
		"limit":           strconv.Itoa(limit),
		"timeout":         strconv.Itoa(timeout),
		"allowed_updates": "[\"" + strings.Join(allowed_updates, "\",\"") + "\"]"})
	if err != nil {
		panic(err)
	}

	body, err = io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		panic(err)
	}

	return &res.Result
}

type Response struct {
	Result []Update `json:"result"`
}

func (telegram *Telegram) GetMe() *User {
	//var response *http.Response
	var err error
	_, err = telegram.ApiCaller.GetNonParamsCallApi("getMe")
	if err != nil {
		panic(err)
	}
	var botInfo *User
	//response.Decoder(&botInfo)
	return botInfo
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

//	struct {
//		message,
//		edited_message,
//		channel_post,
//		edited_channel_post,
//		inline_query,
//		chosen_inline_result,
//		callback_query,
//		shipping_query,
//		pre_checkout_query,
//		poll,
//		my_chat_member,
//		chat_member,
//		chat_join_request bool
//	}
// func NewAllowedUpdates(message bool,
// 	edited_message bool,
// 	channel_post bool,
// 	edited_channel_post bool,
// 	inline_query bool,
// 	chosen_inline_result bool,
// 	callback_query bool,
// 	shipping_query bool,
// 	pre_checkout_query bool,
// 	poll bool,
// 	my_chat_member bool,
// 	chat_member bool,
// 	chat_join_request bool) *AllowedUpdates {
// 	return &AllowedUpdates{message: message,
// 		edited_message:       edited_message,
// 		channel_post:         channel_post,
// 		edited_channel_post:  edited_channel_post,
// 		inline_query:         inline_query,
// 		chosen_inline_result: chosen_inline_result,
// 		callback_query:       callback_query,
// 		shipping_query:       shipping_query,
// 		pre_checkout_query:   pre_checkout_query,
// 		poll:                 poll,
// 		my_chat_member:       my_chat_member,
// 		chat_member:          chat_member,
// 		chat_join_request:    chat_join_request}
// }

// func (allowed_updates *AllowedUpdates) getAllowedUpdates() []string {
// 	var allowed_updates_array []string
// 	for _, value := range allowed_updates {
// 		allowed_updates_array = append(allowed_updates_array[:], value)
// 	}
// 	return allowed_updates_array
// }
