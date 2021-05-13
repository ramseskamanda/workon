package system

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
	"gopkg.in/yaml.v2"
)

type InternalIndex struct {
	MostRecent  []string `yaml:"most-recent"`
	Directories []string `yaml:"directories"`
}

func LoadInternalIndex(index *InternalIndex) error {
	return cleanenv.ReadConfig("index.yml", index)
}

func (index *InternalIndex) Overwrite() error {
	config_file, err := yaml.Marshal(index)
	if err != nil {
		return err
	}

	path, path_err := filepath.Abs("index.yml")
	if path_err != nil {
		return path_err
	}

	return ioutil.WriteFile(path, config_file, fs.ModeAppend)
}

func (index *InternalIndex) FindProjectByName(project string) (string, error) {
	for _, path := range index.Directories {
		if filepath.Base(path) == project {
			return path, nil
		}
	}

	return "", &ProjectNotFound{Name: project}
}

type ProjectNotFound struct {
	Name string
}

func (m *ProjectNotFound) Error() string {
	return fmt.Sprintf("No Project found with name: %s", m.Name)
}
