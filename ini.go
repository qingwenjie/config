package config

import (
	"errors"
	"fmt"
	i "gopkg.in/ini.v1"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	Ini *ini
)

type ini struct {
	ErrorNew error
	File     *i.File
}

type Options struct {
	ConfigBaseDir string   `json:"configBaseDir"`
	ConfigDirs    []string `json:"configDirs"`
	ConfFileExts  []string `json:"confFileExts"`
}

func (s *ini) newIniFile(files []interface{}) *i.File {
	if len(files) == 0 {
		return nil
	}
	cfg, err := i.Load(files[:1][0], files[1:]...)
	if err != nil {
		s.ErrorNew = errors.New("Configuration read failure: " + err.Error())
		return nil
	}
	return cfg
}

func (s *ini) getFiles(config *Options) []interface{} {
	if len(config.ConfigDirs) == 0 {
		s.ErrorNew = errors.New("directory configuration cannot be empty")
		return nil
	}
	var files []interface{}
	for _, dir := range config.ConfigDirs {
		filesInfo, err := ioutil.ReadDir(dir)
		if err != nil {
			s.ErrorNew = errors.New("Configuration file loading failed: " + err.Error())
			return nil
		}
		for _, file := range filesInfo {
			ext := strings.Trim(strings.ToLower(filepath.Ext(file.Name())), ".")
			if inStrings(ext, config.ConfFileExts) {
				files = append(files, filepath.Join(config.ConfigBaseDir, dir, file.Name()))
			}
		}
	}
	return files
}

func NewIni(options *Options) *ini {
	if options == nil {
		options = DefaultOptions()
	}
	if options.ConfigBaseDir == "" {
		baseDir, _ := os.Getwd()
		options.ConfigBaseDir = baseDir
	}
	ini := ini{}
	//获取文件列表
	files := ini.getFiles(options)
	//获取配置
	ini.File = ini.newIniFile(files)
	if ini.ErrorNew != nil {
		fmt.Println("err:", ini.ErrorNew.Error())
		return nil
	}
	return &ini
}

func DefaultOptions() *Options {
	baseDir, _ := os.Getwd()
	return &Options{
		ConfigBaseDir: baseDir,
		ConfigDirs:    []string{"conf"},
		ConfFileExts:  []string{"ini"},
	}
}

func inStrings(str string, strs []string) bool {
	if len(strs) == 0 {
		return false
	}
	for _, v := range strs {
		if v == str {
			return true
		}
	}
	return false
}
