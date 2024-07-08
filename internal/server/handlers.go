package server

import (
	"MailBeacon/internal/util"
	"encoding/json"
	"net/http"
)

func (s *Server) NewsletterSignup(w http.ResponseWriter, r *http.Request) {
	util.WriteJson(w, http.StatusOK, map[string]string{"message": "Thank you for signing up for our newsletter!"})
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	util.WriteJson(w, http.StatusOK, map[string]string{"message": "Hello, World!"})
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	util.WriteJson(w, http.StatusOK, jsonResp)
}
