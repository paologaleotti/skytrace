package main

import (
	"net/http"
	"skytrace/internal/api"
	"skytrace/pkg/util"
)

func main() {
	router := api.InitService()

	util.Log.Info("Server started at http://localhost:3000")
	err := http.ListenAndServe("0.0.0.0:3000", router)
	if err != nil {
		util.Log.Fatalf("Could not start http server: %v\n", err)
	}

}
