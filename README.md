### Simple bot for twitch, who can speech message from chat

### Requirements

install
[mplayer](http://www.mplayerhq.hu/design7/dload.html) and add mplayer binary to PATH variable

### setup .env file

create application at [link](https://dev.twitch.tv/console) with redirect uri http://localhost

open on your favorite browser this link
```https://id.twitch.tv/oauth2/authorize?response_type=token &client_id={clientId}&redirect_uri=http://localhost&scope=channel%3Amanage%3Apolls+channel%3Aread%3Apolls &state=c3ab8aa609ea11e793ae92361f002671```

### Build

##### Building for windows set this environment variable

GOARCH=amd64 GOOS=windows
```make build && main.exe```

##### Building for linux set this environment variable

GOOS=linux GOARCH=amd64
```go build -o main main.go```
```chmod +x ./main```

##### Building for macos set this environment variable

GOOS=darwin GOARCH=amd64
```go build -o main main.go```
```chmod +x ./main```





