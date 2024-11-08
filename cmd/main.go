package main

import (
	"log"
	"os"

	tg "github.com/OvyFlash/telegram-bot-api"
	"github.com/joho/godotenv"
  "cf-bot/internal/handlers"
)

func main() {
  err := godotenv.Load(".env")
  if err != nil {
    panic("can't load token")
  }

  token := os.Getenv("TOKEN")
  adminsChatID := int64(-1002269839756)

  bot, err := tg.NewBotAPI(token)

  log.Printf("авторизован под %s", bot.Self.UserName)

  updateConfig := tg.NewUpdate(0)
  updateConfig.Timeout = 60

  updates := bot.GetUpdatesChan(updateConfig)

  for u := range updates {
    //если нет сообщения пропускаем итерацию
    if u.Message == nil {
      continue
    }
    
    //переменные
    chatID := u.Message.Chat.ID
    msgText := u.Message.Text
    userName := u.SentFrom().UserName
    
    //обрабатываем команды
    switch msgText {
    case "/start":
      logTxt(msgText, userName)
      handlers.Start(chatID, bot)
    case "анон":
      logTxt(msgText, userName)
      handlers.TakeTxt(chatID, bot)

      //цикл для ожидания тейка
      for u := range updates {
        //если сообщения нет пропускаем итерацию
        if u.Message == nil {
          continue
        }

        //переменные
        chatID := u.Message.Chat.ID
        msgText := u.Message.Text
        userName := u.SentFrom().UserName    
        
        //если нажата кнопка отмены тейка
        if msgText == "не хочу отправлять тейк" {
          handlers.WontWriteTake(chatID, bot)
      
          //возврат в стартовое меню
          handlers.Start(chatID, bot)
          break
        }
        
        handlers.AnonTxt(chatID, bot, adminsChatID, msgText, userName)
        //возврат в стартовое меню
        handlers.Start(chatID, bot)
        break
      }
    case "неанон":
      logTxt(msgText, userName)
      handlers.TakeTxt(chatID, bot)

      //цикл для ожидания тейка
      for u := range updates {
        //если сообщения нет пропускаем итерацию
        if u.Message == nil {
          continue
        }

        //переменные
        chatID := u.Message.Chat.ID
        msgText := u.Message.Text
        userName := u.SentFrom().UserName
        
        //если нажата кнопка отмены тейка
        if msgText == "не хочу отправлять тейк" {
          handlers.WontWriteTake(chatID, bot)
      
          //возврат в стартовое меню
          handlers.Start(chatID, bot)
          break
        }
        
        handlers.NeanonTxt(chatID, bot, adminsChatID, msgText, userName) 
        //возврат в стартовое меню
        handlers.Start(chatID, bot)
        break
      }
    default:
      handlers.Default(chatID, bot) 
    }
  }
}
  
func logTxt(text string, username string) {
  log.Printf("сообщение от @%s: %s", username, text)
}
