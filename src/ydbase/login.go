package ydbase

func Login(uname, pwd string) bool {
	if uname == config.uname && pwd == config.pwd {
		config.isLogin = true
	} else {
		config.isLogin = false
	}

	return config.isLogin
}
