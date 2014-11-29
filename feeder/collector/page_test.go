package collector

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServerNorm(t *testing.T) {
	ns := map[string]string{
		"www.google.com":   "google.com",
		"www12.google.com": "google.com",
		"google.com":       "google.com",
		"some.google.com":  "some.google.com",
	}

	for src, dst := range ns {
		p := page{server: src}
		assert.Equal(t, dst, p.event().Domain)
	}
}

func TestTimestamp(t *testing.T) {
	p := page{createdAt: "2014-11-01 12:14:20"}
	l, _ := time.LoadLocation("Local")
	assert.Equal(t, time.Date(2014, 11, 1, 12, 14, 20, 0, l),
		p.event().Timestamp)
}

func TestActionSimple(t *testing.T) {
	p := page{
		method:     "GET",
		controller: "users",
		action:     "new",
		data:       sql.NullString{String: "huh"},
	}
	assert.Equal(t, "GET users/new", p.description())
}

func TestActionNewForm(t *testing.T) {
	p := page{
		method:     "GET",
		controller: "saved_forms",
		action:     "new",
		params:     "record_id=1&form_id=123&other=whatever",
	}
	assert.Equal(t, "GET saved_forms/new/123", p.description())
}

func TestActionCreateForm(t *testing.T) {
	data := `--- '{"form_id"=>"123", "record_id"=>"1", ` +
		`"something"=>"else\#", "action"=>"create", ` +
		`"controller"=>"saved_forms"}'` + "\r\n"
	p := page{
		method:     "POST",
		controller: "saved_forms",
		action:     "create",
		data:       sql.NullString{String: data},
	}
	assert.Equal(t, "POST saved_forms/create/123", p.description())
}
