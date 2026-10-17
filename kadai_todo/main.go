package main

import "fmt"

// ToDoは、すべきこととその完了状態を管理する。
type ToDo struct {
	ID        int
	Title     string
	Completed bool
}

// CompleteはToDoを完了状態に更新する。
func (t *ToDo) Complete() {
	fmt.Printf("ID：%dのToDoを完了に更新します\n", t.ID)
	t.Completed = true
}

// printTodosは各ToDoの完了状態、ID、タイトルを表示する。
func printTodos(todos []ToDo) {
	for _, todo := range todos {
		status := "未完了"
		if todo.Completed {
			status = "完　了"
		}
		fmt.Printf("[%s]（ID：%d）%s\n", status, todo.ID, todo.Title)
	}
}

func main() {
	todos := []ToDo{
		{ID: 1, Title: "学習計画", Completed: false},
		{ID: 2, Title: "環境構築", Completed: false},
		{ID: 3, Title: "基礎文法", Completed: false},
	}

	fmt.Println("--- ToDoリスト（初期状態） ---")
	printTodos(todos)
	fmt.Println()

	// インデックスでアクセスし、スライス内の元の要素を更新する。
	for i := range todos {
		if todos[i].ID == 1 || todos[i].ID == 2 {
			todos[i].Complete()
		}
	}

	fmt.Println()
	fmt.Println("--- ToDoリスト（最終状態） ---")
	printTodos(todos)
}
