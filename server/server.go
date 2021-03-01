package server

import "github.com/laybatin/bodymirror-app-backend/config"

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("server.port"))
}
