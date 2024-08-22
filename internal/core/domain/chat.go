package domain

type Chat struct {
	Id Id `json:"Id" bson:"_id"`

	Users   []Id `json:"users" bson:"users"`
	IsGroup bool `json:"isGroup" bson:"isGroup"`

	ProfileImage string `json:"profileImage" bson:"profileImage"` // esto es solo para grupos

	ChatThemeId Id `json:"chatThemeId" bson:"chatThemeId"`
}
