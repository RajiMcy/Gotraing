package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/raji802/gotraining/pivot/config"
	"github.com/raji802/gotraining/pivot/module"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

var t *template.Template

var store = sessions.NewCookieStore([]byte("super-secret-get-this-value-from-config"))

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello World!</h1>")

}
func login(w http.ResponseWriter, r *http.Request) {
	//p := homePage{Message: "Welcome to home page", Time: time.Now().Format(time.Stamp)}
	t, _ = template.ParseFiles("./template/w3temp.html")
	//t.ExecuteTemplate(w, "./template/w3temp.html", nil)
	t.Execute(w, nil)

}

func profile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing profile func")

	s := r.FormValue("attuid")
	var uid string
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Failed to connect to database")
	}
	fmt.Println(err)
	stmt := "SELECT attuid FROM employee WHERE attuid = ?"
	row := db.QueryRow(stmt, s)
	fmt.Println("row", row)
	err1 := row.Scan(&uid)
	if err1 != nil {
		fmt.Println("attuid from db:", uid)
		fmt.Println(err1)
		fmt.Println("User does not exist in database")
		t, _ = template.ParseFiles("./template/w3temp.html")
		t.Execute(w, "User does not exist. Kindly contact admin.")
		return
	} else {
		fmt.Println("Attuid:", s)
		session, _ := store.Get(r, "session")
		session.Values["userID"] = s
		session.Save(r, w)
		t, _ = template.ParseFiles("./template/profile.html")
		sk := module.Skills(s)
		//fmt.Println(sk.Items)
		//var v = detail{Attuid: s, Skills: []string{"Python", "Go", "Snowflake", "AWS"}}
		//t.ExecuteTemplate(w, "profile.html", v)
		t.Execute(w, sk)
	}

}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("***Running logout handler*****")
	session, _ := store.Get(r, "session")
	delete(session.Values, "userID")
	session.Save(r, w)
	t, _ = template.ParseFiles("./template/w3temp.html")
	t.Execute(w, "Logged out sucessfully.")
}

func main() {

	//t, _ = t.ParseGlob("./template/*.html")
	http.HandleFunc("/default", handlerFunc)

	http.Handle("/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/home/", login)
	http.HandleFunc("/profile", profile)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe("localhost:3000", context.ClearHandler(http.DefaultServeMux))
}
