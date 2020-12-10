package api

import (
	"go-user/internal/database/config"
	"go-user/internal/model"
	"log"
	"net/http"
)

// Get Addresses
func GetAddress(w http.ResponseWriter, r *http.Request) {
	var address model.Address
	var getAddresses []model.Model

	connection, err := config.ConnectSQL()
	defer connection.SQL.Close()

	rows, err := connection.SQL.Query("SELECT id, user_id, address, city, province FROM addresses")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(
			&address.ID,
			&address.UserID,
			&address.Address,
			&address.City,
			&address.Province,
		); err != nil {
			log.Fatal(err.Error())
		} else {
			getAddresses = append(getAddresses, address)
		}
	}

	responseWithJson(w, model.Response{
		Status:  200,
		Message: "Ok",
		Data:    []model.Model{getAddresses},
	})
}

// User address by id user
func UserAddress(w http.ResponseWriter, r *http.Request) {
	connection, _ := config.ConnectSQL()
	defer connection.SQL.Close()

	userID := r.URL.Query().Get("user_id")

	row := connection.SQL.QueryRow("SELECT * FROM addresses WHERE user_id = ?", userID)

	var address model.Address
	err := row.Scan(&address.ID, &address.UserID, &address.Address, &address.City, &address.Province)
	if err != nil {
		responseWithJson(w, model.Response{
			Status:  404,
			Message: "Not Found" + userID,
			Data:    []model.Model{address},
		})
		return
	}

	responseWithJson(w, model.Response{
		Status:  200,
		Message: "Ok",
		Data:    []model.Model{address},
	})
}

func AddUserAddress(w http.ResponseWriter, r *http.Request) {
	connection, err := config.ConnectSQL()
	defer connection.SQL.Close()

	parseErr := r.ParseForm()
	if parseErr != nil {
		panic(parseErr)
	}

	userID := r.Form.Get("user_id")
	address := r.Form.Get("address")
	city := r.Form.Get("city")
	province := r.Form.Get("province")

	_, err = connection.SQL.Exec("INSERT INTO addresses (user_id, address, city, province) values (?,?,?,?)",
		userID,
		address,
		city,
		province,
	)

	if err != nil {
		log.Print(err)
	}

	response := model.Response{
		Status:  201,
		Message: "Success",
	}

	responseWithJson(w, response)
}

func UpdateUserAddress(w http.ResponseWriter, r *http.Request) {
	connection, err := config.ConnectSQL()
	defer connection.SQL.Close()

	parseErr := r.ParseForm()
	if parseErr != nil {
		panic(parseErr)
	}

	userID := r.Form.Get("user_id")
	address := r.Form.Get("address")
	city := r.Form.Get("city")
	province := r.Form.Get("province")

	_, err = connection.SQL.Exec("UPDATE addresses SET address = ?, city = ?, province = ? WHERE user_id = ?",
		address,
		city,
		province,
		userID,
	)

	if err != nil {
		log.Print(err)
	}

	response := model.Response{
		Status:  201,
		Message: "Success",
	}

	responseWithJson(w, response)
}

func DeleteUserAddress(w http.ResponseWriter, r *http.Request) {
	connection, err := config.ConnectSQL()
	defer connection.SQL.Close()

	parseErr := r.ParseForm()
	if parseErr != nil {
		panic(parseErr)
	}

	userID := r.Form.Get("user_id")

	_, err = connection.SQL.Exec("DELETE from addresses WHERE user_id = ?",
		userID,
	)

	if err != nil {
		log.Print(err)
	}

	response := model.Response{
		Status:  201,
		Message: "Success",
	}

	responseWithJson(w, response)
}
