# rich-go fork

An implementation of Discord's rich presence in Golang for Linux, macOS and Windows

This is a forked version of [rich-go](https://github.com/hugolgst/rich-go). During development, I was faced with the inability
to change activity in Discord due to the error "The pipe is being closed", since there was no error return in the rich-go code
when trying to access Discord RPC, it was impossible to reconnect. For this reason , I made a fork in which
when trying to access Discord RPC, an error is returned and when the error "The pipe is being closed" occurs,
an attempt is made to automatically reconnect to Discord RPC.

Это форк [rich-go](https://github.com/hugolgst/rich-go). Во время разработки я столкнулся с невозможностью
изменения активности в Discord из-за ошибки "The pipe is being closed", так как в коде rich-go не было возврата
ошибок при попытке обращения к Discord RPC было невозможно сделать переподключение. По этой причине
я сделал форк в котором при попытке обращения к Discord RPC, возвращается ошибка и при 
возникновении ошибки "The pipe is being closed" происходит попытка автоматического переподключения к Discord RPC.

## Install:
**Put this text in you go.mod**
```
   replace github.com/hugolgst/rich-go master => github.com/ledxdeliveryflopp/rich-go master
```

```bash 
    go get github.com/hugolgst/rich-go
 ```

## Установка:
**Поместите этот текст в go.mod вашего проекта**
```
replace github.com/hugolgst/rich-go master => github.com/ledxdeliveryflopp/rich-go master
```
**Установите оригинальный rich-go**
```bash 
  go get github.com/hugolgst/rich-go
 ```

## Usage

First of all import rich-go
```golang
import "github.com/hugolgst/rich-go/client"
```

then login by sending the first handshake
```golang
err := client.Login("DISCORD_APP_ID")
if err != nil {
	panic(err)
}
```

and you can set the Rich Presence activity (parameters can be found :
```golang
err = client.SetActivity("DISCORD_APP_ID", client.Activity{
	State:      "Heyy!!!",
	Details:    "I'm running on rich-go :)",
	LargeImage: "largeimageid",
	LargeText:  "This is the large image :D",
	SmallImage: "smallimageid",
	SmallText:  "And this is the small image",
	Party: &client.Party{
		ID:         "-1",
		Players:    15,
		MaxPlayers: 24,
	},
	Timestamps: &client.Timestamps{
		Start: time.Now(),
	},
})

if err != nil {
	panic(err)
}
```

More details in the [example](https://github.com/ananagame/rich-go/blob/master/example/main.go)

## Contributing

1. Fork it (https://github.com/ledxdeliveryflopp/rich-go/fork)
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create a new Pull Request

## Contributors

- [ledxdeliveryflopp](https://github.com/ledxdeliveryflopp) - forker, fork maintainer
- [hugolgst](https://github.com/hugolgst) - original creator, maintainer of original version
- [donovansolms](https://github.com/donovansolms) - contributor
- [heroslender](https://github.com/heroslender) - contributor
