package pkg

import (
	"github.com/saketbairoliya2/sampleCRUD/src/backend/dao"
	"github.com/saketbairoliya2/sampleCRUD/src/backend/executors"
)

// FetchAllRecords will get all the records from the dB.
func FetchAllRecords()map[string]dao.RegistrationForm{
	allResponse := executors.Registrar.GetResponse()
	return allResponse
}

func SaveNewRecord(form dao.RegistrationForm)error{
	executors.Registrar.SaveResponse(form)
	return nil
}

func EditRecord(form dao.RegistrationForm)error{
	executors.Registrar.EditResponse(form)
	return nil
}

func DeleteRecord(ID string)error{
	executors.Registrar.DeleteResponse(ID)
	return nil
}
