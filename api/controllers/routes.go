package controllers

import (
	"crudAPIs/api/constants"
	"crudAPIs/api/middlewares"
)

func (s *Server) initializeRoutes() {

	//roller speed data routes
	s.Router.HandleFunc("/rollerSpeed", middlewares.SetMiddlewareJSON(s.GetRollerDataByID)).Methods("GET").Headers(constants.Uuid, "")
	s.Router.HandleFunc("/rollerSpeed/create", middlewares.SetMiddlewareJSON(s.CreateRollerData)).Methods("POST")
	s.Router.HandleFunc("/rollerSpeed/update", middlewares.SetMiddlewareJSON(s.UpdateRollerSpeedData)).Methods("PUT").Headers(constants.Uuid, "").Headers(constants.Roller1Speed, "")
	s.Router.HandleFunc("/rollerSpeed/delete", middlewares.SetMiddlewareJSON(s.DeleteRollerSpeedData)).Methods("PUT").Headers(constants.Uuid, "")
}
