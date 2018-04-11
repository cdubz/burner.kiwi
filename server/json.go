package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Response is the root response for every api call
type Response struct {
	Success bool        `json:"success"`
	Errors  Errors      `json:"errors"`
	Result  interface{} `json:"result"`
	Meta    Meta        `json:"meta"`
}

// Errors is our error struct for if something goes wrong
type Errors struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// Meta contains our version number and by
type Meta struct {
	Version string `json:"version"`
	By      string `json:"by"`
}

func GetMeta() Meta {
	return Meta{
		Version: "0.1 Alpha",
		By:      "Hayden Woodhead",
	}
}

// NewEmailJSON generates a new email address and returns it to the caller
func (s *Server) NewEmailJSON(w http.ResponseWriter, r *http.Request) {
	e := NewEmail()
	resp := Response{Meta: GetMeta()}

	e.Address = s.eg.NewRandom()

	exist, err := s.emailExists(e.Address) // while it's VERY unlikely that the email already exists but lets check anyway

	if err != nil {
		log.Printf("JSON Index: failed to check if email exists: %v", err)
		returnJSON500(w, r, "Failed to generate email")
		return
	}

	if exist {
		log.Printf("JSON Index: email already exisists: %v", err)
		returnJSON500(w, r, "Failed to generate email")
		return
	}

	id, err := uuid.NewRandom()

	if err != nil {
		log.Printf("JSON Index: failed to generate new random id: %v", err)
		returnJSON500(w, r, "Failed to generate random id")
		return
	}

	e.ID = id.String()
	e.CreatedAt = time.Now().Unix()
	e.TTL = time.Now().Add(time.Hour * 24).Unix()

	// Mailgun takes a really long time to register a route (sometimes up to 2 seconds) so
	// we should do this out of the request thread and then update our db with the results
	go s.createRouteAndUpdate(e)

	err = s.saveNewEmail(e)

	if err != nil {
		log.Printf("JSON Index: failed to save email: %v", err)
		returnJSON500(w, r, "Failed to save email")
		return
	}

	token := s.tg.NewToken(e.ID)

	res := struct {
		Email Email  `json:"email"`
		Token string `json:"token"`
	}{
		Email: e,
		Token: string(token),
	}

	resp.Success = true
	resp.Result = res

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Printf("JSON Index: failed to marshal result var: %v", err)
		returnJSON500(w, r, "Failed to marshal response")
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(jsonResp)
}

func (s *Server) GetEmailDetailsJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["emailID"]

	e, err := s.getEmailByID(id)

	if err != nil {
		log.Printf("GetEmailDetailsJSON: failed to retrieve email from db: %v", err)
		returnJSON500(w, r, "Failed to get email details")
		return
	}

	res := Response{
		Success: true,
		Result:  e,
		Meta:    GetMeta(),
	}

	resJSON, err := json.Marshal(res)

	if err != nil {
		log.Printf("GetEmailDetailsJSON: failed to marhsal json: %v", err)
		returnJSON500(w, r, "Failed to marshal reponse")
		return
	}

	w.Write(resJSON)
}

// returnJSON500 returns json with custom error message
func returnJSON500(w http.ResponseWriter, r *http.Request, msg string) {
	resp := Response{}
	resp.Success = false
	resp.Result = nil
	resp.Meta = GetMeta()
	resp.Errors = Errors{
		Code: 500,
		Msg:  "Internal Server Error: " + msg,
	}

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Printf("JSON Index: failed to marshal error response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write(jsonResp)
}

func returnJSONError(w http.ResponseWriter, r *http.Request, status int, msg string) {
	resp := Response{}
	resp.Success = false
	resp.Result = nil
	resp.Meta = GetMeta()
	resp.Errors = Errors{
		Code: status,
		Msg:  msg,
	}

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		returnJSON500(w, r, "Failed to marshal error response")
		return
	}

	w.WriteHeader(status)
	w.Write(jsonResp)
}
