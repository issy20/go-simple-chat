package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type JsonView struct {
	w           http.ResponseWriter
	successCode int
}

func jsonView(w http.ResponseWriter, code int, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
	w.WriteHeader(code)
	return nil
}

func (j *JsonView) ErrorView(err error) {
	j.w.WriteHeader(http.StatusInternalServerError)
}

func (j *JsonView) View(i interface{}) {
	if err := jsonView(j.w, j.successCode, i); err != nil {
		j.w.WriteHeader(http.StatusInternalServerError)
	}
}

func (j *JsonView) ViewNoBody() {
	j.w.WriteHeader(j.successCode)
}

var logger = log.New(os.Stdout, "handler", log.LstdFlags)
