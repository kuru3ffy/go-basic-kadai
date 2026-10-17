package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Entry struct {
	Timestamp time.Time `json:"timestamp"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}

func saveDiary(entries []Entry) error {
	data, err := json.Marshal(entries)
	if err != nil {
		return fmt.Errorf("日記のJSON変換に失敗しました：%w", err)
	}
	return os.WriteFile("diary.json", data, 0644)
}

func loadDiary() ([]Entry, error) {
	data, err := os.ReadFile("diary.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Entry{}, nil
		}
		return nil, fmt.Errorf("日記の読み込みに失敗しました：%w", err)
	}
	var entries []Entry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, fmt.Errorf("日記のJSON解析に失敗しました：%w", err)
	}
	return entries, nil
}

func run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("使い方：go run main.go add \"タイトル\" \"本文\" または go run main.go list")
	}
	switch args[0] {
	case "add":
		if len(args) != 3 {
			return fmt.Errorf("使い方：go run main.go add \"タイトル\" \"本文\"")
		}
		entries, err := loadDiary()
		if err != nil {
			return err
		}
		entries = append(entries, Entry{Timestamp: time.Now(), Title: args[1], Content: args[2]})
		if err := saveDiary(entries); err != nil {
			return fmt.Errorf("日記の保存に失敗しました：%w", err)
		}
		fmt.Println("日記を保存しました。")
	case "list":
		if len(args) != 1 {
			return fmt.Errorf("使い方：go run main.go list")
		}
		entries, err := loadDiary()
		if err != nil {
			return err
		}
		if len(entries) == 0 {
			fmt.Println("日記はまだありません。")
		}
		for _, entry := range entries {
			fmt.Println("--- 日記データ ---")
			fmt.Println("日時：" + entry.Timestamp.Format("2006-01-02 15:04:05"))
			fmt.Println("タイトル：" + entry.Title)
			fmt.Println("本文：")
			fmt.Println(entry.Content)
		}
	default:
		return fmt.Errorf("不明なコマンドです：%s", args[0])
	}
	return nil
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
