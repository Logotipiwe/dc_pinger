package model

type PingConfig struct {
	NotifyChatID int64        `json:"notifyChatId"`
	Targets      []PingTarget `json:"targets"`
}

type PingTarget struct {
	Name         string             `json:"name"`
	Message      string             `json:"message"`
	NotifyChatID int64              `json:"notifyChatId"`
	Requests     PingTargetRequests `json:"requests"`
}

type PingTargetRequests struct {
	Url             string `json:"url"`
	IntervalSec     int    `json:"intervalSec"`
	FailIntervalSec int    `json:"failIntervalSec"`
	TimeoutMs       int    `json:"timeoutMs"`
}
