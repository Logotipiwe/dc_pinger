package service

import (
	"dc_pinger/src/model"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/logotipiwe/dc_go_config_lib"
	"time"
)

type PingerService struct {
	NotificationsService *NotificationsService
	Interruptor          *Interruptor
}

func NewPingerService(notificationsService *NotificationsService, i *Interruptor) *PingerService {
	return &PingerService{
		NotificationsService: notificationsService,
		Interruptor:          i,
	}
}

func (p *PingerService) StartPinger() {
	pingConfig := p.getPingConfig()

	for _, target := range pingConfig.Targets {
		go p.startPingTarget(target)
	}
}

func (p *PingerService) getPingConfig() *model.PingConfig {
	configStr := dc_go_config_lib.GetConfig("PING_CONFIG")
	pingConfig := &model.PingConfig{}
	err := json.Unmarshal([]byte(configStr), pingConfig)
	if err != nil {
		panic(err)
	}
	return pingConfig
}

func (p *PingerService) startPingTarget(target model.PingTarget) {
	for {
		println("Ping " + target.Name + " with url " + target.Requests.Url)
		var sleepForSec int
		if !p.Interruptor.isInterrupted {
			err := p.doPing(target.Requests.Url, target.Requests.TimeoutMs)
			if err != nil {
				println("Error pinging " + target.Name + ": " + err.Error())
				err = p.NotificationsService.SendMessage(target.NotifyChatID, target.Message+"\r\n\r\n"+err.Error())
				if err != nil {
					println("Error sending error notification: " + err.Error())
				}
				sleepForSec = target.Requests.FailIntervalSec
			} else {
				sleepForSec = target.Requests.IntervalSec
			}
		} else {
			println("Skip because of interruption...")
			sleepForSec = target.Requests.IntervalSec
		}
		time.Sleep(time.Duration(sleepForSec) * time.Second)
	}
}

func (p *PingerService) doPing(url string, timeoutMs int) error {
	get, err := resty.New().
		SetTimeout(time.Duration(timeoutMs) * time.Millisecond).
		R().
		Get(url)
	if err != nil {
		return err
	}
	if get.StatusCode() != 200 {
		return errors.New("response status code is " + get.Status() + " instead of 200")
	}
	return nil
}
