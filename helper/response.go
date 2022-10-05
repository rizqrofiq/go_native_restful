package helper

import "net/http"

func Json(res http.ResponseWriter, httpCode int, message []byte) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(message)
}

func JsonError(res http.ResponseWriter, httpCode int, message string) {
	Json(res, httpCode, []byte(`{"success": false, "message": "`+message+`"}`))
}
