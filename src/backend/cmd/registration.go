package cmd

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/saketbairoliya2/sampleCRUD/src/backend/dao"
	"github.com/saketbairoliya2/sampleCRUD/src/backend/pkg"
	"io/ioutil"
	"log"
	"net/http"
)

func fetchAllRecords(writer http.ResponseWriter, r *http.Request){
	log.Println("Getting all records.")
	errList := []error{}
	message := ""

	response := pkg.FetchAllRecords()
	HandleResponse(writer, errList, message, response)
}

func saveNewRecord(writer http.ResponseWriter, r *http.Request){
	log.Println("Saving a new record.")
	errList := []error{}
	message := ""

	request, err := getRegistrationRequest(r)
	if err != nil{
		errList = append(errList, err)
	}

	if len(errList) == 0{
		pkg.SaveNewRecord(request)
	}

	HandleResponse(writer, errList, message, nil)
}

func editRecord(writer http.ResponseWriter, r *http.Request){
	log.Println("Editing a record. ")
	errList := []error{}
	message := ""

	request, err := getRegistrationRequest(r)
	if err != nil{
		errList = append(errList, err)
	}

	if len(errList) == 0 {
		pkg.EditRecord(request)
	}

	HandleResponse(writer, errList, message, nil)
}

func deleteRecord(writer http.ResponseWriter, r *http.Request){
	errList := []error{}
	message := ""

	params := mux.Vars(r)
	id := params["id"]
	log.Println("Deleting a record :", id)
	pkg.DeleteRecord(id)
	HandleResponse(writer, errList, message, nil)
}

func getRegistrationRequest(request *http.Request) (dao.RegistrationForm, error){
	if request.Body != nil {
		defer request.Body.Close()
	}

	requestObject := dao.RegistrationForm{}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Println("Error when reading request body:", err.Error())
		return requestObject, err
	}

	err = json.Unmarshal(body, &requestObject)
	if err != nil {
		log.Println("Error in UnMarshall:", err.Error())
		return requestObject, err
	}

	return requestObject, nil
}

func init(){
	Router.HandleFunc("/sampleCRUD/registration", fetchAllRecords).Methods(http.MethodGet)
	Router.HandleFunc("/sampleCRUD/registration", saveNewRecord).Methods(http.MethodPost)
	Router.HandleFunc("/sampleCRUD/registration", editRecord).Methods(http.MethodPut)
	Router.HandleFunc("/sampleCRUD/registration/{id}", deleteRecord).Methods(http.MethodDelete)
}
