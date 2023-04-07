package telegram

type MessageId int

type MessageThreadId int

type MessageAutoDeleteTimerChanged int

type Update struct {
	Update_id            int                 `json:"update_id"`
	Message              *Message            `json:"message,omitempty"`
	Edited_message       *Message            `json:"edited_message,omitempty"`
	Channel_post         *Message            `json:"channtl_post,omitempty"`
	Edited_channel_post  *Message            `json:"edited_channel_post,omitempty"`
	Inline_query         *InlineQuery        `json:"inline_query,omitempty"`
	Chosen_inline_result *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	Callback_query       *CallbackQuery      `json:"callback_query,omitempty"`
	Shipping_query       *ShippingQuery      `json:"shipping_query,omitempty"`
	Pre_checkout_query   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
	Poll                 *Poll               `json:"poll,omitempty"`
	Poll_answer          *PollAnswer         `json:"poll_answer,omitempty"`
	My_chat_member       *ChatMemberUpdated  `json:"my_chat_member,omitempty"`
	Chat_member          *ChatMemberUpdated  `json:"chat_member,omitempty"`
	Chat_join_request    *ChatJoinRequest    `json:"chat_join_request,omitempty"`
}

type User struct {
	Id                          int    `json:"id"`
	Is_bot                      bool   `json:"is_bot"`
	First_name                  string `json:"first_name"`
	Last_name                   string `json:"last_name,omitempty"`
	Username                    string `json:"usernmae,omitempty"`
	Language_code               string `json:"language_code,omitempty"`
	Is_premium                  bool   `json:"is_premium,omitempty"`
	Added_to_attachment_menu    bool   `json:"added_to_attachment_menu,omitempty"`
	Can_join_groups             bool   `json:"can_join_groups,omitempty"`
	Can_read_all_group_messages bool   `json:"can_read_all_group_messages,omitempty"`
	Supports_inline_queries     bool   `json:"supports_inline_queries,omitempty"`
}

type Message struct {
	Message_id                        *MessageId                     `json:"message_id"`
	Message_thread_id                 *MessageThreadId               `json:"message_thread_id,omitempty"`
	From                              *User                          `json:"from,omitempty"`
	Sender_chat                       *Chat                          `json:"sender_chat,omitempty"`
	Date                              int                            `json:"date"`
	Chat                              *Chat                          `json:"chat"`
	Forward_from                      *User                          `json:"forward_from,omitempty"`
	Forward_from_chat                 *Chat                          `json:"forward_from_chat,omitempty"`
	Forward_from_message_id           int                            `json:"forward_from_message_id,omitempty"`
	Forward_signature                 string                         `json:"forward_signature,omitempty"`
	Forward_sender_name               string                         `json:"forward_sender_name,omitempty"`
	Forward_date                      int                            `json:"faorward_date,omitempty"`
	Is_topic_message                  bool                           `json:"is_topic_message,omitempty"`
	Is_automatic_forward              bool                           `json:"is_automatic_forward,omitempty"`
	Reply_to_message                  *Message                       `json:"reply_to_message,omitempty"`
	Via_bot                           *User                          `json:"via_bot,omitempty"`
	Edit_date                         int                            `json:"edit_date,omitempty"`
	Has_protected_content             bool                           `json:"hat_protected_content,omitempty"`
	Media_group_id                    string                         `json:"media_group_id,omitempty"`
	Author_signature                  string                         `json:"author_signature,omitempty"`
	Text                              string                         `json:"text,omitempty"`
	Entities                          *[]MessageEntity               `json:"entities,omitempty"`
	Animation                         *Animation                     `json:"animation,omitempty"`
	Audio                             *Audio                         `json:"audio,omitempty"`
	Document                          *Document                      `json:"document,omitempty"`
	Photo                             *[]PhotoSize                   `json:"photo,omitempty"`
	Sticker                           *Sticker                       `json:"sticker,omitempty"`
	Video                             *Video                         `json:"video,omitempty"`
	Video_note                        *VideoNote                     `json:"video_note,omitempty"`
	Voice                             *Voice                         `json:"voice,omitempty"`
	Caption                           string                         `json:"caption,omitempty"`
	Caption_entities                  *[]MessageEntity               `json:"caption_entities,omitempty"`
	Has_media_spoiler                 bool                           `json:"mhas_media_spoiler,omitempty"`
	Contact                           *Contact                       `json:"contact,omitempty"`
	Dice                              *Dice                          `json:"dice,omitempty"`
	Game                              *Game                          `json:"game,omitempty"`
	Poll                              *Poll                          `json:"poll,omitempty"`
	Venue                             *Venue                         `json:"venue,omitempty"`
	Location                          *Location                      `json:"location,omitempty"`
	New_chat_members                  *[]User                        `json:"new_chat_members,omitempty"`
	Left_chat_member                  *User                          `json:"left_chat_member,omitempty"`
	New_chat_title                    string                         `json:"new_chat_title,omitempty"`
	New_chat_photo                    *[]PhotoSize                   `json:"new_chat_photo,omitempty"`
	Delete_chat_photo                 bool                           `json:"delete_chat_photo,omitempty"`
	Group_chat_created                bool                           `json:"group_chat_created,omitempty"`
	Supergroup_chat_created           bool                           `json:"supergroup_chat_created,omitempty"`
	Channel_chat_created              bool                           `json:"channel_chat_created,omitempty"`
	Message_auto_delete_timer_changed *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	Migrate_to_chat_id                int                            `json:"migrate_to_chat_id,omitempty"`
	Migrate_from_chat_id              int                            `json:"migrate_from_chat_id,omitempty"`
	Pinned_message                    *Message                       `json:"pinned_message,omitempty"`
	Invoice                           *Invoice                       `json:"invoice,omitempty"`
	Successful_payment                *SuccessfulPayment             `json:"successful_payment,omitempty"`
	User_shared                       *UserShared                    `json:"user_shared,omitempty"`
	Chat_shared                       *ChatShared                    `json:"chat_shared,omitempty"`
	Connected_website                 string                         `json:"connected_website,omitempty"`
	Write_access_allowed              *WriteAccessAllowed            `json:"write_access_allowed,omitempty"`
	Passport_data                     *PassportData                  `json:"passport_data,omitempty"`
	Proximity_alert_triggered         *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	Forum_topic_created               *ForumTopicCreated             `json:"forum_topic_created,omitempty"`
	Forum_topic_edited                *ForumTopicEdited              `json:"forum_topic_edited,omitempty"`
	Forum_topic_closed                *ForumTopicClosed              `json:"forum_topic_closed,omitempty"`
	Forum_topic_reopened              *ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`
	General_forum_topic_hidden        *GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`
	General_forum_topic_unhidden      *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`
	Video_chat_scheduled              *VideoChatScheduled            `json:"video_chat_sheduled,omitempty"`
	Video_chat_started                *VideoChatStarted              `json:"video_chat_started,omitempty"`
	Video_chat_ended                  *VideoChatEnded                `json:"video_chat_ended,omitempty"`
	Video_chat_participants_invited   *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`
	Web_app_data                      *WebAppData                    `json:"web_app_data,omitempty"`
	Reply_markup                      *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`
}

type Chat struct {
	Id                                      int
	Type                                    string
	Title                                   string
	Username                                string
	First_name                              string
	Last_name                               string
	Is_forum                                bool
	Photo                                   *ChatPhoto
	Active_usernames                        []string
	Emoji_status_custom_emoji_id            string
	Bio                                     string
	Has_private_forwards                    bool
	Has_restricted_voice_and_video_messages bool
	Join_to_send_messages                   bool
	Join_by_request                         bool
	Description                             string
	Invite_link                             string
	Pinned_message                          *Message
	Permissions                             *ChatPermissions
	Slow_mode_delay                         int
	Message_auto_delete_time                int
	Has_aggressive_anti_spam_enabled        bool
	Has_hidden_members                      bool
	Has_protected_content                   bool
	Sticker_set_name                        string
	Can_set_sticker_set                     bool
	Linked_chat_id                          int
	Location                                *ChatLocation
}

type InlineQuery struct {
	Id        string
	From      *User
	Query     string
	Offset    string
	Chat_type string
	Location  *Location
}

type ChosenInlineResult struct{}

type CallbackQuery struct {
	Id                string
	From              *User
	Message           *Message
	Inline_message_id string
	Chat_instance     string
	Data              string
	Game_short_name   string
}

type ShippingQuery struct {
	Id                 string
	From               *User
	Currency           string
	Total_amount       int
	Invoice_payload    string
	Shipping_option_id string
	Order_info         *OrderInfo
}

type OrderInfo struct {
	Name             string
	Phone_number     string
	Email            string
	Shipping_address *ShippingAddress
}

type ShippingAddress struct {
	Country_code string
	State        string
	City         string
	Street_line1 string
	Street_line2 string
	Post_code    string
}

type PreCheckoutQuery struct {
	Id                 string
	From               *User
	Currency           string
	Total_amount       int
	Invoice_payload    string
	Shipping_option_id string
	Order_info         *OrderInfo
}

type Poll struct {
	Id                      string
	Question                string
	Options                 *[]PollOption
	Total_voter_count       int
	Is_closed               bool
	Is_anonymous            bool
	Type                    string
	Allows_multiple_answers bool
	Correct_option_id       int
	Explanation             string
	Explanation_entities    *[]MessageEntity
	Open_period             int
	Close_date              int
}

type PollOption struct {
	Text        string
	Voter_count int
}

type PollAnswer struct {
	Poll_id    string
	User       *User
	Option_ids []int
}

type ChatMemberUpdated struct {
	Chat            *Chat
	From            *User
	Date            int
	Old_chat_member *ChatMember
	New_chat_member *ChatMember
	Invite_link     *ChatInviteLink
}

type ChatMember struct {
}

type ChatInviteLink struct {
	Invite_link                string
	Creator                    *User
	Creates_join_request       bool
	Is_primary                 bool
	Is_revoked                 bool
	Name                       string
	Expire_date                int
	Member_limit               int
	Pending_join_request_count int
}

type ChatJoinRequest struct {
	Chat         *Chat
	From         *User
	User_chat_id int
	Date         int
	Bio          string
	Invite_link  *ChatInviteLink
}

type MessageEntity struct {
	Type            string
	Offset          int
	Length          int
	Url             string
	User            *User
	Language        string
	Custom_emoji_id int
}

type Animation struct{}

type Audio struct{}

type Document struct{}

type PhotoSize struct{}

type Video struct{}

type VideoNote struct{}

type Voice struct{}

type Contact struct{}

type Dice struct{}

type Game struct{}

type Sticker struct{}

type Venue struct{}

type Location struct{}

type Invoice struct{}

type SuccessfulPayment struct{}

type UserShared struct{}

type ChatShared struct{}

type WriteAccessAllowed struct{}

type PassportData struct{}

type ProximityAlertTriggered struct{}

type ForumTopicCreated struct{}

type ForumTopicEdited struct{}

type ForumTopicClosed struct{}

type ForumTopicReopened struct{}

type GeneralForumTopicHidden struct{}

type GeneralForumTopicUnhidden struct{}

type VideoChatScheduled struct{}

type VideoChatStarted struct{}

type VideoChatEnded struct{}

type VideoChatParticipantsInvited struct{}

type WebAppData struct{}

type InlineKeyboardMarkup struct{}

type ChatLocation struct{}

type ChatPermissions struct{}

type ChatPhoto struct{}
