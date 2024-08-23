package domain

type Chat struct {
	Id Id `json:"Id" bson:"_id"`

	Users       []Id `json:"usersId" bson:"usersId"`
	ChatThemeId Id   `json:"chatThemeId" bson:"chatThemeId"`

	// esto es solo para grupos
	IsGroup bool `json:"isGroup" bson:"isGroup"`

	ProfileImage string `json:"profileImage" bson:"profileImage"`
	IsPublic     bool   `json:"isPublic" bson:"isPublic"`
	ChatName     string `json:"chatName" bson:"chatName"`
}

type ChatWithUser struct {
	Chat  `bson:",inline"`
	Users []User `json:"users" bson:"users"`
}
