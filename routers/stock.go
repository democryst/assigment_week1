package routers

import (
	"database/sql"
	"net/http"
	"strings"
	"webapp/controllers"
)

// StockHandler Manage path under "/stock"
func StockHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	path := r.URL.Path[0:]
	if strings.EqualFold(path, "/stock/list") {
		controllers.StockListControllers(w, r, db)
	} else if strings.EqualFold(path, "/stock") {
		controllers.DefaultStockControllers(w, r, db)
	} else {
		NotFoundHandler(w, r)
	}
}
