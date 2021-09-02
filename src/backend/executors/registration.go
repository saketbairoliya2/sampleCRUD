package executors

import "github.com/saketbairoliya2/sampleCRUD/src/backend/dao"

type InMemForm struct {
	AllRecords map[string]dao.RegistrationForm
}

// SaveResponse will save this in the given form.
func (r *InMemForm)SaveResponse(form dao.RegistrationForm){
	if _, ok := r.AllRecords[form.ID]; !ok{
		// add record.
		r.AllRecords[form.ID] = form
	}
}

func(r *InMemForm) GetResponse()map[string]dao.RegistrationForm{
	return r.AllRecords
}

func (r *InMemForm) DeleteResponse(id string){
	if _, ok := r.AllRecords[id]; ok{
		delete(r.AllRecords, id)
	}
}

func(r *InMemForm)EditResponse(form dao.RegistrationForm){
	if _, ok := r.AllRecords[form.ID]; ok{
		r.AllRecords[form.ID] = form
	}
}
