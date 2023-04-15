package telegopher

import (
	"math"
	"telegopher/core/telegram"
)

type Telegopher struct {
	telegram *telegram.Telegram
	handlers []Handler
	//middlewares []func()
}

func NewTelegopher(token string, tgOpts *telegram.TelegramOpts) *Telegopher {
	return &Telegopher{telegram: telegram.NewTelegram(token, tgOpts)}
}

func (telegopher *Telegopher) UseMiddleware() {

}

func (telegopher *Telegopher) AddHandler(updateType string, trigger string, handler func()) {

}
func (telegopher *Telegopher) AddHandlerMessage(trigger *Trigger, handler func(ctx *Context)) {
	telegopher.handlers = append(telegopher.handlers, *NewHandler("message", trigger, handler))
}
func (telegopher *Telegopher) AddHandlerEditedMessage(trigger string, handler func()) {

}
func (telegopher *Telegopher) AddHandlerChannelPost(trigger string, handler func()) {

}
func (telegopher *Telegopher) AddHandlerEditedChannelPost(trigger string, handler func()) {

}
func (telegopher *Telegopher) AddHandlerInlineQuery(trigger string, handler func()) {

}
func (telegopher *Telegopher) AddHandlerChosenInlineResult(trigger string, handler func()) {

}
func (telegopher *Telegopher) AddHandlerCallbackQuery(trigger string, handler func()) {

}
func (telegopher *Telegopher) AddHandlerShippingQuery(trigger string, handler func()) {

}
func (telegopher *Telegopher) AddHandlerPreCheckoutQuery(trigger string, handler func()) {

}
func (telegopher *Telegopher) AddHandlerPoll(trigger string, handler func()) {

}
func (telegopher *Telegopher) AddHandlerPollAnswer(trigger string, handler func()) {

}
func (telegopher *Telegopher) AddHandlerMyChatMember(trigger string, handler func()) {

}
func (telegopher *Telegopher) AddHandlerChatMember(trigger string, handler func()) {

}
func (telegopher *Telegopher) AddHandlerChatJoinRequest(trigger string, handler func()) {

}

func (telegopher *Telegopher) StartLongPolling(botOpts *BotOpts) {
	var (
		offset         int = -1
		limit          int
		timeout        int
		allowedUpdates telegram.AllowedUpdates
		updates        *[]telegram.Update
	)
	if l := botOpts.Limit; l != 0 {
		limit = l
	} else {
		limit = 100
	}
	if t := botOpts.Timeout; t != 0 {
		timeout = t
	} else {
		timeout = 10
	}
	if aUpdates := botOpts.AllowedUpdates; aUpdates != nil {
		allowedUpdates = aUpdates
	} else {
		allowedUpdates = telegram.AllowedUpdates{
			"message",
			"edited_message",
			"channel_post",
			"edited_channel_post",
			"inline_query",
			"chosen_inline_result",
			"callback_query",
			"shipping_query",
			"pre_checkout_query",
			"poll",
			"poll_answer",
			"my_chat_member",
			"chat_member",
			"chat_join_request",
		}
	}
	for {
		updates = telegopher.telegram.GetUpdates(offset, limit, timeout, allowedUpdates)
		for _, update := range *updates {
			// for _, middleware := range telegopher.middlewares {
			// 	middleware()
			// }
			for _, handler := range telegopher.handlers {
				switch handler.updateType {
				case "message":
					switch handler.triggerType {
					case "text":
						if update.Message.Text == handler.trigger {
							handler.handlerfn(NewContext(telegopher.telegram, &update))
						}
					}
				case "edited_message":
					handler.handlerfn(NewContext(telegopher.telegram, &update))
				case "channel_post":
					handler.handlerfn(NewContext(telegopher.telegram, &update))
				case "edited_channel_post":
					handler.handlerfn(NewContext(telegopher.telegram, &update))
				case "inline_query":
					handler.handlerfn(NewContext(telegopher.telegram, &update))
				case "chosen_inline_result":
					handler.handlerfn(NewContext(telegopher.telegram, &update))
				case "callback_query":
					handler.handlerfn(NewContext(telegopher.telegram, &update))
				case "shipping_query":
					handler.handlerfn(NewContext(telegopher.telegram, &update))
				case "pre_checkout_query":
					handler.handlerfn(NewContext(telegopher.telegram, &update))
				case "poll":
					handler.handlerfn(NewContext(telegopher.telegram, &update))
				case "poll_answer":
					handler.handlerfn(NewContext(telegopher.telegram, &update))
				case "my_chat_member":
					handler.handlerfn(NewContext(telegopher.telegram, &update))
				case "chat_join_request":
					handler.handlerfn(NewContext(telegopher.telegram, &update))
				default:
					handler.handlerfn(NewContext(telegopher.telegram, &update))
				}
			}
			offset = int(math.Max(float64(offset), float64(update.Update_id))) + 1
		}
	}

}

type BotOpts struct {
	Limit          int
	Timeout        int
	AllowedUpdates telegram.AllowedUpdates
}

type Handler struct {
	updateType  string
	triggerType string
	trigger     string
	handlerfn   func(ctx *Context)
}

func NewHandler(updateType string, trigger *Trigger, handlerfn func(ctx *Context)) *Handler {
	return &Handler{updateType: updateType, triggerType: trigger.triggerType, trigger: trigger.trigger, handlerfn: handlerfn}
}

type Context struct {
	telegram *telegram.Telegram
	update   *telegram.Update
}

func NewContext(telegram *telegram.Telegram, update *telegram.Update) *Context {
	return &Context{telegram: telegram, update: update}
}

func (context *Context) Reply(text string, option struct{}) *telegram.Message {
	return context.telegram.SendMessage(context.update.Message.Chat.Id, text, option)
}

type Trigger struct {
	triggerType string
	trigger     string
}

func NewTrigger(triggerType string, trigger string) *Trigger {
	return &Trigger{triggerType: triggerType, trigger: trigger}
}
