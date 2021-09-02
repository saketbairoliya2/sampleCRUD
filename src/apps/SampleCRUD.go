package main

import (
	"github.com/saketbairoliya2/sampleCRUD/src/backend/cmd"
	"github.com/saketbairoliya2/sampleCRUD/src/backend/executors"
	"github.com/saketbairoliya2/sampleCRUD/src/configs"
	"log"
	"net/http"
)

func init(){
	log.Println("---Loading configs---")
	configs.LoadConfigs()

	log.Println("--- Init Registration Manager ---")
	executors.InitRegisterManager()

	log.Println("--- Out of main func ---")
}

func main(){
	http.Handle("/", cmd.Router)

	err := http.ListenAndServe(":" + configs.Config.Port, nil)
	if err != nil{
		log.Println("Err: ", err.Error())
	}
}