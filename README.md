## Telegram poll bot
To interact with bot you create a chat with this bot and then you could paste him a command: `ADD <URL>`.

### Install local development environment
We will be using Docker to simplify installation and to be OS agnostic:
```
$ docker build . -t godev
$ docker run --rm -ti -v ${PWD}:/work godev sh 
/work # go version
go version go1.20.3 linux/arm64
```

### Running a bot
First we need to export API_TOKEN, then initiate golang package and run or build:
```
// example tocken first part is `chatID : token`
/work # export TELEGRAM_APITOKEN="1118344111:AAFuA7Ggp-7Rww74JIOAMCTkleoMCgWzgZw"
/work # go mod init kita-koguta-bot
/work # go run main.go
go: downloading github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
2023/07/12 19:36:16 Endpoint: getUpdates, params: map[allowed_updates:null timeout:30]
```
