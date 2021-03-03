package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	util "github.com/suvajit-sarkar/engine/utilities"
)

//Router ultilites

//Platforms API routes are here
var getPlaformInitData = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	content, _ := newDao("getWorldList")
	var res WorldList
	log.Print(json.Unmarshal(content, &res))
	json.NewEncoder(w).Encode(res)
})
var joinGlobalChatChannel = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	hub := newHub()
	go hub.run()
	serveWs(hub, w, r)
})

var getWorldList = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	content, _ := newDao("getWorldList")
	var res WorldList
	log.Print(json.Unmarshal(content, &res))
	json.NewEncoder(w).Encode(res)
})

//WorldMap API routes are

// 1. Assign a City to user

func assignCity() {

}

var spawnWorld = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	// TODO: validate the configuration of world
	// vars := mux.Vars(r)
	// world := newWorld(vars['name'], "normal")
	// go world.run()
})

var joinWorld = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	userName := r.Header.Get("name")
	worldName := vars["name"]

	_, found := util.Find(worldList[worldName].WorldUserName, userName)
	if !found {
		InitNewUser(userName, worldName)
	}

	//TODO: Check if user is already joined to the world
	// worldAddr, _ := worldList["Rivendel"]
	// worldAddr.AddUser(10)
	// content, _ := newDao("getWorldList")
	// var res WorldList
	// log.Print(json.Unmarshal(content, &res))
	// json.NewEncoder(w).Encode(res)
})

var joinWorldSocket = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	hub := newHub()
	go hub.run()
	serveWs(hub, w, r)

})

func createUsersDummy() {
	// for i := 0; i < 1000; i++ {
	// 	cityCord := utilities.GetNewCityCordinate()
	// 	print(cityCord)
	// 	city := user.City{i, cityCord, rand.Intn(100) / 10, rand.Intn(100) / 10, rand.Intn(100) / 10, rand.Intn(100) / 10, rand.Intn(100) / 10, 100, 100, 100, 100, 100}
	// 	cities := make([]user.City, 1)
	// 	cities = append(cities, city)
	// 	userObj := user.User{i, strings.Join([]string{"Dummy", strconv.Itoa(i)}, ""), cities}
	// 	//print(userObj.UserName, userObj.FarmRate, userObj.GoldRate, userObj.IronRate, userObj.StoneRate, userObj.WoodRate)
	// 	userStore := user.NewUserStore()

	// 	userStore.UpdateUser(ctx, userObj)
	// 	// worldObj, _ := worldList["rivendel"]
	// 	// worldObj.AddUser(userObj.UserName)
	// }
	// print("Done adding dummy")
	// println(len(worldList["rivendel"].WorldUserName))

}
