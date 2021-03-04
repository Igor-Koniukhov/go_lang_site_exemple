package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)


func index(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	CheckErr(err)
	t.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	CheckErr(err)
	t.ExecuteTemplate(w, "create", nil)
}

func save_article(w http.ResponseWriter, r *http.Request){
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/go_site?charset=utf8")
	CheckErr(err)
	defer db.Close()

	insert, err := db.Query(fmt.Sprintf(" INSERT INTO `articles` (`title`, `anons`, `full_text`) VALUES ('%s', '%s','%s') ", title, anons, full_text))
	CheckErr(err)

	defer insert.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)


}


func handleFunc(){
	//нижняя строка кода необходима для подлючения стилей в header по пути static/css/main.css (а так же js, img)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save_article", save_article)
	http.ListenAndServe(":8080", nil)

}
func main(){
	handleFunc()
}


func CheckErr(err error){
	if err !=nil {
		fmt.Println(err.Error())
	}
}