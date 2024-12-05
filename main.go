package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var users = map[string]string{
	"student1": "password123",
	"student2": "password456",
	"student3": "password789",
}

func loginFormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Login</title>
        </head>
        <body>
            <form method="post" action="/login">
                <label for="username">Username:</label>
                <input type="text" id="username" name="username"><br>
                <label for="password">Password:</label>
                <input type="password" id="password" name="password"><br>
                <input type="submit" value="Login">
            </form>
        </body>
        </html>
    `
	t, _ := template.New("login").Parse(tmpl)
	t.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		user := r.FormValue("username")
		pass := r.FormValue("password")
		if validateCredentials(user, pass) {
			http.Redirect(w, r, "/welcome", http.StatusSeeOther)
			return
		}
		http.Error(w, "Unauthorized: Invalid credentials", http.StatusUnauthorized)
	}
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome! You've successfully logged in.")
}

func validateCredentials(user, pass string) bool {
	if storedPass, exists := users[user]; exists {
		return storedPass == pass
	}
	return false
}

func main() {
	http.HandleFunc("/", loginFormHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/welcome", welcomeHandler)
	http.ListenAndServe(":8080", nil)
}
