# MajsoulGo

MajsoulGo is a module for connecting to and communicating with Mahjong Soul's game servers and contest management servers.

## Getting Started

MajsoulGo provides two different clients for connecting to Majsoul's servers. The package `dhs` contains the client and protobuf message types for the contest management server. The package `lq` contains the client and protobuf message types for the game server.

The following examples will use the contest management client (`dhs.ContestManagerClient`), but the process is the same for the game client (`lq.MajsoulGameClient`).

### Connecting to the server

Import the the `majsoulgo` and `dhs` packages:
```go
import (
    "github.com/oscarfzs/majsoulgo"
    "github.com/oscarfzs/dhs"
)
```

Create the `dhs.ContestManagerClient` object, then retrieve the latest contest management server url:
```go
client := dhs.NewContestManagerClient()

url, err := majsoulgo.GetContestManagementServerUrl()
if err != nil {
    log.Fatal(err)
}

err = client.Connect(url)
```
`client.Connect()` will spawn an event loop that listens for messages from the server. It will also periodically ping the server to keep the connection alive. If the connection closes unexpectedly, or an interrupt signal (ctrl-c) is sent, then ```client.Connect()``` will return with a non-nil error.

Alternatively, you may run `client.Connect` as a goroutine, then retrieve the return value through `client.ExitValue()`:
```go
go client.Connect(url)

//Do some stuff

err = client.ExitValue()
```

### Sending and receiving messages

```go
req := &dhs.ReqContestManageOauth2Login{
    Type: 10,
    AccessToken: "YOUR-TOKEN-HERE"
}

res, err := client.Oauth2LoginContestManager(req)
if err != nil {
    log.Fatal(err)
}

log.Println(res)
```

### Handling notifications
```go
//Run this function whenever we receive a dhs.NotifyContestGameEnd message from the server
func foo(pbMsg proto.Message) {
    msg := pbMsg.(*dhs.NotifyContestGameEnd)  //convert the protobuf message in order to access its fields.
    log.Println(msg.game_uuid)
}
```

```go
client.AddNotificationHandler(&dhs.NotifyContestGameEnd{}, foo)
```
