package routers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ErrorResponse struct error array object
type ErrorResponse struct {
	Error []ErrorBody `json:"error"`
}

// ErrorBody struct error object
type ErrorBody struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

//PostgresDB sturct
type PostgresDB struct {
	DataMapper *sql.DB
}

// Router Public
func (postgresDB PostgresDB) Router(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[0:]
	fmt.Println(fmt.Sprintf("routing to url: %s", path))
	if strings.Contains(path, "/stock") {
		StockHandler(w, r, postgresDB.DataMapper)
	} else {
		NotFoundHandler(w, r)
	}

}

// NotFoundHandler response Error
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	errorResponse := ErrorResponse{}
	errorBody := ErrorBody{"E404", "Service Not Found"}
	errorArray := []ErrorBody{errorBody}
	errorResponse.Error = errorArray

	jsonResponse, err := json.Marshal(errorResponse)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		// return
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(jsonResponse)
}
