package broker

import (
	"encoding/json"
	"net/http"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{
		"msg": "Welcome to e-commerce backend",
	}
	b_msg, _ := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b_msg)
}
