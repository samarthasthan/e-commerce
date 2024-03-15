package broker

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/samarthasthan/e-commerce-backend/internal/auth"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "applciation/json")
	s := &auth.User{}
	if err := json.NewDecoder(r.Body).Decode(s); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	s.UUID = uuid.New().String()
	if validErrs := s.Validate(); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s)
}
