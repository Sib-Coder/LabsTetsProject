package main

import (
	"awesomeProject9/database"
	"awesomeProject9/pages"
	"net/http"
)

func main() {
	database.Conect()
	http.HandleFunc("/", pages.Hello)
	http.HandleFunc("/data_user", pages.ReceiveSend)
	http.HandleFunc("/update_data", pages.UpdateDAtaHTTP)
	http.HandleFunc("/delete_data", pages.DeleteDAtaHTTP)
	http.HandleFunc("/add_data", pages.AddDAtaHTTP)
	http.HandleFunc("/all_data_user", pages.ExtractAllUserHTTP)
	//реализовать /добавление /удаление /по какому-то идентификатору /редактирование
	http.ListenAndServe(":8090", nil)

}
