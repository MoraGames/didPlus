package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

/*func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)
	//get time in nanoseconds
	t := time.Now().UnixNano()
	//connect to the db
	db, err := sql.Open("mysql", "root:root@tcp(db:3306)/benchmarks")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not connect to the database"))
	}
	defer db.Close()
	//generate a random number between 1 and 10
	// rand := int(t) % 10
	// if rand == 0 {
	// 	useless := 0
	// 	for i := 0; i < 100000000; i++ {
	// 		useless += i
	// 	}
	// }

	//calcolate the time passed
	t = time.Now().UnixNano() - t

	//save the result in the db
	_, err = db.Exec("INSERT INTO benchmarks (backType, execTime) VALUES ('go', ?)", t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not save the result: " + err.Error()))
		return
	}

	w.Write([]byte(fmt.Sprintf("go %d ns\n", t)))
}*/

var db *sql.DB
func init(){
	var err error
	db, err = sql.Open("mysql", "root:rootPassword@tcp(db:58138)/didPlus")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/index.html", homePage)
	r.HandleFunc("/signin.html", signinPage)
	r.HandleFunc("/signup.html", signupPage)
	r.HandleFunc("/api/signin", signin).Methods("POST")
	r.HandleFunc("/api/signup", signup).Methods("POST")

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", r)
}
