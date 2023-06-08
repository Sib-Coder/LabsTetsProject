package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

type UserInfo struct { // кусок пакета модели
	Name      string
	LastName  string
	SurName   string
	Gender    string
	Status    string
	DateBirth string
	DateAdded string
}

// подключение к бд
func Conect() {
	db, err := sqlx.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db = db
	//defer Db.Close() //Закрытие БД
	fmt.Println("DataBase_is_WORK")
}

// получение пользователя из бд
func ExtractUserData() UserInfo {
	var u UserInfo
	res, err := Db.Query("SELECT name, lastname, surname,status FROM employees WHERE name='Daniil';")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Name, &u.LastName, &u.SurName, &u.Status)
		if err != nil {
			panic(err)
		}
	}
	result := u
	return result

}

// надо переделать на передачу структуры и везде работать через вызов полей структуры
func AddUserData(u UserInfo) bool {
	var count_users int

	res, err := Db.Query("SELECT COUNT(id) FROM employees WHERE name = $1 AND lastname = $2 AND surname = $3;", u.Name, u.LastName, u.SurName)
	fmt.Println(res)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&count_users)
		//fmt.Println(count_users)
		if err != nil {
			panic(err)
		}
	}

	if count_users == 0 {
		result, err := Db.Exec("insert into employees (name, lastname, surname, gender, status ) values ($1, $2, $3, $4, $5);", u.Name, u.LastName, u.SurName, u.Gender, u.Status)
		if err != nil {
			panic(err)
		}
		fmt.Println("Пользователь успешно добавлен")
		fmt.Println(result.RowsAffected()) // количество затронутых строк
		return true
	} else {
		return false
	}

}

// подумать над параметрами обновлений
func UpdateUser(u UserInfo) {
	result, err := Db.Exec("UPDATE employees set  surname =$3, gender = $4 WHERE name =$1 AND lastname =$2 ;", u.Name, u.LastName, u.SurName, u.Gender)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Пользователь обновлён успешно")
		fmt.Println(result.RowsAffected()) // количество затронутых строк
	}

}

// удаление пользователя на основе его статуса и имени
func DeleteUser(u UserInfo) {
	result, err := Db.Exec("DELETE FROM employees WHERE name = $1 and status = $2 and lastname =$3", u.Name, u.Status, u.LastName)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Пользователь удалён успешно")
		fmt.Println(result.RowsAffected())
	}

}
func main() {
	user := UserInfo{Name: "Anna", LastName: "Sinitsyna", SurName: "25-17Pipa", Gender: "woman", Status: "active", DateBirth: "2008-05-04", DateAdded: "2023-06-08"}
	Conect()
	dan := ExtractUserData()
	fmt.Println(dan)
	AddUserData(user)
	UpdateUser(user)
	//DeleteUser(user)
}
