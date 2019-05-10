package main

import (
	"os"
	"strings"

	"github.com/alecthomas/kingpin"
)

/*
	架构：
	https://www.ctolib.com/kingpin.html

	usage: curl [<flags>] <command> [<args> ...]

	写在 command 前面的 flags 是全局参数，[] 表示可选
	command 表示具体子服务命令
	command 后面可以用 flag 也可以用 args 表示，args 表示位置参数，require 的要放在前面，通过 * 取值

	命令：
	new：application name
	command：子命令，服务通过 flag 或 Args 绑定不同的命令参数，返回值是地址，所以需要 * 取值
	Flag：前面必须加 --
	Args：是位置参数，require 放在最前面，不用 --，Flag 必须使用 --
		PlaceHolder
		require
		default
	version：打印后直接退出程序
	short：简写
*/

/*
	例一：

	usage: kingpinTest [<flags>] <ip> [<count>]

	Flags:
		--help        Show context-sensitive help (also try --help-long and --help-man).
		--debug       Enable debug mode.
	-t, --timeout=5s  Timeout waiting for ping.
		--version     Show application version.

	Args:
	<ip>       IP address to ping.
	[<count>]  Number of packets to send

	命令行形式：
	go run kingpinTest.go --debug -t 5s 127.0.0.1 3
	go run kingpinTest.go --debug --timeout=5s 127.0.0.1 3

	文件形式：
	go run kingpinTest.go @filename
	内容:
	-t5s
	127.0.0.1
*/
// var (
// 	debug   = kingpin.Flag("debug", "Enable debug mode.").Bool()
// 	timeout = kingpin.Flag("timeout", "Timeout waiting for ping.").Default("5s").OverrideDefaultFromEnvar("PING_TIMEOUT").Short('t').Duration()
// 	ip      = kingpin.Arg("ip", "IP address to ping.").Required().IP()
// 	count   = kingpin.Arg("count", "Number of packets to send").Int()
// )

// func main() {
// 	kingpin.Version("0.0.1")
// 	kingpin.Parse()
// 	fmt.Printf("Would ping: %s with timeout %s and count %d\n", *ip, *timeout, *count)
// }

/*
	例二：

	$ chat --help
		usage: chat [<flags>] <command> [<flags>] [<args> ...]

		A command-line chat application.

		Flags:
		--help              Show help.
		--debug             Enable debug mode.
		--server=127.0.0.1  Server address.

		Commands:
		help [<command>]
			Show help for a command.

		register <nick> <name>
			Register a new user.

		post [<flags>] <channel> [<text>]
			Post a message to a channel.

	$ chat help post
		usage: chat [<flags>] post [<flags>] <channel> [<text>]

		Post a message to a channel.

		Flags:
		--image=IMAGE  Image to post.

		Args:
		<channel>  Channel to post to.
		[<text>]   Text to post.

	$ chat post --image=~/Downloads/owls.jpg pics
		...
*/
var (
	app      = kingpin.New("cmd", "A command-line chat application.")
	debug    = app.Flag("debug", "Enable debug mode.").Bool()
	serverIP = app.Flag("server", "Server address.").Default("127.0.0.1").IP()

	register     = app.Command("register", "Register a new user.")
	registerNick = register.Arg("nick", "Nickname for user.").Required().String()
	registerName = register.Arg("name", "Name of user.").Required().String()

	post        = app.Command("post", "Post a message to a channel.")
	postImage   = post.Flag("image", "Image to post.").File()
	postChannel = post.Arg("channel", "Channel to post to.").Required().String()
	postText    = post.Arg("text", "Text to post.").Strings()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	// Register user

	// case "register": 等价
	case register.FullCommand():
		println(*registerNick)

	// Post message
	case post.FullCommand():
		if *postImage != nil {
		}
		text := strings.Join(*postText, " ")
		println("Post:", text)
	}
}
