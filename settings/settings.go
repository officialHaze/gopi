package settings

import (
	"fmt"
)

func Generate() error {
	conf, err := readSettingsConf()
	if err != nil {
		return fmt.Errorf("error generating settings: %v", err)
	}

	MySettings = &Settings{
		ENV_FILE_NAME:  conf.Env_File_Name,
		DEF_GO_VERSION: conf.Default_Go_Version,
	}

	return nil
}

type Settings struct {
	ENV_FILE_NAME  string
	DEF_GO_VERSION string
}
