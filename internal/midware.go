package internal

import (
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
)

func midware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.RequestURI, handlerPath) && !strings.HasPrefix(r.RequestURI, "/metrics") {
			ll.Printf("\n%s %s\n%s\n%s", r.Method, r.RequestURI, prettyHeaders(r), prettyBody(r))
		}
		next.ServeHTTP(w, r)
	})
}

func prettyHeaders(r *http.Request) string {
	var h []string

	var keys []string
	for k := range r.Header {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for _, k := range keys {
		v := r.Header[k]
		header := fmt.Sprintf("%s (%d):\t%s", k, len(v), strings.Join(v, ", "))
		h = append(h, header)
	}

	if len(h) == 0 {
		h = append(h, "no headers")
	}

	headers := ""
	for _, header := range h {
		if headers == "" {
			headers = fmt.Sprintf("  * %s", header)
		} else {
			headers = fmt.Sprintf("%s\n  * %s", headers, header)
		}
	}

	return headers
}

func prettyBody(r *http.Request) string {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Sprintf("error reading body: %v", err)
	} else if len(bytes) > 0 {
		return string(bytes)
	}
	return "no body"
}
