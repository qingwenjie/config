package main

import (
	"fmt"
	"github.com/qingwenjie/config"
)

func main() {
	config.Ini = config.NewIni(&config.Options{
		ConfigDirs:   []string{"testdata"},
		ConfFileExts: []string{"ini"},
	})
	if config.Ini == nil {
		panic("err")
		return
	}
	t()
}

func t() {
	host := config.Ini.File.Section("App").Key("Host").String()
	fmt.Println(host)
	name := config.Ini.File.Section("App").Key("Name").String()
	fmt.Println(name)
	port, err := config.Ini.File.Section("App").Key("Port").Int()
	fmt.Println(port, err)
}
