package model

//Server - Model
type Server struct {
	Address  string `json:"ipAddress"`
	SslGrade string `json:"grade"`
	Country  string `json:"country"`
	Owner    string `json:"owner"`
}
