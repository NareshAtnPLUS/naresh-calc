package models

type Calculation struct {
	X         int    `bson:"x,omitempty"`
	Y         int    `bson:"y,omitempty"`
	Operation string `bson:"operation,omitempty"`
	Result    int    `bson:"result,omitempty"`
}
