package cmd

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/saketbairoliya2/sampleCRUD/src/backend/dao"
	"net/http"
)

var Router = mux.NewRouter()

// HandleResponse is generic function to get response for all the requests.
func HandleResponse(writer http.ResponseWriter, errList []error, message string, data interface{}) {
	response := dao.ResponseStruct{
		Message: message,
		Result:  data,
		Success: true,
	}

	for _, err := range errList {
		if err != nil {
			response.Message = err.Error()
			response.Success = false
			break
		}
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	// TODO: Based on the error, send correct error code. Like 4xx
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(response)
}