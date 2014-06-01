package routing

import (
		"testing"
		"net/http"
		)

type TestModule struct {
}

func (t TestModule) Version() string {
	return "1.0"
}

func (t TestModule) SubModules() []SubModule {
	return nil	
}

func (t TestModule) Post(r *http.Request, res http.ResponseWriter) error {
	return nil
}

func (t TestModule) Get(r *http.Request, res http.ResponseWriter) error {
	return nil
}

func (t TestModule) Put(r *http.Request, res http.ResponseWriter) error {
	return nil
}

func (t TestModule) Delete(r *http.Request, res http.ResponseWriter) error {
	return nil
}

var testModulesAdded = 0

func TestAddWebModule(t *testing.T) {
	mod := TestModule{}
	RegisterModule(mod)
	RegisterModule(nil)
	RegisterModule(mod)
	out, expected := NumberOfModules(),1
	if(out != expected) {
		t.Errorf("NumberOfModules after adding module=%v expected 1", out)
	} else {
		testModulesAdded++
	}
}

func TestCountNumberOfWebModules(t *testing.T) {
	out := NumberOfModules()
	if(out != testModulesAdded) {
		t.Errorf("NumberOfModules=%v, expected 0", out)
	}
}