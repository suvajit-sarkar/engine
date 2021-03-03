package user

//NodeTile is an enum

//TODO: change this to village as village will hold resources

//User model
type User struct {
	UserID   int    `json:"userID"`
	UserName string `json:"userName"`
	//Cities   []world.City `json:"cities"`
}
