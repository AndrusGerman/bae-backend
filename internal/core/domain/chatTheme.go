package domain

type BackgroundType uint

const (
	BackgroundTypeSolid BackgroundType = 0
)

type ChatTheme struct {
	Id             Id             `json:"Id" bson:"_id"`
	BackgroundType BackgroundType `json:"backgroundTypeId" bson:"backgroundTypeId"`
}
