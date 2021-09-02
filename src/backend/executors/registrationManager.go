package executors

import "github.com/saketbairoliya2/sampleCRUD/src/backend/dao"

type RegisterManager interface {
	SaveResponse(form dao.RegistrationForm)
	GetResponse()map[string]dao.RegistrationForm
	DeleteResponse(id string)
	EditResponse(form dao.RegistrationForm)
}


var Registrar RegisterManager

func InitRegisterManager(){
	// Make an obj which impls above interface. (Note, obj can be made based on the factory layout pattern)
	obj := InMemForm{AllRecords: make(map[string]dao.RegistrationForm)}

	// Assign to Registrar
	Registrar = &obj
}


