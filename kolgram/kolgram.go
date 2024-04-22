package kolgram

import (
	methods "github.com/koltypka/kolgram/kolgram/methods"
)

type Handler struct {
	token string
}

func New(token string) Handler {
	return Handler{token}
}

func (Handler *Handler) GetUpdates() methods.GetUpdates {
	return methods.New(Handler.token).GetUpdates()
}

func (Handler *Handler) SetWebhook(url string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("url", url)

	return Telegram.Post("/setWebhook")
}

func (Handler *Handler) DeleteWebhook() map[string]interface{} {
	Telegram := methods.New(Handler.token)
	return Telegram.Post("/deleteWebhook")
}

func (Handler *Handler) GetWebhookInfo() methods.GetWebhookInfo {
	return methods.New(Handler.token).GetWebhookInfo()
}

func (Handler *Handler) GetMe() methods.GetMe {
	return methods.New(Handler.token).GetMe()
}

func (Handler *Handler) LogOut() map[string]interface{} {
	Telegram := methods.New(Handler.token)
	return Telegram.Post("/logOut")
}

func (Handler *Handler) Close() map[string]interface{} {
	Telegram := methods.New(Handler.token)
	return Telegram.Post("/close")
}

func (Handler *Handler) SendMessage(chat_id, text string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("text", text)
	return Telegram.Post("/sendMessage")
}

func (Handler *Handler) ForwardMessage(chat_id, from_chat_id, message_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("from_chat_id", from_chat_id)
	Telegram.AddParam("message_id", message_id)
	return Telegram.Post("/forwardMessage")
}

func (Handler *Handler) CopyMessage(chat_id, from_chat_id, message_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("from_chat_id", from_chat_id)
	Telegram.AddParam("message_id", message_id)
	return Telegram.Post("/copyMessage")
}

func (Handler *Handler) SendPhoto(chat_id, photo string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("photo", photo)
	return Telegram.Post("/sendPhoto")
}

func (Handler *Handler) SendAudio(chat_id, audio string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("audio", audio)
	return Telegram.Post("/sendAudio")
}

func (Handler *Handler) SendDocument(chat_id, document string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("document", document)
	return Telegram.Post("/sendDocument")
}

func (Handler *Handler) SendVideo(chat_id, video string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("video", video)
	return Telegram.Post("/sendVideo")
}

func (Handler *Handler) Animation(chat_id, animation string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("animation", animation)
	return Telegram.Post("/animation")
}

func (Handler *Handler) SendVoice(chat_id, voice string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("voice", voice)
	return Telegram.Post("/sendVoice")
}

func (Handler *Handler) SendVideoNote(chat_id, video_note string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("video_note", video_note)
	return Telegram.Post("/sendVideoNote")
}

func (Handler *Handler) SendMediaGroup(chat_id, media string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("media", media)
	return Telegram.Post("/sendMediaGroup")
}

func (Handler *Handler) SendLocation(chat_id, latitude, longitude string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("latitude", latitude)
	Telegram.AddParam("longitude", longitude)
	return Telegram.Post("/sendLocation")
}

func (Handler *Handler) SendVenue(chat_id, latitude, longitude, title, address string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("latitude", latitude)
	Telegram.AddParam("longitude", longitude)
	Telegram.AddParam("title", title)
	Telegram.AddParam("address", address)
	return Telegram.Post("/sendVenue")
}

func (Handler *Handler) SendContact(chat_id, phone_number, first_name string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("phone_number", phone_number)
	Telegram.AddParam("first_name", first_name)
	return Telegram.Post("/sendContact")
}

func (Handler *Handler) SendPoll(chat_id, question, options string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("question", question)
	Telegram.AddParam("options", options)
	return Telegram.Post("/sendPoll")
}

func (Handler *Handler) SendDice(chat_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.Post("/sendDice")
}

func (Handler *Handler) SendChatAction(chat_id, action string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("action", action)
	return Telegram.Post("/sendChatAction")
}

/*
func (Handler *Handler) GetUserProfilePhotos(user_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("user_id", user_id)
	return Telegram.get("/getUserProfilePhotos")
}

func (Handler *Handler) GetFile(file_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("file_id", file_id)
	return Telegram.get("/getFile")
}*/

func (Handler *Handler) BanChatMember(chat_id, user_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("user_id", user_id)
	return Telegram.Post("/banChatMember")
}

func (Handler *Handler) UnbanChatMember(chat_id, user_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("user_id", user_id)
	return Telegram.Post("/unbanChatMember")
}

func (Handler *Handler) RestrictChatMember(chat_id, user_id, permissions string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("user_id", user_id)
	Telegram.AddParam("permissions", permissions)
	return Telegram.Post("/restrictChatMember")
}

func (Handler *Handler) PromoteChatMember(chat_id, user_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("user_id", user_id)
	return Telegram.Post("/promoteChatMember")
}

func (Handler *Handler) SetChatAdministratorCustomTitle(chat_id, user_id, custom_title string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("user_id", user_id)
	Telegram.AddParam("custom_title", custom_title)
	return Telegram.Post("/setChatAdministratorCustomTitle")
}

func (Handler *Handler) BanChatSenderChat(chat_id, sender_chat_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("sender_chat_id", sender_chat_id)
	return Telegram.Post("/banChatSenderChat")
}

func (Handler *Handler) UnbanChatSenderChat(chat_id, sender_chat_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("sender_chat_id", sender_chat_id)
	return Telegram.Post("/unbanChatSenderChat")
}

/*
func (Handler *Handler) SetChatPermissions(chat_id, permissions string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("permissions", permissions)
	return Telegram.Post("/setChatPermissions")
}
*/
func (Handler *Handler) ExportChatInviteLink(chat_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.Post("/exportChatInviteLink")
}

func (Handler *Handler) CreateChatInviteLink(chat_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.Post("/createChatInviteLink")
}

func (Handler *Handler) EditChatInviteLink(chat_id, invite_link string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("invite_link", invite_link)
	return Telegram.Post("/editChatInviteLink")
}

func (Handler *Handler) RevokeChatInviteLink(chat_id, invite_link string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("invite_link", invite_link)
	return Telegram.Post("/revokeChatInviteLink")
}

func (Handler *Handler) ApproveChatJoinRequest(chat_id, user_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("user_id", user_id)
	return Telegram.Post("/approveChatJoinRequest")
}

func (Handler *Handler) DeclineChatJoinRequest(chat_id, user_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("user_id", user_id)
	return Telegram.Post("/declineChatJoinRequest")
}

func (Handler *Handler) SetChatPhoto(chat_id, photo string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("photo", photo)
	return Telegram.Post("/setChatPhoto")
}

func (Handler *Handler) DeleteChatPhoto(chat_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.Post("/deleteChatPhoto")
}

func (Handler *Handler) SetChatTitle(chat_id, title string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("title", title)
	return Telegram.Post("/setChatTitle")
}

func (Handler *Handler) SetChatDescription(chat_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.Post("/setChatDescription")
}

func (Handler *Handler) PinChatMessage(chat_id, message_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("message_id", message_id)
	return Telegram.Post("/pinChatMessage")
}

func (Handler *Handler) UnpinChatMessage(chat_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.Post("/unpinChatMessage")
}

func (Handler *Handler) UnpinAllChatMessages(chat_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.Post("/unpinAllChatMessages")
}

func (Handler *Handler) LeaveChat(chat_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.Post("/leaveChat")
}

/*
func (Handler *Handler) GetChat(chat_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.get("/getChat")
}

func (Handler *Handler) GetChatAdministrators(chat_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.get("/getChatAdministrators")
}

func (Handler *Handler) GetChatMemberCount(chat_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.get("/getChatMemberCount")
}

func (Handler *Handler) GetChatMember(chat_id, user_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("user_id", user_id)
	return Telegram.get("/getChatMember")
}
*/
func (Handler *Handler) SetChatStickerSet(chat_id, sticker_set_name string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("sticker_set_name", sticker_set_name)
	return Telegram.Post("/setChatStickerSet")
}

func (Handler *Handler) DeleteChatStickerSet(chat_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.Post("/deleteChatStickerSet")
}

// func (Handler *Handler) GetForumTopicIconStickers() map[string]interface{} {
// 	Telegram := methods.New(Handler.token)
// 	return Telegram.get("/deleteChatStickerSet")
// }

func (Handler *Handler) CreateForumTopic(chat_id, name string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("name", name)
	return Telegram.Post("/createForumTopic")
}

func (Handler *Handler) EditForumTopic(chat_id, message_thread_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("message_thread_id", message_thread_id)
	return Telegram.Post("/editForumTopic")
}

func (Handler *Handler) CloseForumTopic(chat_id, message_thread_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("message_thread_id", message_thread_id)
	return Telegram.Post("/closeForumTopic")
}

func (Handler *Handler) ReopenForumTopic(chat_id, message_thread_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("message_thread_id", message_thread_id)
	return Telegram.Post("/reopenForumTopic")
}

func (Handler *Handler) DeleteForumTopic(chat_id, message_thread_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("message_thread_id", message_thread_id)
	return Telegram.Post("/deleteForumTopic")
}

func (Handler *Handler) UnpinAllForumTopicMessages(chat_id, message_thread_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("message_thread_id", message_thread_id)
	return Telegram.Post("/unpinAllForumTopicMessages")
}

func (Handler *Handler) EditGeneralForumTopic(chat_id, name string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("name", name)
	return Telegram.Post("/editGeneralForumTopic")
}

func (Handler *Handler) CloseGeneralForumTopic(chat_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.Post("/closeGeneralForumTopic")
}

func (Handler *Handler) ReopenGeneralForumTopic(chat_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.Post("/reopenGeneralForumTopic")
}

func (Handler *Handler) HideGeneralForumTopic(chat_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.Post("/hideGeneralForumTopic")
}

func (Handler *Handler) UnhideGeneralForumTopic(chat_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.Post("/unhideGeneralForumTopic")
}

func (Handler *Handler) AnswerCallbackQuery(callback_query_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("callback_query_id", callback_query_id)
	return Telegram.Post("/answerCallbackQuery")
}

func (Handler *Handler) SetMyCommands(commands string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("commands", commands)
	return Telegram.Post("/setMyCommands")
}

func (Handler *Handler) DeleteMyCommands() map[string]interface{} {
	Telegram := methods.New(Handler.token)
	return Telegram.Post("/deleteMyCommands")
}

// func (Handler *Handler) GetMyCommands() map[string]interface{} {
// 	Telegram := methods.New(Handler.token)
// 	return Telegram.get("/getMyCommands")
// }

func (Handler *Handler) SetMyDescription() map[string]interface{} {
	Telegram := methods.New(Handler.token)
	return Telegram.Post("/setMyDescription")
}

/*
func (Handler *Handler) GetMyDescription() map[string]interface{} {
	Telegram := methods.New(Handler.token)
	return Telegram.get("/getMyDescription")
}

func (Handler *Handler) SetMyShortDescription() map[string]interface{} {
	Telegram := methods.New(Handler.token)
	return Telegram.Post("/setMyShortDescription")
}

func (Handler *Handler) GetMyShortDescription() map[string]interface{} {
	Telegram := methods.New(Handler.token)
	return Telegram.get("/getMyShortDescription")
}*/

func (Handler *Handler) SetChatMenuButton() map[string]interface{} {
	Telegram := methods.New(Handler.token)
	return Telegram.Post("/setChatMenuButton")
}

/*
func (Handler *Handler) GetChatMenuButton() map[string]interface{} {
	Telegram := methods.New(Handler.token)
	return Telegram.get("/getChatMenuButton")
}

func (Handler *Handler) SetMyDefaultAdministratorRights() map[string]interface{} {
	Telegram := methods.New(Handler.token)
	return Telegram.Post("/setMyDefaultAdministratorRights")
}

func (Handler *Handler) GetMyDefaultAdministratorRights() map[string]interface{} {
	Telegram := methods.New(Handler.token)
	return Telegram.get("/getMyDefaultAdministratorRights")
}*/

func (Handler *Handler) EditMessageText() map[string]interface{} {
	Telegram := methods.New(Handler.token)
	return Telegram.Post("/editMessageText")
}

func (Handler *Handler) EditMessageCaption() map[string]interface{} {
	Telegram := methods.New(Handler.token)
	return Telegram.Post("/editMessageCaption")
}

func (Handler *Handler) EditMessageMedia(media string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("media", media)
	return Telegram.Post("/editMessageMedia")
}

func (Handler *Handler) EditMessageLiveLocation(latitude, longitude string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("latitude", latitude)
	Telegram.AddParam("longitude", longitude)
	return Telegram.Post("/editMessageLiveLocation")
}

func (Handler *Handler) StopMessageLiveLocation() map[string]interface{} {
	Telegram := methods.New(Handler.token)
	return Telegram.Post("/stopMessageLiveLocation")
}

func (Handler *Handler) EditMessageReplyMarkup() map[string]interface{} {
	Telegram := methods.New(Handler.token)
	return Telegram.Post("/editMessageReplyMarkup")
}

func (Handler *Handler) StopPoll(chat_id, message_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("message_id", message_id)
	return Telegram.Post("/stopPoll")
}

func (Handler *Handler) DeleteMessage(chat_id, message_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("message_id", message_id)
	return Telegram.Post("/deleteMessage")
}

func (Handler *Handler) SendSticker(chat_id, sticker string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("sticker", sticker)
	return Telegram.Post("/sendSticker")
}

func (Handler *Handler) GetStickerSet(name string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("name", name)
	return Telegram.Post("/getStickerSet")
}

func (Handler *Handler) GetCustomEmojiStickers(custom_emoji_ids string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("custom_emoji_ids", custom_emoji_ids)
	return Telegram.Post("/getCustomEmojiStickers")
}

func (Handler *Handler) UploadStickerFile(user_id, sticker, sticker_format string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("user_id", user_id)
	Telegram.AddParam("sticker", sticker)
	Telegram.AddParam("sticker_format", sticker_format)
	return Telegram.Post("/uploadStickerFile")
}

func (Handler *Handler) CreateNewStickerSet(user_id, name, title, stickers, sticker_format string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("user_id", user_id)
	Telegram.AddParam("name", name)
	Telegram.AddParam("title", title)
	Telegram.AddParam("stickers", stickers)
	Telegram.AddParam("sticker_format", sticker_format)
	return Telegram.Post("/createNewStickerSet")
}

func (Handler *Handler) AddStickerToSet(user_id, name, sticker string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("user_id", user_id)
	Telegram.AddParam("name", name)
	Telegram.AddParam("sticker", sticker)
	return Telegram.Post("/addStickerToSet")
}

func (Handler *Handler) SetStickerPositionInSet(sticker, position string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("sticker", sticker)
	Telegram.AddParam("position", position)
	return Telegram.Post("/setStickerPositionInSet")
}

func (Handler *Handler) DeleteStickerFromSet(sticker string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("sticker", sticker)
	return Telegram.Post("/deleteStickerFromSet")
}

func (Handler *Handler) SetStickerEmojiList(sticker, emoji_list string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("sticker", sticker)
	Telegram.AddParam("emoji_list", emoji_list)
	return Telegram.Post("/setStickerEmojiList")
}

func (Handler *Handler) SetStickerKeywords(sticker string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("sticker", sticker)
	return Telegram.Post("/setStickerKeywords")
}

func (Handler *Handler) SetStickerMaskPosition(sticker string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("sticker", sticker)
	return Telegram.Post("/setStickerMaskPosition")
}

func (Handler *Handler) SetStickerSetTitle(name, title string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("name", name)
	Telegram.AddParam("title", title)
	return Telegram.Post("/setStickerSetTitle")
}

func (Handler *Handler) SetStickerSetThumbnail(name, user_id string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("name", name)
	Telegram.AddParam("user_id", user_id)
	return Telegram.Post("/setStickerSetThumbnail")
}

func (Handler *Handler) SetCustomEmojiStickerSetThumbnail(name string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("name", name)
	return Telegram.Post("/setCustomEmojiStickerSetThumbnail")
}

func (Handler *Handler) DeleteStickerSet(name string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("name", name)
	return Telegram.Post("/deleteStickerSet")
}

func (Handler *Handler) AnswerInlineQuery(inline_query_id, results string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("inline_query_id", inline_query_id)
	Telegram.AddParam("results", results)
	return Telegram.Post("/answerInlineQuery")
}

func (Handler *Handler) AnswerWebAppQuery(web_app_query_id, results string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("web_app_query_id", web_app_query_id)
	Telegram.AddParam("results", results)
	return Telegram.Post("/answerWebAppQuery")
}

func (Handler *Handler) SendInvoice(chat_id, title, description, payload, provider_token, currency, prices string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("title", title)
	Telegram.AddParam("description", description)
	Telegram.AddParam("payload", payload)
	Telegram.AddParam("provider_token", provider_token)
	Telegram.AddParam("currency", currency)
	Telegram.AddParam("prices", prices)
	return Telegram.Post("/sendInvoice")
}

func (Handler *Handler) CreateInvoiceLink(chat_id, title, description, payload, provider_token, currency, prices string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("title", title)
	Telegram.AddParam("description", description)
	Telegram.AddParam("payload", payload)
	Telegram.AddParam("provider_token", provider_token)
	Telegram.AddParam("currency", currency)
	Telegram.AddParam("prices", prices)
	return Telegram.Post("/createInvoiceLink")
}

func (Handler *Handler) AnswerShippingQuery(shipping_query_id, ok string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("shipping_query_id", shipping_query_id)
	Telegram.AddParam("ok", ok)
	return Telegram.Post("/answerShippingQuery")
}

func (Handler *Handler) AnswerPreCheckoutQuery(pre_checkout_query_id, ok string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("pre_checkout_query_id", pre_checkout_query_id)
	Telegram.AddParam("ok", ok)
	return Telegram.Post("/answerPreCheckoutQuery")
}

func (Handler *Handler) SetPassportDataErrors(user_id, errors string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("user_id", user_id)
	Telegram.AddParam("errors", errors)
	return Telegram.Post("/setPassportDataErrors")
}

func (Handler *Handler) SendGame(chat_id, game_short_name string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("game_short_name", game_short_name)
	return Telegram.Post("/sendGame")
}

func (Handler *Handler) SetGameScore(user_id, score string) map[string]interface{} {
	Telegram := methods.New(Handler.token)
	Telegram.AddParam("user_id", user_id)
	Telegram.AddParam("score", score)
	return Telegram.Post("/setGameScore")
}

// func (Handler *Handler) GetGameHighScores(user_id string) map[string]interface{} {
// 	Telegram := methods.New(Handler.token)
// 	Telegram.AddParam("user_id", user_id)
// 	return Telegram.get("/getGameHighScores")
// }
