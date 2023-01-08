package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	e := r.ParseForm()
	log.Println(e)

	w.Header().Add("Content-Type", "application/json")
	content, _ := ioutil.ReadAll(r.Body)

	fmt.Print("content", content)

	var t TestStruct

	if err := json.Unmarshal(content, &t); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Print("value", t)

	res, err := json.Marshal(t)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

type TestStruct struct {
	RoomID int `json:"room_id"`
	UserID int `json:"user_id"`
}
