package configs

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

var Config Configurations

func LoadConfigs(){
	configPath := flag.String("ConfigFile", "configs.json", "Path to config file. ")
	flag.Parse()

	c := Configurations{}
	err := c.setConfigFromPath(*configPath)
	if err != nil{
		log.Fatal("Error when loading func -", err.Error())
	}
}

func (c *Configurations)setConfigFromPath(filename string)error{
	log.Println("FileName for Config is:", filename)
	configFile, err := os.Open(filename)
	if configFile != nil{
		defer configFile.Close()
	}

	if err != nil{
		log.Fatal("Error loading config", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&c)
	Config = *c
	return err
}

func (c *Configurations)GetConfigFromPath(){
	Config = *c
}

