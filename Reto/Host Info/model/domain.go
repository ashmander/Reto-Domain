package model

//Domain - Model
type Domain struct {
	Host             string   `json:"host"`
	ServersChange    bool     `json:"servers_change"`
	IsDown           bool     `json:"is_down"`
	SslGrade         string   `json:"ssl_grade"`
	PreviousSslGrade string   `json:"previous_ssl_grade"`
	Logo             string   `json:"logo"`
	Title            string   `json:"title"`
	Servers          []Server `json:"endpoints"`
}
