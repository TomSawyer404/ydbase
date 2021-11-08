package ydbase

var config struct {
	uname   string
	pwd     string
	isLogin bool
}

func init() {
	config.uname = "root"
	config.pwd = "123456"
	config.isLogin = false
}
