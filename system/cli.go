package system

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
	"gopkg.in/yaml.v2"
)

const PROJECT_DEFAULT = "__empty__"

type Config struct {
	Editor  string `yaml:"editor" env:"WORKON_EDITOR" env-default:"neovim"`
	Profile string `yaml:"profile" env:"WORKON_TERM_PROFILE" env-default:"default"`
	Path    string `yaml:"root-path" env:"WORKON_ROOTPATH" env-default:"/"`
}

func LoadConfig(cfg *Config) error {
	return cleanenv.ReadConfig("config.yml", cfg)
}

func (cfg *Config) Create() error {
	config_file, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	path, path_err := filepath.Abs("./config.yml")
	if path_err != nil {
		return path_err
	}

	return ioutil.WriteFile(path, config_file, fs.ModeAppend)
}

func SelectDirectory(options []fs.FileInfo) (fs.FileInfo, error) {
	names := []string{}
	for _, option := range options {
		names = append(names, option.Name())
	}

	// ch := chooser.NewChooser(5, 40) // height, width
	// choice := ch.Choose("", names)  // options

	// fmt.Println("You Chose:", choice)
	// for _, option := range options {
	// 	if choice == option.Name() {
	// 		return option, nil
	// 	}
	// }

	return nil, errors.New("Failed to find option user picked in options array")
}

func Scan(directory string, c chan error) {
	fmt.Println("Scanning for new projects...")
	directories := []string{}
	err := filepath.Walk(directory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				fmt.Println(path, info.Size())
				directories = append(directories, path)
			}
			return nil
		})
	if err != nil {
		c <- err
	}

	//TODO: Overwrite the old list with the new list
	fmt.Println(len(directories))

	c <- nil
}
