package dtos

type PlayerListQuery struct {
	PageNo  int    `query:"page" validate:"omitempty,min=1"`               // Optional, must be >= 1 if provided
	Status  string `query:"status" validate:"omitempty,oneof=true false"`  // Optional, must be a valid email if provided
	Gender  string `query:"gender" validate:"omitempty,oneof=male female"` // Optional, must be "male" or "female"
	Country int    `query:"country" validate:"omitempty,min=1"`            // Optional, must be >= 1 if provided
}
