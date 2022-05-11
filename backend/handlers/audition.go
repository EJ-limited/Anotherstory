package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/brianloveswords/airtable"
	"github.com/ej-limited/auditions/pkg/mail"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type AuditionHandler struct {
	l      *log.Logger
	mail   *mail.MailClient
	client *airtable.Client
}

func NewAuditionHandler(m *mail.MailClient, c *airtable.Client) *AuditionHandler {
	l := log.New(os.Stdout, "{AUDITIONHANDLER}", log.LstdFlags)
	return &AuditionHandler{l, m, c}
}

type AuditionRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type AirtableRecord struct {
	airtable.Record
	Fields AuditionRequest
}

type RequestErr struct {
	Msg string `json:"msg"`
}

type RequestResponse struct {
	Msg string `json:"msg"`
}

func (ar *AuditionRequest) Validate() error {
	return validation.ValidateStruct(&ar,
		validation.Field(&ar.FirstName, validation.Required),
		validation.Field(&ar.LastName, validation.Required),
		validation.Field(&ar.PhoneNumber, validation.Required),
		validation.Field(ar.Email, is.EmailFormat),
	)
}

func (ah *AuditionHandler) SignUP(rw http.ResponseWriter, r *http.Request) {
	ah.l.Println("Auditon request recieved ")
	req := AuditionRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		ah.l.Printf("unable to decode json, %s", err.Error())
		json.NewEncoder(rw).Encode(RequestErr{"unable to process request"})
		return
	}
	aRecord := AirtableRecord{Fields: req}
	// aRecord.Fields.FirstName = req.FirstName
	// aRecord.Fields.LastName = req.LastName
	// aRecord.Fields.Email = req.Email
	// aRecord.Fields.PhoneNumber = req.PhoneNumber

	auditionsTable := ah.client.Table("Audition Signups")
	err = auditionsTable.Create(&aRecord)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		ah.l.Printf("unable to insert aritable record %s", err.Error())
		ah.l.Printf("%v", aRecord)
		json.NewEncoder(rw).Encode(RequestErr{"unable to sign you up at this time"})
		return
	}
	err = ah.mail.SendConfirmEmail(req.Email, aRecord.Record.ID, aRecord.Fields.FirstName)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		ah.l.Printf("unable to send email  %s", err.Error())
		json.NewEncoder(rw).Encode(RequestErr{"unable to sign you up at this time"})
		return
	}
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(RequestResponse{"success pls check your email"})

}
