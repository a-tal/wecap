package internal

import (
	"encoding/json"
	"net/http"
)

func writeResp(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	if _, e := w.Write([]byte(msg)); e != nil {
		ll.Printf("failed to write response: %v", e)
	}
}

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	writeResp(w, 200, "hi")
}

func reportHandler(w http.ResponseWriter, r *http.Request) {
	p := newPayload(parseQuery(r.RequestURI))

	updateMetrics(p)

	bytes, err := json.Marshal(p)
	if err != nil {
		ll.Printf("failed to JSON dump: %v", err)
	} else {
		ll.Printf("%s", string(bytes))
	}

	writeResp(w, 200, "ok")
}
