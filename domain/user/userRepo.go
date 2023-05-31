package user

import (
	"gotest/dbc"
)

const sql_select_one = "select id,uname from users where id=?;"

const sql_select_pages = "select id,uname from users limit :offset,:limit;"

const sql_insert_one = "INSERT INTO users (id, uname) VALUES (?, ?);"

// SelectOne selects a entity by ID
func GetOne(id string) (User, error) {

	db := dbc.MySQL()
	stmt, err := db.Preparex(sql_select_one)
	if err != nil {
		return User{
			ID:    "",
			Uname: "",
		}, err
	}
	var u User
	e := stmt.Get(&u, id)
	if e != nil {
		return User{
			ID:    "",
			Uname: "",
		}, e
	}
	return u, err
}

func Getpage(offset int, limit int) ([]User, error) {
	users := make([]User, 0)
	db := dbc.MySQL()
	rows, err := db.NamedQuery(sql_select_pages, map[string]interface{}{"offset": offset, "limit": limit})
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var u User
		rows.StructScan(&u)
		users = append(users, u)
	}
	return users, err
}

// SelectOne selects a entity by ID
func PostOne(user *User) (User, error) {
	db := dbc.MySQL()
	tx, err := db.Begin()

	if err != nil {
		return *user, err
	}
	//准备sql语句
	stmt, err := tx.Prepare(sql_insert_one)
	if err != nil {
		return *user, err
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(user.ID, user.Uname)
	if err != nil {
		return *user, err
	}
	//将事务提交
	res.LastInsertId()
	tx.Commit()
	return *user, err
}
