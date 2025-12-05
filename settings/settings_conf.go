package settings

import (
	"encoding/json"
	"fmt"
	"os"
	"projinit/path"
)

func readSettingsConf() (*SettingsConf, error) {
	configPath := path.Join("settings", "settings.conf.jsonc")

	f, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %v", configPath, err)
	}

	dec := json.NewDecoder(f)

	conf := &SettingsConf{}
	if err := dec.Decode(conf); err != nil {
		return nil, fmt.Errorf("error while decoding settings configuration: %v", err)
	}

	return conf, nil
}

type SettingsConf struct {
	Env_File_Name      string `json:"env_file_name"` // which env file to use
	Default_Go_Version string `json:"default_go_version"`
}
