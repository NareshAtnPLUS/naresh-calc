package types

type CalcParameters struct {
	X int `json:"x" binding:"required"`
	Y int `json:"y" binding:"required"`
}
