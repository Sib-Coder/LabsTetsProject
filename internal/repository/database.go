package database

import (
	"LabsTetsProject/internal/data"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

// подключение к бд
func Conect() {
	db, err := sqlx.Open("postgres", "host='10.10.0.136' port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Println("DataBase NOT WORK")
	}
	Db = db
	//defer Db.Close() //Закрытие БД
	fmt.Println("DataBase_is_WORK")
}

// получение пользователя из бд - готова по имени
func ExtractUserData(t string) model.UserInfo {
	var u model.UserInfo
	res, err := Db.Query("SELECT name,lastname, surname, status, gender, TO_CHAR(datebirth,'YYYY-MM-DD'), TO_CHAR(dateadded,'YYYY-MM-DD')  FROM employees WHERE name=$1;", t)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Name, &u.LastName, &u.SurName, &u.Status, &u.Gender, &u.DateBirth, &u.DateAdded)
		if err != nil {
			panic(err)
		}
	}

	return u

}

// добавление данных пользователя - готова
func AddUserData(u model.UserInfo) string {
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
		result, err := Db.Exec("insert into employees (name, lastname, surname, gender, status,datebirth,dateadded ) values ($1, $2, $3, $4, $5,$6,$7);", u.Name, u.LastName, u.SurName, u.Gender, u.Status, u.DateBirth, u.DateAdded)
		if err != nil {
			panic(err)
		}
		fmt.Println("Пользователь успешно добавлен")
		fmt.Println(result.RowsAffected()) // количество затронутых строк
		return "пользователь успешно добавлен"
	} else {
		return "ошибка добавления пользователя"
	}

}

// подумать над параметрами обновлений параметрами будет имя и фамилия - готова
func UpdateUser(u model.UserInfo) string {
	var count_users int

	res, err := Db.Query("SELECT COUNT(id) FROM employees WHERE name = $1 AND lastname = $2 AND surname = $3;", u.Name, u.LastName, u.SurName)
	fmt.Println(res)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&count_users)
		fmt.Println(count_users)
		if err != nil {
			panic(err)
		}
	}
	if count_users != 0 {
		result, err := Db.Exec("UPDATE employees set  surname =$3, gender = $4 ,status =$5,datebirth=$6,dateadded=$7 WHERE name =$1 AND lastname =$2 ;", u.Name, u.LastName, u.SurName, u.Gender, u.Status, u.DateBirth, u.DateAdded)
		if err != nil {
			panic(err)
		} else {
			return "Пользователь обновлён успешно"
			fmt.Println(result.RowsAffected()) // количество затронутых строк
		}
		return "Ошибка обновления пользователя"
	}
	return "Ошибка обновления пользователя"

}

// удаление пользователя на основе его статуса и имени и фамилии- готова
func DeleteUser(u model.UserInfo) string {
	result, err := Db.Exec("DELETE FROM employees WHERE name = $1 and status = $2 and lastname =$3", u.Name, u.Status, u.LastName)
	if err != nil {
		panic(err)
	} else {
		return "Пользователь удалён успешно"
		fmt.Println(result.RowsAffected())
	}
	return "ошибка удаления пользователя"
}

func ExtractUserDataMas() []model.UserInfo { //получение всех пользователей из бд
	var u model.UserInfo
	var u_mas []model.UserInfo
	res, err := Db.Query("SELECT name,lastname,surname,gender,status, TO_CHAR(datebirth,'YYYY-MM-DD'), TO_CHAR(dateadded,'YYYY-MM-DD') FROM employees;")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Name, &u.LastName, &u.SurName, &u.Gender, &u.Status, &u.DateBirth, &u.DateAdded)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
		u_mas = append(u_mas, u)
	}
	fmt.Println(u_mas) //пример как вырывать параметры из запроса
	return u_mas

}

func ExtractUserDataMasfForIdexGender(g model.UserInfo) []model.UserInfo { //получение всех пользователей из бд
	var u model.UserInfo
	var u_mas []model.UserInfo
	res, err := Db.Query("SELECT name,lastname,surname,gender,status, TO_CHAR(datebirth,'YYYY-MM-DD'), TO_CHAR(dateadded,'YYYY-MM-DD') FROM employees WHERE gender =$1;", g.Gender)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Name, &u.LastName, &u.SurName, &u.Gender, &u.Status, &u.DateBirth, &u.DateAdded)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
		u_mas = append(u_mas, u)
	}
	fmt.Println(u_mas) //пример как вырывать параметры из запроса
	return u_mas

}
func ExtractUserDataMasfForIdexStatus(g model.UserInfo) []model.UserInfo { //получение всех пользователей из бд
	var u model.UserInfo
	var u_mas []model.UserInfo
	res, err := Db.Query("SELECT name,lastname,surname,gender,status, TO_CHAR(datebirth,'YYYY-MM-DD'), TO_CHAR(dateadded,'YYYY-MM-DD') FROM employees WHERE status =$1;", g.Status)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Name, &u.LastName, &u.SurName, &u.Gender, &u.Status, &u.DateBirth, &u.DateAdded)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
		u_mas = append(u_mas, u)
	}
	fmt.Println(u_mas) //пример как вырывать параметры из запроса
	return u_mas

}
func ExtractUserDataMasDes() []model.UserInfo { //получение всех пользователей из бд
	var u model.UserInfo
	var u_mas []model.UserInfo
	res, err := Db.Query("SELECT name,lastname,surname,gender,status, TO_CHAR(datebirth,'YYYY-MM-DD'), TO_CHAR(dateadded,'YYYY-MM-DD') FROM employees WHERE ORDER BY datebirth DESC ;")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Name, &u.LastName, &u.SurName, &u.Gender, &u.Status, &u.DateBirth, &u.DateAdded)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
		u_mas = append(u_mas, u)
	}
	fmt.Println(u_mas) //пример как вырывать параметры из запроса
	return u_mas

}
func ExtractUserDataMasASC() []model.UserInfo { //получение всех пользователей из бд
	var u model.UserInfo
	var u_mas []model.UserInfo
	res, err := Db.Query("SELECT name,lastname,surname,gender,status, TO_CHAR(datebirth,'YYYY-MM-DD'), TO_CHAR(dateadded,'YYYY-MM-DD') FROM employees WHERE ORDER BY datebirth ASC ;")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Name, &u.LastName, &u.SurName, &u.Gender, &u.Status, &u.DateBirth, &u.DateAdded)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
		u_mas = append(u_mas, u)
	}
	fmt.Println(u_mas) //пример как вырывать параметры из запроса
	return u_mas

}
