package handlers

import (
	"cf-bot/internal/keyboards"
	"cf-bot/internal/texts"
	"log"

	tg "github.com/OvyFlash/telegram-bot-api"
)

func Start(chatID int64, bot *tg.BotAPI) {
  msg := tg.NewMessage(chatID, texts.Start)
  msg.ReplyMarkup = keyboards.StartKB

  bot.Send(msg) 
}

func TakeTxt(chatID int64, bot *tg.BotAPI) {
  msg := tg.NewMessage(chatID, texts.Take)
  msg.ReplyMarkup = tg.NewRemoveKeyboard(true)
  msg.ReplyMarkup = keyboards.TakeKB

  bot.Send(msg)
}

func AnonTxt(chatID int64, bot *tg.BotAPI, adminsChatID int64, usrMsg string, userName string) {
  msgToAdmins := tg.NewMessage(adminsChatID, usrMsg + "\n\n#тейк")
  bot.Send(msgToAdmins)
  msg := tg.NewMessage(chatID, "тейк был отправлен админам")
  msg.ReplyMarkup = tg.NewRemoveKeyboard(true)
  bot.Send(msg)
  log.Printf("анонимный тейк от @%s был отправлен в чат админов", userName)
}

func NeanonTxt(chatID int64, bot *tg.BotAPI, adminsChatID int64, usrMsg string, userName string) {
  msgToAdmins := tg.NewMessage(adminsChatID, usrMsg + "\n\n#тейк")
  msgToAdmins2 := tg.NewMessage(adminsChatID, "тейк от @" + userName)
  bot.Send(msgToAdmins)
  bot.Send(msgToAdmins2)
  msg := tg.NewMessage(chatID, "тейк был отправлен админам")
  msg.ReplyMarkup = tg.NewRemoveKeyboard(true)
  bot.Send(msg)
  log.Printf("неанонимный тейк от @%s был отправлен в чат админов", userName)
}

func WontWriteTake(chatID int64, bot *tg.BotAPI) {
  msg := tg.NewMessage(chatID, "")
  msg.ReplyMarkup = tg.NewRemoveKeyboard(true) 
}

func Default(chatID int64, bot *tg.BotAPI) {
  msg := tg.NewMessage(chatID, "я не знаю такой команды")
  bot.Send(msg)
}
