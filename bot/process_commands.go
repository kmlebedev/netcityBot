package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
	"time"
)

func ProcessCommand(updateMsg *tgbotapi.Message, sendMsg *tgbotapi.MessageConfig, bot *tgbotapi.BotAPI) {
	netCityApi := GetLoginWebApi(sendMsg.ChatID)
	if netCityApi == nil && updateMsg.Command() != "start" {
		sendMsg.Text = "Вы не вошли в дневник"
		return
	}
	user := GetChatUser(sendMsg.ChatID)
	switch updateMsg.Command() {
	case "contacts":
		if mobilePhone, email, err := netCityApi.GetContacts(); err == nil {
			sendMsg.Text = fmt.Sprintf("Мобильный телефон: %s и email: %s", mobilePhone, email)
		} else {
			log.Errorf("netCityApi.GetContacts: %v", err)
		}

	case "start":
		ReplySelectCity(sendMsg)
	case "track_marks":
		if user.Marks != nil {
			sendMsg.Text = ""
			if user.TrackMarksCn != nil {
				user.TrackMarksCn <- true
			}
			return
		}
		var err error
		user.Marks, err = user.NetCityApi.GetLessonAssignmentMarks()
		if err != nil {
			sendMsg.Text = fmt.Sprintf("Что то пошло не так: %+v", err)
			return
		}
		sendMsg.Text = fmt.Sprintf("Включено отслеживание отметок")
		user.TrackMarksCn = make(chan bool)
		go func(chatID int64, bot *tgbotapi.BotAPI, login *User) {
			tick := time.Tick(time.Duration(5) * time.Minute)
			for {
				select {
				case <-login.TrackMarksCn:
					login.Marks = nil
					if _, err := bot.Send(tgbotapi.NewMessage(chatID,
						fmt.Sprintf("Отключено отслеживание отметок"))); err != nil {
						log.Warningf("bot.Send: %+v", err)
					}
					return
				case <-tick:
					if msg, err := trackMarks(login); err == nil && msg != "" {
						if _, err = bot.Send(tgbotapi.NewMessage(chatID, msg)); err != nil {
							log.Warningf("bot.Send: %+v", err)
						}
					}
				}
			}
		}(sendMsg.ChatID, bot, user)

	case "hello":
		sendMsg.Text = fmt.Sprintf("И тебе привет %s", user.LoginName)
	case "login":
		sendMsg.Text = "login"
		ReplySelectCity(sendMsg)
	case "logout":
		sendMsg.Text = "logout"
		netCityApi.Logout()
	}
}