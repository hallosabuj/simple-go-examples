package calculator

type Numbers struct {
	First  int16 `json:"num1" example:"Pepsi Inc."`
	Second int16 `json:"num2" example:"Pepsi Inc."`
}
type Result struct {
	First     int16  `json:"num1" example:"Pepsi Inc."`
	Second    int16  `json:"num2" example:"Pepsi Inc."`
	Operation string `json:"operation" example:"Pepsi Inc."`
	Result    int32  `json:"result" example:"Pepsi Inc."`
}
