package main

import (
	"gohttp101/data"
	"gohttp101/server"
)

func main() {

	db := data.NewSliceDatabase()
	svr := server.GetServer(db)
	svr.RegisterRoute(server.Get)
	svr.RegisterRoute(server.Post)
	svr.RegisterRoute(server.Update)
	//svr.RegisterRoute(server.Post)
	svr.Start()
}
