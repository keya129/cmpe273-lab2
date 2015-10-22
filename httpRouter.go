package main
import (
    "fmt"
    "log"
    "encoding/json"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "io"
  	"io/ioutil"

)
type Todo struct {
    Name string `json:"name"`
}
type NewTodo struct {
    Greeting string `json:"greeting"`
}

type Todos []Todo
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}
func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request,p httprouter.Params) {
  //  fmt.Fprintln(rw, "posts index")
  //String id;
  var todo Todo

  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		rw.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(rw).Encode(err); err != nil {
			panic(err)
		}
	}

/*
    todos := Todos{
    Todo{Name: r.FormValue("name")},
    Todo{Name: "Host meetup"},
}*/
var newtodo NewTodo;
newtodo.Greeting="Hello, "+todo.Name+"!"
json.NewEncoder(rw).Encode(newtodo)

}

func PostsIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
fmt.Fprintln(rw, "Welcome ",p.ByName("name"))
}
func main() {
    mux := httprouter.New()
    mux.GET("/", Index)
    mux.GET("/posts/", PostsIndexHandler)
    mux.POST("/posts", PostsCreateHandler)

  //  mux.POST("/posts", PostsCreateHandler)

    mux.GET("/hello/:name", hello)

    log.Fatal(http.ListenAndServe(":8080", mux))


}
