// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/suvajit-sarkar/engine/auth"
)

var addr = flag.String("addr", ":8080", "http service address")
var ctx = context.Background()

// ConfigureRouters to configure the api requests
func ConfigureRouters() *mux.Router {
	//print(utilites.Test)
	router := mux.NewRouter()

	//static file serve path
	router.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	router.HandleFunc("/", serveHome)
	hub := newHub()
	go hub.run()
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	router.HandleFunc("/authenticate", auth.Authenticate)

	// API routes are should be provided here

	// API for fetching initial data for a user in plaforms page
	router.Handle("/api/getPlaformInitData", auth.Middleware(getPlaformInitData))
	router.Handle("/api/worldlist", auth.Middleware(getWorldList))
	router.Handle("/globalChatSocket", auth.Middleware(joinGlobalChatChannel))
	router.Handle("/api/joinWorld/{name}", auth.Middleware(joinWorld))
	router.Handle("/worldWs/{token}", auth.Middleware(joinWorldSocket))

	// For now admin specfic routes
	router.Handle("/api/spawnWorld/{name}", auth.Middleware(spawnWorld))
	//router.Handle("/ws/{token}", auth.Middleware(joinWorldSocket))

	return router
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "./static/home.html")
}

// func StartEngines() {

// }

func main() {
	router := ConfigureRouters()

	//Run this at your own risk
	//world.InitWorldDetails()

	// GameInit()
	// SpawnWorlds()

	UVUSimulation()

	err := http.ListenAndServe(*addr, handlers.LoggingHandler(os.Stdout, router))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
