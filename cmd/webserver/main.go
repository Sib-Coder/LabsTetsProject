package main

import (
	pages "LabsTetsProject/internal/delivery/http"
	database "LabsTetsProject/internal/repository"
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
	http.HandleFunc("/search_genger", pages.ExtractAllUserHTTPGender)    //work
	http.HandleFunc("/search_status", pages.ExtractAllUserHTTPStatus)    //work
	http.HandleFunc("/all_data_user_desc", pages.ExtractAllUserHTTPDESC) // work
	http.HandleFunc("/all_data_user_asc", pages.ExtractAllUserHTTPASC)   // work
	http.HandleFunc("/limit", pages.ExtractAllUserHTTPLimit)             //work требудет доработки с передачей параметра
	http.HandleFunc("/offset", pages.ExtractAllUserHTTPOffset)			//work требудет доработки с передачей параметра
	//осталось реализовать запрос с (LIMIT 10;)  на вывод строк
	//осталось реализовать с лимитом и сдвигом
	//реализовать /добавление /удаление /по какому-то идентификатору /редактирование
	http.ListenAndServe(":8090", nil)

}
