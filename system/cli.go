package system

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
	"gopkg.in/yaml.v2"
)

const PROJECT_DEFAULT = "__empty__"
const NEW_PROJECT_PROMPT = "__new"

type Config struct {
	Editor             string   `yaml:"editor" env:"WORKON_EDITOR" env-default:"neovim"`
	Profile            string   `yaml:"profile" env:"WORKON_TERM_PROFILE" env-default:"default"`
	Path               string   `yaml:"root-path" env:"WORKON_ROOTPATH" env-default:"/"`
	SupportedLanguages []string `yaml:"supported-languages" env:"WORKON_SUPPORTED_LANGUAGES"`
}

func LoadConfig(cfg *Config) error {
	return cleanenv.ReadConfig("config.yml", cfg)
}

func (cfg *Config) Save() error {
	config_file, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	path, path_err := filepath.Abs("config.yml")
	if path_err != nil {
		return path_err
	}

	return ioutil.WriteFile(path, config_file, fs.ModeAppend)
}

func Scan(directory string, c chan error, index *InternalIndex) {
	fmt.Println("Scanning for new projects...")
	directories := []string{}
	err := filepath.Walk(directory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				// fmt.Println(path, info.Size())
				directories = append(directories, path)
			}
			return nil
		})
	if err != nil {
		c <- err
		return
	}

	index.Directories = directories
	err = index.Overwrite()
	if err != nil {
		c <- err
		return
	}

	c <- nil
}

func CreateProject(name string, path string, project_type string) error {
	// Create the project with the path
	err := os.Mkdir(path+name, os.ModeAppend)
	if err != nil {
		return err
	}
	// Run startup script with the name
	fmt.Printf("Running startup script for %s", project_type)
	fmt.Println()
	return nil
}
