package handlersfunc

import (
	"net/http"
	"encoding/json"
	"log"
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types/container"
)

type Images struct {
	Name string `json:"name"`
}

//Create container
var CreateContainer = http.HandlerFunc( func (w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	images := Images{}

	err := decoder.Decode(&images)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	imageName := images.Name

	ctx := context.Background()
	cli,err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:imageName,
	}, nil,nil,"")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte([]byte("Container create with ID: "+resp.ID[:12])))

})

//Run container
var RunContainer = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	idContainer := r.URL.Path[len("/c/") : len("/c/") + 12]

	ctx := context.Background()
	cli,err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	if err:= cli.ContainerStart(ctx, idContainer, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Container " + idContainer + " run"))
})

//Stop container
var StopContainer = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  {
	idContainer := r.URL.Path[len("/containers/"):len("/containers/") + 12]
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStop(ctx, idContainer, nil);
		err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Container " + idContainer + " stop"))
})

//Delete conteiner
var DeleteContainer = http.HandlerFunc( func (w http.ResponseWriter, r *http.Request){

	idContainer := r.URL.Path[len("/containers/"):len("/containers/") + 12]
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerRemove(ctx, idContainer, types.ContainerRemoveOptions{});
		err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Container " + idContainer + " delete"))
})

//Information about container
var InfoContainer = http.HandlerFunc (func (w http.ResponseWriter, r *http.Request)  {
	idContainer := r.URL.Path[len("/containers/"):len("/containers/") + 12]
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	inf, err := cli.ContainerInspect(ctx, idContainer)
	if	err != nil {
		panic(err)
	}

	productsJson, err := json.Marshal(inf)
	if	err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(productsJson)
})

var ListContainer = http.HandlerFunc ( func (w http.ResponseWriter, r *http.Request) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("List running containers"))

	for _, container := range containers {
		w.Write([]byte("Container " + container.ID[:12] + " "))
	}
})