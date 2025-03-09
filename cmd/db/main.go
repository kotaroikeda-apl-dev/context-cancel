package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func queryDB(ctx context.Context, db *sql.DB) {
	query := "SELECT SLEEP(10)" // 10秒スリープ
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		fmt.Println("クエリがキャンセルされました:", err)
	} else {
		fmt.Println("クエリ成功")
	}
}

func main() {
	dsn := "root:context_password@tcp(127.0.0.1:3306)/context_database"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("データベース接続エラー:", err)
		return
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	queryDB(ctx, db)
}
