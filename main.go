package main

import (
	"dc_pinger/src/http"
	"dc_pinger/src/service"
	"github.com/logotipiwe/dc_go_config_lib"
)

func main() {
	dc_go_config_lib.LoadDcConfigDynamically(3)
	services := service.CreateServices()
	go services.NotificationsService.HandleUpdates()
	go services.PingerService.StartPinger()
	http.StartServer()
}
