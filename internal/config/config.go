package config

import (
	"os"
	"encoding/json"
	"fmt"
)

type  Config struct{
	DbUrl   string   `json:"db_url"`
	UserName  string  `json:"current_user_name"`
}

const configfileName = "/.gatorconfig.json"

func getConfigPath() (string,error) {
  path,err := os.UserHomeDir()

   if err != nil {
     return "",err
   }
   return path+configfileName,nil
}

func Read() (Config,error) {
   path,err := getConfigPath()
   fmt.Println(path)
   
   if err != nil {
     return Config{},err
   }

   file,err := os.Open(path);
   
 

   if err != nil {
	return Config{},err
   }
     defer file.Close()
   var config Config
   decoder := json.NewDecoder(file)

   if err := decoder.Decode(&config);err != nil {
	return Config{},err
   }
  
   return config,nil

}


func (c *Config) SetUser(username string) error {
    if username == "" || len(username) <= 0 {
		return fmt.Errorf("No username given")
	}
    c.UserName = username

	path,err := getConfigPath()

	if err != nil {
      return err
	}

	data,err := json.MarshalIndent(c,""," ")
	if err != nil {
      return err
	}
    
	err = os.WriteFile(path,data,0644)
	if err != nil {
		return err
	}

    return nil
}

