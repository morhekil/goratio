package collector

import (
	"encoding/json"
	"log"
	"net/url"
	"strings"
)

type mapper func(*page) string

func mapDataFormID(p *page) string {
	psf := postSavedForm{}
	return psf.get(p)
}

func mapParamsFormID(p *page) string {
	v, _ := url.ParseQuery(p.params)
	return v["form_id"][0]
}

var mappers = map[string]mapper{
	"POST saved_forms/create": mapDataFormID,
	"GET saved_forms/new":     mapParamsFormID,
}

type postSavedForm struct {
	ID string `json:"form_id"`
}

func (psf *postSavedForm) get(p *page) string {
	p.mapFormData(&psf)
	return psf.ID
}

func (p *page) mapFormData(d interface{}) {
	data := strings.Replace(
		strings.Trim(p.data.String, "' -\r\n"),
		"=>", ": ", -1)
	err := json.Unmarshal([]byte(data), &d)
	if err != nil {
		log.Fatal(err)
	}
}
