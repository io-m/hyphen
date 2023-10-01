package models

type CustomerFollowsBusiness struct {
	ProfileId  int `json:"profile_id,omitempty" db:"profile_id"`
	BusinessId int `json:"business_id,omitempty" db:"business_id"`
}
