package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/er-vamsi/PracticeGo/TunaGoApp/web/config/db"
	"github.com/er-vamsi/PracticeGo/TunaGoApp/web/model"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/er-vamsi/PracticeGo/TunaGoApp/web/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

//for Homepage GET
func HomePageHandler(w http.ResponseWriter, request *http.Request) {
	var tmpls = template.Must(template.ParseFiles("web/templates/home.html"))
	if err := tmpls.ExecuteTemplate(w, "home.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// for GET
func RegisterPageHandler(w http.ResponseWriter, request *http.Request) {
	var tmpls = template.Must(template.ParseFiles("web/templates/signup.html"))
	if err := tmpls.ExecuteTemplate(w, "signup.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

//for POST
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()

	for key, value := range r.Form {
		fmt.Printf("%s = %s \n", key, value)
	}

	uName := r.PostFormValue("username")
	fName := r.PostFormValue("firstname")
	lName := r.PostFormValue("lastname")
	uEmail := r.PostFormValue("email")
	uPwd := r.PostFormValue("password")

	fmt.Println("#################### ", uName, " ########################")

	_uName, _fName, _lName, _uEmail, _uPwd := false, false, false, false, false

	_uName = !helpers.IsEmpty(uName)
	_fName = !helpers.IsEmpty(fName)
	_lName = !helpers.IsEmpty(lName)
	_uEmail = !helpers.IsEmpty(uEmail)
	_uPwd = !helpers.IsEmpty(uPwd)

	if _uName && _fName && _lName && _uEmail && _uPwd {
		fmt.Fprintln(w, "Username for Register : ", uName)
		fmt.Fprintln(w, "First Name : ", fName)
		fmt.Fprintln(w, "Last Name : ", lName)
		fmt.Fprintln(w, "Email : ", uEmail)
	} else {
		fmt.Fprintln(w, "This fields can not be blank!")
	}

	//var user model.User
	//user.UserName = uName

	user := &model.User{uName, uEmail, fName, lName, uPwd, ""}
	fmt.Println(user.Email)

	var res model.ResponseResult

	collection, err := db.GetDBCollection()

	if err != nil {
		log.Print("Error")
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	var result model.User
	err = collection.FindOne(context.TODO(), bson.D{{"username", user.UserName}}).Decode(&result)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

			if err != nil {
				res.Error = "Error While Hashing Password, Try Again"
				json.NewEncoder(w).Encode(res)
				return
			}
			user.Password = string(hash)

			_, err = collection.InsertOne(context.TODO(), user)
			if err != nil {
				res.Error = "Error While Creating User, Try Again"
				json.NewEncoder(w).Encode(res)
				return
			}
			res.Result = "Registration Successful"
			json.NewEncoder(w).Encode(res)
			return
		}

		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Result = "Username already Exists!!"
	json.NewEncoder(w).Encode(res)
	http.Redirect(w, r, "/login", 302)

	return

}

// for GET
func LoginPageHandler(w http.ResponseWriter, request *http.Request) {
	var tmpls = template.Must(template.ParseFiles("web/templates/login.html"))
	if err := tmpls.ExecuteTemplate(w, "login.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

//for POST
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{}
	uName := r.PostFormValue("username")
	uPwd := r.PostFormValue("password")

	user := &model.User{uName, "", "", "", uPwd, ""}
	fmt.Println("User details: ", user)

	collection, err := db.GetDBCollection()

	if err != nil {
		log.Fatal(err)
	}
	var result model.User
	var res model.ResponseResult

	err = collection.FindOne(context.TODO(), bson.D{{"username", user.UserName}}).Decode(&result)

	if err != nil {
		res.Error = "Invalid username"
		//json.NewEncoder(w).Encode(res)
		m["error"] = "Invalid username"
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))

	if err != nil {
		res.Error = "Invalid password"
		json.NewEncoder(w).Encode(res)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  result.UserName,
		"firstname": result.FirstName,
		"lastname":  result.LastName,
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		res.Error = "Error while generating token,Try again"
		json.NewEncoder(w).Encode(res)
		return
	}

	result.Token = tokenString
	result.Password = ""

	json.NewEncoder(w).Encode(result)

}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})
	var result model.User
	var res model.ResponseResult
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result.UserName = claims["username"].(string)
		result.FirstName = claims["firstname"].(string)
		result.LastName = claims["lastname"].(string)

		json.NewEncoder(w).Encode(result)
		return
	} else {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

}
