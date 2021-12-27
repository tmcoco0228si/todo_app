package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

// C create
func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (
		content,
		user_id,
		created_at) values ($1, $2, $3)`
	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// R read
func GetTodo(id int) (todo Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos
	where id = $1`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)
	return todo, err
}

//複数取得
func GetTodos() (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

///特定のUserのTodo取得
func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select id,content,user_id,created_at from todos
	where user_Id = $1`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

//更新
func (t *Todo) UpdateTodo() error {
	cmd := `update todos set content =$1, user_id = $2
	where id = $3`
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//削除
func (t *Todo) DeleteTodo() error {
	cmd := `delete from todos where id = $1`
	_, err := Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
