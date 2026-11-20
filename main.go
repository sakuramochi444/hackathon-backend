package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"

	"kaizen/controller"
	"kaizen/dao"
)

// ① GoプログラムからMySQLへ接続
var db *sql.DB

func init() {
	db = dao.OpenDB()
}

// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controller.UserSearchController(w, r, db)
	case http.MethodPost:
		controller.UserRegisterController(w, r, db)
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	http.HandleFunc("/user", handler)

	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	closeDBWithSysCall()

	// 8000番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
func closeDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}
