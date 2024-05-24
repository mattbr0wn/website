package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type HeadData struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Keywords    string `yaml:"keywords"`
	Author      string `yaml:"author"`
	SocialImg   string `yaml:"social_img"`
	WebsiteUrl  string `yaml:"website_url"`
}

func HeadConfig() (HeadData, error) {
	configData, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return HeadData{}, err
	}

	var headData HeadData
	err = yaml.Unmarshal(configData, &headData)
	if err != nil {
		return HeadData{}, err
	}

	return headData, nil
}
