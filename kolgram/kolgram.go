package kolgram

import (
	request "github.com/koltypka/kolRequest/kolRequest"
)

type Telegram struct {
	Request request.Request
}

func New(token string) Telegram {
	return Telegram{request.New("https://api.telegram.org/bot" + token)}
}

func (Telegram *Telegram) AddParam(key, value string) {
	Telegram.Request.AddParam(key, value)
}

func (Telegram *Telegram) GetUpdates() map[string]interface{} {
	return Telegram.get("/getUpdates")
}

func (Telegram *Telegram) SetWebhook(url string) map[string]interface{} {
	Telegram.AddParam("url", url)
	return Telegram.post("/setWebhook")
}

func (Telegram *Telegram) DeleteWebhook() map[string]interface{} {
	return Telegram.post("/deleteWebhook")
}

func (Telegram *Telegram) GetWebhookInfo() map[string]interface{} {
	return Telegram.get("/getWebhookInfo")
}

func (Telegram *Telegram) GetMe() map[string]interface{} {
	return Telegram.get("/getMe")
}

func (Telegram *Telegram) LogOut() map[string]interface{} {
	return Telegram.post("/logOut")
}

func (Telegram *Telegram) Close() map[string]interface{} {
	return Telegram.post("/close")
}

func (Telegram *Telegram) SendMessage(chat_id, text string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("text", text)
	return Telegram.post("/sendMessage")
}

func (Telegram *Telegram) ForwardMessage(chat_id, from_chat_id, message_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("from_chat_id", from_chat_id)
	Telegram.AddParam("message_id", message_id)
	return Telegram.post("/forwardMessage")
}

func (Telegram *Telegram) CopyMessage(chat_id, from_chat_id, message_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("from_chat_id", from_chat_id)
	Telegram.AddParam("message_id", message_id)
	return Telegram.post("/copyMessage")
}

func (Telegram *Telegram) SendPhoto(chat_id, photo string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("photo", photo)
	return Telegram.post("/sendPhoto")
}

func (Telegram *Telegram) SendAudio(chat_id, audio string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("audio", audio)
	return Telegram.post("/sendAudio")
}

func (Telegram *Telegram) SendDocument(chat_id, document string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("document", document)
	return Telegram.post("/sendDocument")
}

func (Telegram *Telegram) SendVideo(chat_id, video string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("video", video)
	return Telegram.post("/sendVideo")
}

func (Telegram *Telegram) Animation(chat_id, animation string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("animation", animation)
	return Telegram.post("/animation")
}

func (Telegram *Telegram) SendVoice(chat_id, voice string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("voice", voice)
	return Telegram.post("/sendVoice")
}

func (Telegram *Telegram) SendVideoNote(chat_id, video_note string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("video_note", video_note)
	return Telegram.post("/sendVideoNote")
}

func (Telegram *Telegram) SendMediaGroup(chat_id, media string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("media", media)
	return Telegram.post("/sendMediaGroup")
}

func (Telegram *Telegram) SendLocation(chat_id, latitude, longitude string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("latitude", latitude)
	Telegram.AddParam("longitude", longitude)
	return Telegram.post("/sendLocation")
}

func (Telegram *Telegram) SendVenue(chat_id, latitude, longitude, title, address string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("latitude", latitude)
	Telegram.AddParam("longitude", longitude)
	Telegram.AddParam("title", title)
	Telegram.AddParam("address", address)
	return Telegram.post("/sendVenue")
}

func (Telegram *Telegram) SendContact(chat_id, phone_number, first_name string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("phone_number", phone_number)
	Telegram.AddParam("first_name", first_name)
	return Telegram.post("/sendContact")
}

func (Telegram *Telegram) SendPoll(chat_id, question, options string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("question", question)
	Telegram.AddParam("options", options)
	return Telegram.post("/sendPoll")
}

func (Telegram *Telegram) SendDice(chat_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.post("/sendDice")
}

func (Telegram *Telegram) SendChatAction(chat_id, action string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("action", action)
	return Telegram.post("/sendChatAction")
}

func (Telegram *Telegram) GetUserProfilePhotos(user_id string) map[string]interface{} {
	Telegram.AddParam("user_id", user_id)
	return Telegram.get("/getUserProfilePhotos")
}

func (Telegram *Telegram) GetFile(file_id string) map[string]interface{} {
	Telegram.AddParam("file_id", file_id)
	return Telegram.get("/getFile")
}

func (Telegram *Telegram) BanChatMember(chat_id, user_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("user_id", user_id)
	return Telegram.post("/banChatMember")
}

func (Telegram *Telegram) UnbanChatMember(chat_id, user_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("user_id", user_id)
	return Telegram.post("/unbanChatMember")
}

func (Telegram *Telegram) RestrictChatMember(chat_id, user_id, permissions string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("user_id", user_id)
	Telegram.AddParam("permissions", permissions)
	return Telegram.post("/restrictChatMember")
}

func (Telegram *Telegram) PromoteChatMember(chat_id, user_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("user_id", user_id)
	return Telegram.post("/promoteChatMember")
}

func (Telegram *Telegram) SetChatAdministratorCustomTitle(chat_id, user_id, custom_title string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("user_id", user_id)
	Telegram.AddParam("custom_title", custom_title)
	return Telegram.post("/setChatAdministratorCustomTitle")
}

func (Telegram *Telegram) BanChatSenderChat(chat_id, sender_chat_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("sender_chat_id", sender_chat_id)
	return Telegram.post("/banChatSenderChat")
}

func (Telegram *Telegram) UnbanChatSenderChat(chat_id, sender_chat_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("sender_chat_id", sender_chat_id)
	return Telegram.post("/unbanChatSenderChat")
}

func (Telegram *Telegram) SetChatPermissions(chat_id, permissions string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("permissions", permissions)
	return Telegram.post("/setChatPermissions")
}

func (Telegram *Telegram) ExportChatInviteLink(chat_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.post("/exportChatInviteLink")
}

func (Telegram *Telegram) CreateChatInviteLink(chat_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.post("/createChatInviteLink")
}

func (Telegram *Telegram) EditChatInviteLink(chat_id, invite_link string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("invite_link", invite_link)
	return Telegram.post("/editChatInviteLink")
}

func (Telegram *Telegram) RevokeChatInviteLink(chat_id, invite_link string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("invite_link", invite_link)
	return Telegram.post("/revokeChatInviteLink")
}

func (Telegram *Telegram) ApproveChatJoinRequest(chat_id, user_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("user_id", user_id)
	return Telegram.post("/approveChatJoinRequest")
}

func (Telegram *Telegram) DeclineChatJoinRequest(chat_id, user_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("user_id", user_id)
	return Telegram.post("/declineChatJoinRequest")
}

func (Telegram *Telegram) SetChatPhoto(chat_id, photo string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("photo", photo)
	return Telegram.post("/setChatPhoto")
}

func (Telegram *Telegram) DeleteChatPhoto(chat_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.post("/deleteChatPhoto")
}

func (Telegram *Telegram) SetChatTitle(chat_id, title string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("title", title)
	return Telegram.post("/setChatTitle")
}

func (Telegram *Telegram) SetChatDescription(chat_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.post("/setChatDescription")
}

func (Telegram *Telegram) PinChatMessage(chat_id, message_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("message_id", message_id)
	return Telegram.post("/pinChatMessage")
}

func (Telegram *Telegram) UnpinChatMessage(chat_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.post("/unpinChatMessage")
}

func (Telegram *Telegram) UnpinAllChatMessages(chat_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.post("/unpinAllChatMessages")
}

func (Telegram *Telegram) LeaveChat(chat_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.post("/leaveChat")
}

func (Telegram *Telegram) GetChat(chat_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.get("/getChat")
}

func (Telegram *Telegram) GetChatAdministrators(chat_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.get("/getChatAdministrators")
}

func (Telegram *Telegram) GetChatMemberCount(chat_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.get("/getChatMemberCount")
}

func (Telegram *Telegram) GetChatMember(chat_id, user_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("user_id", user_id)
	return Telegram.get("/getChatMember")
}

func (Telegram *Telegram) SetChatStickerSet(chat_id, sticker_set_name string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("sticker_set_name", sticker_set_name)
	return Telegram.post("/setChatStickerSet")
}

func (Telegram *Telegram) DeleteChatStickerSet(chat_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.post("/deleteChatStickerSet")
}

func (Telegram *Telegram) GetForumTopicIconStickers() map[string]interface{} {
	return Telegram.get("/deleteChatStickerSet")
}

func (Telegram *Telegram) CreateForumTopic(chat_id, name string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("name", name)
	return Telegram.post("/createForumTopic")
}

func (Telegram *Telegram) EditForumTopic(chat_id, message_thread_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("message_thread_id", message_thread_id)
	return Telegram.post("/editForumTopic")
}

func (Telegram *Telegram) CloseForumTopic(chat_id, message_thread_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("message_thread_id", message_thread_id)
	return Telegram.post("/closeForumTopic")
}

func (Telegram *Telegram) ReopenForumTopic(chat_id, message_thread_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("message_thread_id", message_thread_id)
	return Telegram.post("/reopenForumTopic")
}

func (Telegram *Telegram) DeleteForumTopic(chat_id, message_thread_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("message_thread_id", message_thread_id)
	return Telegram.post("/deleteForumTopic")
}

func (Telegram *Telegram) UnpinAllForumTopicMessages(chat_id, message_thread_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("message_thread_id", message_thread_id)
	return Telegram.post("/unpinAllForumTopicMessages")
}

func (Telegram *Telegram) EditGeneralForumTopic(chat_id, name string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	Telegram.AddParam("name", name)
	return Telegram.post("/editGeneralForumTopic")
}

func (Telegram *Telegram) CloseGeneralForumTopic(chat_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.post("/closeGeneralForumTopic")
}

func (Telegram *Telegram) ReopenGeneralForumTopic(chat_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.post("/reopenGeneralForumTopic")
}

func (Telegram *Telegram) HideGeneralForumTopic(chat_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.post("/hideGeneralForumTopic")
}

func (Telegram *Telegram) UnhideGeneralForumTopic(chat_id string) map[string]interface{} {
	Telegram.AddParam("chat_id", chat_id)
	return Telegram.post("/unhideGeneralForumTopic")
}

func (Telegram *Telegram) AnswerCallbackQuery(callback_query_id string) map[string]interface{} {
	Telegram.AddParam("callback_query_id", callback_query_id)
	return Telegram.post("/answerCallbackQuery")
}

func (Telegram *Telegram) SetMyCommands(commands string) map[string]interface{} {
	Telegram.AddParam("commands", commands)
	return Telegram.post("/setMyCommands")
}

func (Telegram *Telegram) DeleteMyCommands() map[string]interface{} {
	return Telegram.post("/deleteMyCommands")
}

func (Telegram *Telegram) GetMyCommands() map[string]interface{} {
	return Telegram.get("/getMyCommands")
}

func (Telegram *Telegram) SetMyDescription() map[string]interface{} {
	return Telegram.post("/setMyDescription")
}

func (Telegram *Telegram) GetMyDescription() map[string]interface{} {
	return Telegram.get("/getMyDescription")
}

func (Telegram *Telegram) SetMyShortDescription() map[string]interface{} {
	return Telegram.post("/setMyShortDescription")
}

func (Telegram *Telegram) GetMyShortDescription() map[string]interface{} {
	return Telegram.get("/getMyShortDescription")
}

func (Telegram *Telegram) SetChatMenuButton() map[string]interface{} {
	return Telegram.post("/setChatMenuButton")
}

func (Telegram *Telegram) GetChatMenuButton() map[string]interface{} {
	return Telegram.get("/getChatMenuButton")
}

func (Telegram *Telegram) SetMyDefaultAdministratorRights() map[string]interface{} {
	return Telegram.post("/setMyDefaultAdministratorRights")
}

func (Telegram *Telegram) GetMyDefaultAdministratorRights() map[string]interface{} {
	return Telegram.get("/getMyDefaultAdministratorRights")
}

func (Telegram *Telegram) EditMessageText() map[string]interface{} {
	return Telegram.post("/editMessageText")
}

//todo editMessageCaption()

func (Telegram *Telegram) get(method string) map[string]interface{} {
	getResult, _ := Telegram.Request.Get(method)
	result, _ := getResult.ToJson()

	return result
}

func (Telegram *Telegram) post(method string) map[string]interface{} {
	getResult, _ := Telegram.Request.Post(method)
	result, _ := getResult.ToJson()

	return result
}
