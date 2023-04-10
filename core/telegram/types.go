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
	Id                                      int              `json:"id"`
	Type                                    string           `json:"type"`
	Title                                   string           `json:"title,omitempty"`
	Username                                string           `json:"username,omitempty"`
	First_name                              string           `json:"first_name,omitempty"`
	Last_name                               string           `json:"last_name,omitempty"`
	Is_forum                                bool             `json:"is_forum,omitempty"`
	Photo                                   *ChatPhoto       `json:"photo,omitempty"`
	Active_usernames                        []string         `json:"active_usernames,omitempty"`
	Emoji_status_custom_emoji_id            string           `json:"emoji_status_custom_emoji_id,omitempty"`
	Bio                                     string           `json:"bio,omitempty"`
	Has_private_forwards                    bool             `json:"has_private_forwards,omitempty"`
	Has_restricted_voice_and_video_messages bool             `json:"has_restricted_voice_and_video_messages,omitempty"`
	Join_to_send_messages                   bool             `json:"join_to_send_messages,omitempty"`
	Join_by_request                         bool             `json:"join_by_request,omitempty"`
	Description                             string           `json:"description,omitempty"`
	Invite_link                             string           `json:"invite_link,omitempty"`
	Pinned_message                          *Message         `json:"pinned_message,omitempty"`
	Permissions                             *ChatPermissions `json:"permissions,omitempty"`
	Slow_mode_delay                         int              `json:"slow_mode_delay,omitempty"`
	Message_auto_delete_time                int              `json:"message_auto_delete_time,omitempty"`
	Has_aggressive_anti_spam_enabled        bool             `json:"has_aggressive_anti_spam_enabled,omitempty"`
	Has_hidden_members                      bool             `json:"has_hidden_members,omitempty"`
	Has_protected_content                   bool             `json:"has_protected_content,omitempty"`
	Sticker_set_name                        string           `json:"sticker_set_name,omitempty"`
	Can_set_sticker_set                     bool             `json:"can_set_sticker_set,omitempty"`
	Linked_chat_id                          int              `json:"linked_chat_id,omitempty"`
	Location                                *ChatLocation    `json:"location,omitempty"`
}

type InlineQuery struct {
	Id        string    `json:"id"`
	From      *User     `json:"from"`
	Query     string    `json:"query"`
	Offset    string    `json:"offset"`
	Chat_type string    `json:"chat_type,omitempty"`
	Location  *Location `json:"location,omitempty"`
}

type ChosenInlineResult struct {
	Result_id         string    `json:"result_id"`
	From              *User     `json:"from"`
	Location          *Location `json:"location,omitempty"`
	Inline_message_id string    `json:"inline_message_id,omitempty"`
	Query             string    `json:"query,omitempty"`
}

type CallbackQuery struct {
	Id                string   `json:"id"`
	From              *User    `json:"from"`
	Message           *Message `json:"message,omitempty"`
	Inline_message_id string   `json:"inline_message_id,omitempty"`
	Chat_instance     string   `json:"chat_instance"`
	Data              string   `json:"data,omitempty"`
	Game_short_name   string   `json:"game_short_name,omitempty"`
}

type ShippingQuery struct {
	Id               string           `json:"id"`
	From             *User            `json:"from"`
	Invoice_payload  string           `json:"invoice_payload"`
	Shipping_address *ShippingAddress `json:"shipping_address"`
}

type OrderInfo struct {
	Name             string           `json:"name,omitempty"`
	Phone_number     string           `json:"phone_number,omitempty"`
	Email            string           `json:"email,omitempty"`
	Shipping_address *ShippingAddress `json:"shipping_address,omitempty"`
}

type ShippingAddress struct {
	Country_code string `json:"country_code"`
	State        string `json:"state"`
	City         string `json:"city"`
	Street_line1 string `json:"street_line1"`
	Street_line2 string `json:"street_line2"`
	Post_code    string `json:"post_code"`
}

type PreCheckoutQuery struct {
	Id                 string     `json:"id"`
	From               *User      `json:"from"`
	Currency           string     `json:"currency"`
	Total_amount       int        `json:"total_amount"`
	Invoice_payload    string     `json:"invoice_payload"`
	Shipping_option_id string     `json:"shipping_option_id,omitempty"`
	Order_info         *OrderInfo `json:"order_info,omitempty"`
}

type Poll struct {
	Id                      string           `json:"id"`
	Question                string           `json:"question"`
	Options                 *[]PollOption    `json:"options"`
	Total_voter_count       int              `json:"total_voter_count"`
	Is_closed               bool             `json:"is_closed"`
	Is_anonymous            bool             `json:"is_anonymous"`
	Type                    string           `json:"type"`
	Allows_multiple_answers bool             `json:"allows_multiple_answers"`
	Correct_option_id       int              `json:"correct_option_id,omitempty"`
	Explanation             string           `json:"explanation,omitempty"`
	Explanation_entities    *[]MessageEntity `json:"explanation_entities,omitempty"`
	Open_period             int              `json:"open_period,omitempty"`
	Close_date              int              `json:"close_date,omitempty"`
}

type PollOption struct {
	Text        string `json:"text"`
	Voter_count int    `json:"voter_count"`
}

type PollAnswer struct {
	Poll_id    string `json:"poll_id"`
	User       *User  `json:"user"`
	Option_ids []int  `json:"option_ids"`
}

type ChatMemberUpdated struct {
	Chat            *Chat           `json:"chat"`
	From            *User           `json:"from"`
	Date            int             `json:"date"`
	Old_chat_member *ChatMember     `json:"old_chat_member"`
	New_chat_member *ChatMember     `json:"new_chat_member"`
	Invite_link     *ChatInviteLink `json:"invite_link,omitempty"`
}

type ChatMember struct {
	Status                    string `json:"status"`
	User                      *User  `json:"user"`
	Is_anonymous              bool   `json:"is_anonymous,omitempty"`
	Custom_title              string `json:"custom_title,omitempty"`
	Can_be_edited             bool   `json:"can_be_edited,omitempty"`
	Can_manage_chat           bool   `json:"can_manage_chat,omitempty"`
	Can_delete_messages       bool   `json:"can_delete_messages,omitempty"`
	Can_manage_video_chats    bool   `json:"can_manage_video_chats,omitempty"`
	Can_restrict_members      bool   `json:"can_restrict_members,omitempty"`
	Can_promote_members       bool   `json:"can_promote_members,omitempty"`
	Can_change_info           bool   `json:"can_change_info,omitempty"`
	Can_invite_users          bool   `json:"can_invite_users,omitempty"`
	Can_post_messages         bool   `json:"can_post_messages,omitempty"`
	Can_edit_messages         bool   `json:"can_edit_messages,omitempty"`
	Can_pin_messages          bool   `json:"can_pin_messages,omitempty"`
	Can_manage_messages       bool   `json:"can_manage_messages,omitempty"`
	Can_send_messages         bool   `json:"can_send_messages,omitempty"`
	Can_send_audios           bool   `json:"can_send_audios,omitempty"`
	Can_send_documents        bool   `json:"can_send_documents,omitempty"`
	Can_send_photos           bool   `json:"can_send_photos,omitempty"`
	Can_send_videos           bool   `json:"can_send_videos,omitempty"`
	Can_send_video_notes      bool   `json:"can_send_video_notes,omitempty"`
	Can_send_voice_notes      bool   `json:"can_send_voice_notes,omitempty"`
	Can_send_polls            bool   `json:"can_send_polls,omitempty"`
	Can_send_other_messages   bool   `json:"can_send_other_messages,omitempty"`
	Can_add_web_page_previews bool   `json:"can_add_web_page_previews,omitempty"`
	Can_manage_topics         bool   `json:"can_manage_topics,omitempty"`
	Until_date                int    `json:"until_date,omitempty"`
}

type ChatInviteLink struct {
	Invite_link                string `json:"invite_link"`
	Creator                    *User  `json:"creator"`
	Creates_join_request       bool   `json:"creates_join_request"`
	Is_primary                 bool   `json:"is_primary"`
	Is_revoked                 bool   `json:"is_revoked"`
	Name                       string `json:"name,omitempty"`
	Expire_date                int    `json:"expire_date,omitempty"`
	Member_limit               int    `json:"member_limit,omitempty"`
	Pending_join_request_count int    `json:"pending_join_request_count,omitempty"`
}

type ChatJoinRequest struct {
	Chat         *Chat           `json:"chat"`
	From         *User           `json:"from"`
	User_chat_id int             `json:"user_chat_id"`
	Date         int             `json:"date"`
	Bio          string          `json:"bio,omitempty"`
	Invite_link  *ChatInviteLink `json:"invite_link,omitempty"`
}

type MessageEntity struct {
	Type            string `json:"type"`
	Offset          int    `json:"offset"`
	Length          int    `json:"length"`
	Url             string `json:"url,omitempty"`
	User            *User  `json:"user,omitempty"`
	Language        string `json:"language,omitempty"`
	Custom_emoji_id int    `json:"custom_emoji_id,omitempty"`
}

type Animation struct {
	File_id        string     `json:"file_id"`
	File_unique_id string     `json:"file_unique_id"`
	Width          int        `json:"width"`
	Height         int        `json:"height"`
	Duration       int        `json:"duration"`
	Thumbnail      *PhotoSize `json:"thumbnail,omitempty"`
	File_name      string     `json:"file_name,omitempty"`
	Mime_type      string     `json:"mime_type,omitempty"`
	File_size      int        `json:"file_size,omitempty"`
}

type Audio struct {
	File_id        string `json:"file_id"`
	File_unique_id string `json:"file_unique_id"`
	Duration       int    `json:"duration"`
	Performer      string `json:"performer,omitempty"`
	Title          string `json:"title,omitempty"`
	File_name      string `json:"file_name,omitempty"`
	Mime_type      string `json:"mime_type,omitempty"`
	File_size      int    `json:"file_size,omitempty"`
}

type Document struct {
	File_id        string     `json:"file_id"`
	File_unique_id string     `json:"file_unique_id"`
	Thumbnail      *PhotoSize `json:"thumbnail,omitempty"`
	File_name      string     `json:"file_name,omitempty"`
	Mime_type      string     `json:"mime_type,omitempty"`
	File_size      int        `json:"file_size,omitempty"`
}

type PhotoSize struct {
	File_id        string `json:"file_id"`
	File_unique_id string `json:"file_unique_id"`
	Width          int    `json:"width"`
	File_size      int    `json:"file_size,omitempty"`
}

type Video struct {
	File_id        string     `json:"file_id"`
	File_unique_id string     `json:"file_unique_id"`
	Width          int        `json:"width"`
	Height         int        `json:"height"`
	Duration       int        `json:"duration"`
	Thumbnail      *PhotoSize `json:"thumbnail,omitempty"`
	File_name      string     `json:"file_name,omitempty"`
	Mime_type      string     `json:"mime_type,omitempty"`
	File_size      int        `json:"file_size,omitempty"`
}

type VideoNote struct {
	File_id        string     `json:"file_id"`
	File_unique_id string     `json:"file_unique_id"`
	Length         int        `json:"lenght"`
	Duration       int        `json:"duration"`
	Thumbnail      *PhotoSize `json:"thumbnail,omitempty"`
	File_size      int        `json:"file_size,omitempty"`
}

type Voice struct {
	File_id        string `json:"file_id"`
	File_unique_id string `json:"file_unique_id"`
	Duration       int    `json:"duration"`
	Mime_type      string `json:"mime_type,omitempty"`
	File_size      int    `json:"file_size,omitempty"`
}

type Contact struct {
	Phone_number string `json:"phone_number"`
	First_name   string `json:"first_name"`
	Last_name    string `json:"last_name,omitempty"`
	User_id      int    `json:"user_id,omitempty"`
	Vcard        string `json:"vcard,omitempty"`
}

type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

type Game struct {
	Title         string           `json:"title"`
	Description   string           `json:"description"`
	Photo         []*PhotoSize     `json:"photo"`
	Text          string           `json:"text,omitempty"`
	Text_Entities []*MessageEntity `json:"text_entities,omitempty"`
	Animation     *Animation       `json:"animation,omitempty"`
}

type Sticker struct {
	File_id           string        `json:"file_id"`
	File_unique_id    string        `json:"file_unique_id"`
	Type              string        `json:"type"`
	Width             int           `json:"width"`
	Height            int           `json:"height"`
	Is_animated       bool          `json:"is_animated"`
	Is_video          bool          `json:"is_video"`
	Thumbnail         *PhotoSize    `json:"thumbnail,omitempty"`
	Emoji             string        `json:"emoji,omitempty"`
	Set_name          string        `json:"set_name,omitempty"`
	Premium_animation *File         `json:"premium_animation,omitempty"`
	Mask_position     *MaskPosition `json:"mask_position,omitempty"`
	Custom_emoji_id   string        `json:"custom_emoji_id,omitempty"`
	Needs_repainting  bool          `json:"needs_repainting,omitempty"`
	File_size         int           `json:"file_size,omitempty"`
}

type File struct{}

type MaskPosition struct{}

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
