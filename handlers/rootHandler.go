package handlers

import "net/http"

//RootHandler handles the root route
func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Asset not found \n"))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello World \n"))
}
