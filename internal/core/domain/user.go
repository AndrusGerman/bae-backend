package domain

type User struct {
	Id           Id      `json:"Id" bson:"_id"`
	UserName     string  `json:"username" bson:"username"`
	ProfileImage string  `json:"profileImage" bson:"profileImage"`
	Phone        Phone   `json:"phone" bson:"phone"`
	Country      Country `json:"country" bson:"countryId"`
}
