package blog

import (
	"encoding/json"
	"fmt"
	"go-gemin/models/blog"
	//"log"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	//"strconv"
	//"math/rand"
	//"time"
	//
	//"go-gemin/models/blog"
	"io/ioutil"
)


type Articles struct {
	Title    string    `json:"title"`
	Post     string    `json:"post"`
}
//
//var Article []Articles
var result = []byte("")

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)
	//err := json.Unmarshal(body, &t)
	//var resp blog.ResponseResult
	//if err != nil {
	//	resp.Error = err.Error()
	//	json.NewEncoder(w).Encode(resp)
	//	return
	//}
	var resp blog.ResponseResult
	var t Articles
	err := json.Unmarshal(body, &t)
	if err != nil {
		resp.Error = err.Error()
		json.NewEncoder(w).Encode(resp)
		panic(err)
		return
	}
	title := t.Title
	post := t.Post

	if title == "" && post == ""{
		resp.Result = "Username already Exists!!"
		json.NewEncoder(w).Encode(resp)
		return
	}

 	 _, result := blog.InsertArticles(title, post)
 	 fmt.Println(result)
	if result != nil {

		resp.Result = result
		json.NewEncoder(w).Encode(result)
		fmt.Println("OKEsss")

	}


}

func ArticleGETALLHanddler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(Article)
}