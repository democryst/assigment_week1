package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// StockListResponse struct stock list object
type StockListResponse struct {
	StockList []StockListBody `json:"stock_list"`
}

// StockListBody struct stock object
type StockListBody struct {
	ID       uint32  `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity uint16  `json:"quantity"`
}

// DefaultStockControllers Response Default Stock Controller
func DefaultStockControllers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//1. Check Stock
	//2. Response Stock
	fmt.Fprintf(w, "Welcome to stock management on %s!", r.URL.Path[1:])
}

// StockListControllers Response Stock List Controller
func StockListControllers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	stockListResponse := StockListResponse{}
	stockList := []StockListBody{}
	//1. Check Stock
	queryStmt := `SELECT id, name, price, quantity FROM products WHERE quantity >= 1;`
	rows, err := db.Query(queryStmt)
	if err != nil {
		// panic(err)
		switch err {
		case sql.ErrNoRows:
			//locally handle SQL error, abstract for caller
			fmt.Println("No Data Record")
		default:
			fmt.Println("Error Query")
			// log.Fatal(err)
			panic(err)
		}
	}
	defer rows.Close()
	for rows.Next() {
		var id uint32
		var name string
		var price float32
		var quantity uint16
		err = rows.Scan(&id, &name, &price, &quantity)
		if err != nil {
			// handle this error
			fmt.Println("Error Get Data From Rows" + err.Error())
			// log.Fatal(err)
			panic(err)
		}
		stockList = append(stockList, StockListBody{id, name, price, quantity})
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		// log.Fatal(err)
		fmt.Println("Error Iteration Rows" + err.Error())
		panic(err)
	}

	//2. Response Stock
	//fmt.Fprintf(w, "Welcome to stock List on %s!", r.URL.Path[1:])

	// stockListBody1 := StockListBody{"1", "Test1", 1.00, 1}
	// stockListBody2 := StockListBody{"2", "Test2", 2.00, 1}
	// stockListBody3 := StockListBody{"3", "Test3", 3.00, 1}
	// stockList := []StockListBody{stockListBody1, stockListBody2, stockListBody3}

	// for i := 0; i < 10; i++ {
	// 	id := fmt.Sprintf("%d", i+1)
	// 	stockListBody := StockListBody{id, "Test" + id, float32(i + 1), uint16(i + 1)}
	// 	stockList = append(stockList, stockListBody)
	// }
	stockListResponse.StockList = stockList

	jsonResponse, err := json.Marshal(stockListResponse)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		// return
		fmt.Println("Error Marshal JSON result" + err.Error())
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
