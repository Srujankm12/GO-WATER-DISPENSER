package middlewares

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Shubhangcs/go-water-dispenser/models"
)

func LoggerMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter ,r *http.Request){
		log.Printf("The Request Type Is %v on Route %v" , r.Method , r.URL)
		next.ServeHTTP(w,r)
	})
}

func AuthorizeUser(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data , err := io.ReadAll(r.Body)
		var d models.UserModel
		if err != nil{
			panic(err)
		}
		json.Unmarshal(data , &d)
		fmt.Println(d.UserId)
		next.ServeHTTP(w,r)
	})
}