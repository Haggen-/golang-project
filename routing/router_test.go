package routing

import (
		"testing"
		"net/http"
		)

type TestModule struct {
}

func (t TestModule) GetPath() string {
	return "/test/"
}

func (t TestModule) GetVersion() string {
	return "1.0"
}

func (t TestModule) Post(res http.ResponseWriter, r *http.Request) {
} 
func (t TestModule) Get(res http.ResponseWriter, r *http.Request) {
}

func (t TestModule) Put(res http.ResponseWriter, r *http.Request) {
}

func (t TestModule) Delete(res http.ResponseWriter, r *http.Request) {
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