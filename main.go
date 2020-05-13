package main

import (
	"os"
	"os/exec"

	"flag"
	"fmt"
	"github.com/armon/go-socks5"
)

var user, pass, addr string
var daemon bool

func init() {
	flag.StringVar(&user, "u", "", "用户名，未设定表示无需验证")
	flag.StringVar(&pass, "p", "", "密码，指定用户名时必填")
	flag.StringVar(&addr, "a", ":21000", "绑定地址")
	flag.BoolVar(&daemon, "d", false, "后台进程模式")
	flag.Parse()

	if !flag.Parsed() {
		flag.Parse()
	}

	if daemon {
		cmd := exec.Command(os.Args[0], flag.Args()[1:]...)
		cmd.Start()
		fmt.Printf("%s [PID] %d running...\n", os.Args[0], cmd.Process.Pid)
		daemon = false
		os.Exit(0)
	}
}

func main() {
	conf := &socks5.Config{}

	if len(addr) == 0 {
		panic("绑定地址不能为空")
	}

	if len(user) > 0 {
		if len(pass) == 0 {
			panic("密码不能为空")
		}
		cator := socks5.UserPassAuthenticator{Credentials: socks5.StaticCredentials{
			user: pass,
		}}

		conf.AuthMethods = []socks5.Authenticator{
			cator,
		}
	}

	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	fmt.Println("启动代理服务器", addr)
	if err := server.ListenAndServe("tcp", addr); err != nil {
		panic(err)
	}
}
