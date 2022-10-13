package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"honeypot/controllers"
	"honeypot/models"
	"honeypot/seed"

	"github.com/joho/godotenv"
)
var (
	appLogger  *log.Logger
    server = controllers.Server{}
    fakeServer = controllers.Server{}
)


func handleIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)
    var parsedTemplate *template.Template; 
    var err error;

    if strings.Contains(r.URL.Path, "api") {
        handleAPI(w, r)
    } else {
        if r.URL.Path == "/login" || r.URL.Path == "/login.html" {
            parsedTemplate, err = template.ParseFiles("templates/login.html")
        } else {
            parsedTemplate, err = template.ParseFiles("templates/index.html")
        }

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(err.Error()))
            log.Fatal(err)
        }

        if err := parsedTemplate.Execute(w, nil); err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(err.Error()))
            log.Fatal(err)
        }
    }

}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

type LogEntry struct {
    IP string `json:"IP"`
    Form string `json:"Form"`
    Body string `json:"Body"`
}

var logEntries []LogEntry

func handleAPI(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        var user models.User

        if r.URL.Path == "/api/login" {

            r.ParseForm()
            
            if r.FormValue("email") != "" || r.FormValue("password") != "" {
                // Send through to fake db
                _,err := user.GetUserByEmailFake(fakeServer.DB, r.FormValue("email"))

                logEntries = append(logEntries, LogEntry{IP: r.RemoteAddr, Form: fmt.Sprint(r.Form), Body: fmt.Sprint(r.Body)})

                if err != nil {
                    log.Panic(err)
                    JSON(w, http.StatusInternalServerError, err.Error())
                } else {
                    JSON(w, http.StatusInternalServerError, "token: ewiojojiewjio193nfu432i")
                }

            } else {
                _, err := user.GetUserByEmail(server.DB, r.FormValue("__e_m_a_i_l"))

                if err != nil {
                    log.Println(err)
                    JSON(w, http.StatusInternalServerError, "User not found")
                } else {
                    if err != nil {
                        log.Println(err)
                        JSON(w, http.StatusInternalServerError, "Something went wrong")
                    } else {

                        JSON(w, http.StatusOK, logEntries)
                    }
                }
            }


        } else if r.URL.Path == "/api/getUsers" {
            users, err := user.FindAllUsers(server.DB)

            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                w.Write([]byte("Error getting all users"))
                log.Fatal("Error getting all users")
            } else {
                w.WriteHeader(http.StatusOK)
                JSON(w, http.StatusOK, users)
            }

        } 
        
    } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte(""))
        log.Fatal("Method not allowed")
    }
}


func main() {
    err := godotenv.Load()

	logFile, err := os.OpenFile("./mylog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("unable to open %s\n", "./mylog.log")
		log.Fatal(err)
	}

	defer logFile.Close()
	appLogger = log.New(logFile, "", 0)

    if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

    server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

    fakeServer.Initialize(os.Getenv("OTHER_DB_DRIVER"), os.Getenv("OTHER_DB_USER"), os.Getenv("OTHER_DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("OTHER_DB_HOST"), os.Getenv("OTHER_DB_NAME"))

	seed.Load(server.DB)
	seed.Load(fakeServer.DB)

	http.HandleFunc("/", handleIndex)
    http.HandleFunc("/api", handleAPI)

    fmt.Printf("Running server on port %s\n", "8080")
	if err := http.ListenAndServe(":"+"8080", nil); err != nil {
		log.Fatal(err)
	}

}
