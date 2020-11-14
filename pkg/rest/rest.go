package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"rest/pkg/dto"
)

func WriteAsJSONErr(w http.ResponseWriter, err error, wHeader int) {
	log.Println(err)
	data := &dto.ErrDTO{Err: err.Error()}
	WriteAsJSON(w, data, wHeader)
}

func WriteAsJSON(w http.ResponseWriter, dto interface{}, wHeader int) {
	data, err := json.Marshal(dto)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(wHeader)
	_, err = w.Write(data)
	if err != nil {
		log.Print(err)
		return
	}
}
