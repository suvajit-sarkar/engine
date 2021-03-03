package world

import (
	"log"
	"time"
)

//World model for world detials
type World struct {
	WorldName         string          `json:"worldName"`
	WorldID           int             `json:"worldID"`
	WorldTime         int             `json:"worldTime"`
	WorldCreationTime string          `json:"worldCreationTime"`
	WorldPopulation   int             `json:"worldPopulation"`
	WorldType         string          `json:"worldType"`
	WorldSize         int             `json:"worldSize"`
	WorldUserName     []string        `json:"worldUserName"`
	GridDetails       map[string]Node `json:"gridDetails"`
	WorldTicker       *time.Ticker
}

// NewWorld is a overloaded constructor
func NewWorld(name string, worldType string) World {
	log.Printf("Creating the world %s", name)
	return World{
		WorldName:       name,
		WorldTime:       0,
		WorldPopulation: 0,
		WorldType:       worldType,
		WorldTicker:     time.NewTicker(1 * time.Second),
	}

}

//AddUser to world
func (world *World) AddUser(userName string) {
	world.WorldUserName = append(world.WorldUserName, userName)
}

//NodeTile Contains enum of tiles
type NodeTile int

const (
	CityTile NodeTile = iota
	SmallResoureNodeFood
	MediumResourceNodeFood
	LargeResourceNodeFood
	SmallResoureNodeIron
	LargeResourceNodeIron
	MediumResourceNodeIron
	SmallResoureNodeStone
	MediumResourceNodeStone
	LargeResourceNodeStone
	SmallResoureNodeWood
	MediumResourceNodeWood
	LargeResourceNodeWood
	SmallResoureNodeGold
	MediumResourceNodeGold
	LargeResourceNodeGold
	SmallResourceNodeFOODWOOD
	MediumResourceNodeFOODWOOD
	LargeResourceNodeFOODWOOD
	SmallResourceNodeStoneIronGold
	MediumResourceNodeStoneIronGold
	LargeResourceNodeStoneIronGold
	NPCTiles
)

//City struct
type City struct {
	Coordinate string   `json:"coordinate"`
	FoodRate   int      `json:"foodRate"`
	IronRate   int      `json:"ironRate"`
	GoldRate   int      `json:"goldRate"`
	StoneRate  int      `json:"stoneRate"`
	WoodRate   int      `json:"woodRate"`
	Food       int      `json:"food"`
	Iron       int      `json:"iron"`
	Gold       int      `json:"gold"`
	Stone      int      `json:"stone"`
	Wood       int      `json:"wood"`
	Attack     int      `json:"attack"`
	Defence    int      `json:"defence"`
	NodeType   NodeTile `json:"nodetype"`
}

//Resource structure
type Resource struct {
	Coordinate string   `json:"cordinate"`
	Type       NodeTile `json:"type"`
	FoodRate   int      `json:"foodRate"`
	IronRate   int      `json:"ironRate"`
	GoldRate   int      `json:"goldRate"`
	StoneRate  int      `json:"stoneRate"`
	WoodRate   int      `json:"woodRate"`
	Attack     int      `json:"attack"`
	Defence    int      `json:"defence"`
}

//Node Contains the struct of the tile
type Node struct {
	Coordinate string      `json:"coordinate"`
	Owner      int         `json:"owner"`
	Object     interface{} `json:"object"`
}

//NewNode creates and returns a new instance of Node
func NewNode(coordinate string, obj interface{}) *Node {
	return &Node{
		Coordinate: coordinate,
		Owner:      0,
		Object:     obj,
	}
}

//NewCity creates and returns a new instance of Node
func NewCity(coordinate string) *City {
	return &City{
		Coordinate: coordinate,
		FoodRate:   10,
		IronRate:   10,
		GoldRate:   10,
		StoneRate:  10,
		WoodRate:   10,
		Food:       100,
		Iron:       100,
		Gold:       100,
		Stone:      100,
		Wood:       100,
		Attack:     100,
		Defence:    100,
		NodeType:   0,
	}
}

//NewResourceType1 creates and returns a new instance of Node
func NewResourceType1(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   10,
		IronRate:   0,
		GoldRate:   0,
		StoneRate:  0,
		WoodRate:   0,
		Attack:     50,
		Defence:    50,
	}
}

//NewResourceType2 creates and returns a new instance of Node
func NewResourceType2(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   15,
		IronRate:   0,
		GoldRate:   0,
		StoneRate:  0,
		WoodRate:   0,
		Attack:     100,
		Defence:    100,
	}
}

//NewResourceType3 creates and returns a new instance of Node
func NewResourceType3(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   20,
		IronRate:   0,
		GoldRate:   0,
		StoneRate:  0,
		WoodRate:   0,
		Attack:     150,
		Defence:    150,
	}
}

//NewResourceType4 creates and returns a new instance of Node
func NewResourceType4(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   0,
		IronRate:   10,
		GoldRate:   0,
		StoneRate:  0,
		WoodRate:   0,
		Attack:     50,
		Defence:    50,
	}
}

//NewResourceType5 creates and returns a new instance of Node
func NewResourceType5(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   0,
		IronRate:   15,
		GoldRate:   0,
		StoneRate:  0,
		WoodRate:   0,
		Attack:     100,
		Defence:    100,
	}
}

//NewResourceType6 creates and returns a new instance of Node
func NewResourceType6(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   0,
		IronRate:   20,
		GoldRate:   0,
		StoneRate:  0,
		WoodRate:   0,
		Attack:     150,
		Defence:    150,
	}
}

//NewResourceType7 creates and returns a new instance of Node
func NewResourceType7(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   0,
		IronRate:   0,
		GoldRate:   0,
		StoneRate:  10,
		WoodRate:   0,
		Attack:     50,
		Defence:    50,
	}
}

//NewResourceType8 creates and returns a new instance of Node
func NewResourceType8(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   0,
		IronRate:   0,
		GoldRate:   0,
		StoneRate:  15,
		WoodRate:   0,
		Attack:     100,
		Defence:    100,
	}
}

//NewResourceType9 creates and returns a new instance of Node
func NewResourceType9(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   0,
		IronRate:   0,
		GoldRate:   0,
		StoneRate:  20,
		WoodRate:   0,
		Attack:     150,
		Defence:    150,
	}
}

//NewResourceType10 creates and returns a new instance of Node
func NewResourceType10(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   0,
		IronRate:   0,
		GoldRate:   0,
		StoneRate:  0,
		WoodRate:   10,
		Attack:     50,
		Defence:    50,
	}
}

//NewResourceType11 creates and returns a new instance of Node
func NewResourceType11(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   0,
		IronRate:   0,
		GoldRate:   0,
		StoneRate:  0,
		WoodRate:   15,
		Attack:     100,
		Defence:    100,
	}
}

//NewResourceType12 creates and returns a new instance of Node
func NewResourceType12(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   0,
		IronRate:   0,
		GoldRate:   0,
		StoneRate:  0,
		WoodRate:   20,
		Attack:     150,
		Defence:    150,
	}
}

//NewResourceType13 creates and returns a new instance of Node
func NewResourceType13(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   0,
		IronRate:   0,
		GoldRate:   10,
		StoneRate:  0,
		WoodRate:   0,
		Attack:     50,
		Defence:    50,
	}
}

//NewResourceType14 creates and returns a new instance of Node
func NewResourceType14(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   0,
		IronRate:   0,
		GoldRate:   15,
		StoneRate:  0,
		WoodRate:   0,
		Attack:     100,
		Defence:    100,
	}
}

//NewResourceType15 creates and returns a new instance of Node
func NewResourceType15(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   0,
		IronRate:   0,
		GoldRate:   20,
		StoneRate:  0,
		WoodRate:   0,
		Attack:     150,
		Defence:    150,
	}
}

//NewResourceType16 creates and returns a new instance of Node
func NewResourceType16(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   10,
		IronRate:   0,
		GoldRate:   0,
		StoneRate:  0,
		WoodRate:   10,
		Attack:     100,
		Defence:    100,
	}
}

//NewResourceType17 creates and returns a new instance of Node
func NewResourceType17(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   15,
		IronRate:   0,
		GoldRate:   0,
		StoneRate:  0,
		WoodRate:   15,
		Attack:     150,
		Defence:    150,
	}
}

//NewResourceType18 creates and returns a new instance of Node
func NewResourceType18(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   20,
		IronRate:   0,
		GoldRate:   0,
		StoneRate:  0,
		WoodRate:   20,
		Attack:     200,
		Defence:    200,
	}
}

//NewResourceType19 creates and returns a new instance of Node
func NewResourceType19(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   0,
		IronRate:   10,
		GoldRate:   10,
		StoneRate:  10,
		WoodRate:   0,
		Attack:     150,
		Defence:    150,
	}
}

//NewResourceType20 creates and returns a new instance of Node
func NewResourceType20(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   0,
		IronRate:   15,
		GoldRate:   15,
		StoneRate:  15,
		WoodRate:   0,
		Attack:     200,
		Defence:    200,
	}
}

//NewResourceType21 creates and returns a new instance of Node
func NewResourceType21(coordinate string) *Resource {
	return &Resource{
		Coordinate: coordinate,
		FoodRate:   0,
		IronRate:   20,
		GoldRate:   20,
		StoneRate:  20,
		WoodRate:   0,
		Attack:     250,
		Defence:    250,
	}
}

// VaccantCityList Contains list of vaccant city cordinates
type VaccantCityList struct {
	ListObj []string `json:"listObj"`
}

//UserDetails is a struct to store the whole user info top-bottom
type UserDetails struct {
	UserName  string                       `json:"userName"`
	WorldTime int                          `json:"worldTime"`
	WorldList map[string]*UserWorldDetails `json:"worldList"`
}

//UserWorldDetails is a struct to store the whole user info top-bottom
type UserWorldDetails struct {
	CityList map[string]*City `json:"cityList"`
}
