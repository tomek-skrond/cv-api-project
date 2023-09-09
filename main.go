package main

import (
	"fmt"
	"log"
)

///////////// ENTRYPOINT /////////////

func main() {
	listenPort := ":3000"
	store, err := NewPostgresStore()
	if err != nil {
		fmt.Println(err)
		log.Fatal("no db 4u")
	}

	if err := store.Init(); err != nil {
		log.Fatal(err, "error initializing storage")
	}

	//fmt.Printf("%+v\n", store)

	api := NewAPIServer(listenPort, store)
	api.Run()
}
