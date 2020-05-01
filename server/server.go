package main

import (
	"net/http"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/FlockSupport/graphql/graph"
	"github.com/FlockSupport/graphql/graph/generated"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"

	"fmt"
	"flock-support/back/proto"
	"strconv"
	"google.golang.org/grpc"
)

func add(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	a, err := strconv.ParseUint(chi.URLParam(r, "a"), 10, 64)
	
	if (err != nil){
		fmt.Println(err)
	}
	
	if (err == nil){
		fmt.Println("input: ", a)
	}
	conn, err := grpc.Dial("localhost:8005", grpc.WithInsecure())
	client := proto.NewAddServiceClient(conn)

		req := &proto.Request{A: int64(a), B: int64(a)}
		if response, err := client.Add(ctx, req); err == nil {
			fmt.Println("result: ", response.Result)
		} else {
			fmt.Println(err)
		}
}

func addUser(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	age, err := strconv.ParseUint(chi.URLParam(r, "age"), 10, 64)
	name := chi.URLParam(r, "name")
	
	if (err != nil){
		fmt.Println(err)
	}

	conn, err := grpc.Dial("localhost:8005", grpc.WithInsecure())
	client := proto.NewAddServiceClient(conn)

		req := &proto.AddUserRequest{Id: int64(id), Age: int64(age), Name: string(name)}
		if response, err := client.AddUser(ctx, req); err == nil {
			fmt.Println("result: ", response.Result)
		} else {
			fmt.Println(err)
		}


}
func getSingleUser(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	
	if (err != nil){
		fmt.Println(err)
	}

	conn, err := grpc.Dial("localhost:8005", grpc.WithInsecure())
	client := proto.NewAddServiceClient(conn)

		req := &proto.GetSingleUserRequest{Id: int64(id)}
		if response, err := client.GetSingleUser(ctx, req); err == nil {
			fmt.Println("result: ", response)
		} else {
			fmt.Println(err)
		}
}

func getAllUsers(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	conn, err := grpc.Dial("localhost:8005", grpc.WithInsecure())
	if err != nil{
		fmt.Println(err)
	}
	
	client := proto.NewAddServiceClient(conn)

		req := &proto.GetAllUsersRequest{}
		if response, err := client.GetAllUsers(ctx, req); err == nil {
			fmt.Println("result: ", response.Users)
		} else {
			fmt.Println(err)
		}


}


func main() {
	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	
	
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "example.org"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	fmt.Println("graphql server started on port 8080")

	router.Handle("/", playground.Handler("Starwars", "/query"))
	router.Handle("/query", srv)

	router.Route("/grpc/{a}", func(router chi.Router) {
		router.Get("/", add)
	  })
	router.Route("/adduser/{id}/{age}/{name}", func(router chi.Router) {
		router.Get("/", addUser)
	  })
	router.Route("/getsingleuser/{id}", func(router chi.Router) {
		router.Get("/", getSingleUser)
	  })
	router.Route("/getallusers", func(router chi.Router) {
		router.Get("/", getAllUsers)
	  })

	err := http.ListenAndServe(":8080", router)
	
	if err != nil {
		panic(err)
	}

}
