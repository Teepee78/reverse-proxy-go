package config

func cleanStaticDir() {
	lastCharIndex := len(Vars.StaticDir) - 1

	if Vars.StaticDir[lastCharIndex] == '/' {
		Vars.StaticDir = Vars.StaticDir[:lastCharIndex]
	}
}
