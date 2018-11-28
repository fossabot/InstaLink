package main

type Configuration struct {
	Server struct {
		host string `json:"host"`
		port string `json:"port"`
	}
	Datanbase struct {
		host     string `json:"host"`
		port     string `json:"port"`
		user     string `json:"user"`
		database string `json:"database"`
		password string `json:"password"`
	}
}

func configure() {

}
