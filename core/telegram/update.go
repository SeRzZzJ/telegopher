package telegram

type Updates struct {
	UpdateMessage []*UpdateMessage `json:"message"`
}

type UpdateMessage struct {
	// update_id int `json:"update_id"`
	Message *Message `json:"message"`
}

// func (updateMessage *UpdateMessage) GetUpdateId() int {
// 	return updateMessage.update_id
// }

// Edited_message       *Message
// Channel_post         *Message
// Edited_channel_post  *Message
// Inline_query         *InlineQuery
// Chosen_inline_result *ChosenInlineResult
// Callback_query       *CallbackQuery
// Shipping_query       *ShippingQuery
// Pre_checkout_query   *PreCheckoutQuery
// Poll                 *Poll
// Poll_answer          *PollAnswer
// My_chat_member       *ChatMemberUpdated
// Chat_member          *ChatMemberUpdated
// Chat_join_request    *ChatJoinRequest
type UpdateEditedMessage struct {
	update_id int
	*Message
}

func (updateEditedMessage *UpdateEditedMessage) GetUpdateId() int {
	return updateEditedMessage.update_id
}

type UpdateChannelPost struct {
	update_id int
	*Message
}

func (updateChannelPost *UpdateChannelPost) GetUpdateId() int {
	return updateChannelPost.update_id
}

type UpdateEditedChannelPost struct {
	update_id int
	*Message
}

func (updateEditedChannelPost *UpdateEditedChannelPost) GetUpdateId() int {
	return updateEditedChannelPost.update_id
}

type UpdateInlineQuery struct {
	update_id int
	*InlineQuery
}

func (updateInlineQuery *UpdateInlineQuery) GetUpdateId() int {
	return updateInlineQuery.update_id
}

type UpdateChosenInlineResult struct {
	update_id int
	*ChosenInlineResult
}

func (updateChosenInlineResult *UpdateChosenInlineResult) GetUpdateId() int {
	return updateChosenInlineResult.update_id
}

type UpdateCallbackQuery struct {
	update_id int
	*CallbackQuery
}

func (updateCallbackQuery *UpdateCallbackQuery) GetUpdateId() int {
	return updateCallbackQuery.update_id
}

type UpdateShippingQuery struct {
	update_id int
	*ShippingQuery
}

func (updateShippingQuery *UpdateShippingQuery) GetUpdateId() int {
	return updateShippingQuery.update_id
}

type UpdatePreCheckoutQuery struct {
	update_id int
	*PreCheckoutQuery
}

func (updatePreCheckoutQuery *UpdatePreCheckoutQuery) GetUpdateId() int {
	return updatePreCheckoutQuery.update_id
}

type UpdatePoll struct {
	update_id int
	*Poll
}

func (updatePoll *UpdatePoll) GetUpdateId() int {
	return updatePoll.update_id
}

type UpdateMyChatMember struct {
	update_id int
	*ChatMemberUpdated
}

func (updateMyChatMember *UpdateMyChatMember) GetUpdateId() int {
	return updateMyChatMember.update_id
}

type UpdateChatMember struct {
	update_id int
	*Message
}

func (updateChatMember *UpdateChatMember) GetUpdateId() int {
	return updateChatMember.update_id
}

type UpdateChatJoinRequest struct {
	update_id int
	*ChatJoinRequest
}

func (updateChatJoinRequest *UpdateChatJoinRequest) GetUpdateId() int {
	return updateChatJoinRequest.update_id
}
