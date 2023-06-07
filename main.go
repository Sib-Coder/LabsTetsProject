package main

import (
	"awesomeProject9/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Привет это сервис, предоставляющий API для работы с данными пользователей.\n")
}

func receiveSend(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		fmt.Println("My request is: ", r.Body)
		var user model.Registr
		decode := json.NewDecoder(r.Body).Decode(&user)

		if decode != nil {
			fmt.Println(http.StatusOK) //надо найти выдачу ошибки
		}
		fmt.Println(user.Name, " ", user.LastName)
		//обработка параметров через бд
		fmt.Fprintln(w, "Полученные данные пользователя "+string(user.Name))

	}
}
func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/data_user", receiveSend)
	//реализовать /добавление /удаление /по какому-то идентификатору /редактирование
	http.ListenAndServe(":8090", nil)

}
