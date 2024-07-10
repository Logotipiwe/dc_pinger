<h1>Pinger env spec:</h1>
To apply config changes service should be <u>restarted</u>

---
<h3>PING_CONFIG</h3>
```json
{
  "targets": [
    {
      "name": "str",
      "notifyChatId": 2435563,
      "message": "Service A is shutted down!!",
      "requests": {
        "url": "https://logotipiwe.ru",
        "intervalSec": 60,
        "failIntervalSec": 7200,
        "timeoutMs": 5000
      }
    },
    ...
  ]
}
```
---
<h3>BOT_TOKEN</h3>
string token of tg bot

---