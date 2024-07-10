package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	. "github.com/logotipiwe/dc_go_config_lib"
	"strconv"
	"strings"
)

type NotificationsService struct {
	bot         *tgbotapi.BotAPI
	Interruptor *Interruptor
}

func NewNotificationsService(token string, i *Interruptor) *NotificationsService {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	service := &NotificationsService{
		bot:         bot,
		Interruptor: i,
	}
	return service
}

func (n *NotificationsService) SendMessage(chatID int64, message string) error {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := n.bot.Send(msg)
	return err
}

func (n *NotificationsService) HandleUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := n.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				if strings.HasPrefix(update.Message.Text, "/stop") {
					n.Interruptor.Interrupt()
					n.SendMessage(update.Message.Chat.ID, "Бот остановлен")
				} else if strings.HasPrefix(update.Message.Text, "/start") {
					n.Interruptor.Resume()
					n.SendMessage(update.Message.Chat.ID, "Бот запущен")
				} else if strings.HasPrefix(update.Message.Text, "/help") {
					n.SendMessage(update.Message.Chat.ID, "Айди чата: "+strconv.FormatInt(update.Message.Chat.ID, 10))
				} else {
					if update.SentFrom().ID != getOwnerID() {
						n.SendMessage(update.Message.Chat.ID, "Командовать мной может только Герман")
					} else {
						n.SendMessage(update.Message.Chat.ID, "Не понял")
					}
				}
			} else {
				if strings.Contains(update.Message.Text, "@LogoPingerBot") {
					n.SendMessage(update.Message.Chat.ID, "@"+update.Message.From.UserName+" дурак")
				}
			}
		}
	}
}

func getOwnerID() int64 {
	id, err := GetConfigInt("TG_OWNER_ID")
	if err != nil {
		panic(err)
	}
	return int64(id)
}
