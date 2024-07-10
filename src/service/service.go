package service

import . "github.com/logotipiwe/dc_go_config_lib"

type Services struct {
	PingerService        *PingerService
	NotificationsService *NotificationsService
	Interruptor          *Interruptor
}

func CreateServices() *Services {
	i := NewInterruptor()
	ns := NewNotificationsService(GetConfig("BOT_TOKEN"), i)
	return &Services{
		PingerService:        NewPingerService(ns, i),
		NotificationsService: ns,
		Interruptor:          i,
	}
}
