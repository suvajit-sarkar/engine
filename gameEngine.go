package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/suvajit-sarkar/engine/world"
)

//Redis Stores
var worldStore = world.NewWorldStore()

//Game server local variables
var worldList = make(map[string]*world.World)
var vaccantCityList []string
var userMap = make(map[string]*world.UserDetails)

//GameInit initalizes the game objects
func GameInit() {
	val, error := worldStore.GetVaccantCityList(context.Background())
	if error == nil {
		vaccantCityList = val
	} else {
		print("Error")
	}
}

//InitNewUser add user to global variables
func InitNewUser(userName string, worldName string) {
	//add the new user to the worldlist
	worldList[worldName].AddUser(userName)
	cord := getNewCityCordinate()
	city := world.NewCity(cord)
	cityList := make(map[string]*world.City)
	cityList[cord] = city
	userWorldDetails := world.UserWorldDetails{CityList: cityList}
	worldList := make(map[string]*world.UserWorldDetails)
	worldList[worldName] = &userWorldDetails
	userDetails := world.UserDetails{UserName: userName, WorldList: worldList}
	userMap[userName] = &userDetails
}

//getNewCityCordinate gives coords to allocate city
func getNewCityCordinate() string {
	vaccany := len(vaccantCityList)
	newCityCoord := vaccantCityList[0]
	vaccantCityList = vaccantCityList[1 : vaccany-1]
	return newCityCoord
}

// SpawnWorlds function creates/recreates worlds from list of active worlds
func SpawnWorlds() {
	var res WorldList
	content, _ := newDao("getWorldList")
	json.Unmarshal(content, &res)
	for _, val := range res.List {
		worldObj := world.NewWorld(val.WorldName, "Normal")
		wordStore := world.NewWorldStore()
		worldList[val.WorldName] = &worldObj
		wordStore.CreateWorld(ctx, worldObj)
		print("About to start world..", val.WorldName)
		go runWorld(&worldObj)
	}
}

//RunWorld runs the world event timer
func runWorld(worldObj *world.World) {
	for {
		select {
		case <-worldObj.WorldTicker.C:
			worldObj.WorldTime++
			// worldStore.UpdateWorld(context.Background(), worldObj)
			for i := 0; i < len(worldObj.WorldUserName); i++ {
				//userStore := user.NewUserStore()
				//user, _ := userStore.FindUser(context.Background(), worldObj.WorldUserName[i])
				user := userMap[worldObj.WorldUserName[i]]
				calculateResource(user, worldObj)
				user.WorldTime = worldObj.WorldTime
				val, error := json.Marshal(user)
				if error == nil {
					clientList[user.UserName].hub.broadcast <- val
				} else {
					print("Marshal error")
				}
			}
			log.Printf("World Time passed in sec %d for %s", worldObj.WorldTime, worldObj.WorldName)
		}
	}
}

func calculateResource(user *world.UserDetails, worldObj *world.World) {
	for key, element := range user.WorldList[worldObj.WorldName].CityList {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}
}
