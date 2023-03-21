package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"crudAPIs/api/constants"
	"crudAPIs/api/models"
	"crudAPIs/api/responses"
	"crudAPIs/api/utils"
)

func (server *Server) GetRollerDataByID(w http.ResponseWriter, r *http.Request) {

	//vars := mux.Vars(r)
	iD := r.Header.Get(constants.Uuid)
	uuid, err := uuid.Parse(iD)
	//fmt.Printf("variables are %v\n", vars)
	fmt.Printf("uuid passed is %v\n", uuid)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	roller := models.RollerSpeed7{}
	rollerSpeed, err := roller.FindRollerDataByID(server.DB, uuid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Printf("response recieved %v\n", rollerSpeed)
	responses.JSON(w, http.StatusOK, rollerSpeed)
}

func (server *Server) CreateRollerData(w http.ResponseWriter, r *http.Request) {

	var rollerSpeed models.RollerSpeed7

	err := json.NewDecoder(r.Body).Decode(&rollerSpeed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	roller := models.RollerSpeed7{}

	fmt.Printf("\n rollerSpeed data recieved in body is %v\n", rollerSpeed)
	rollerData, err := roller.CreateRollerData(server.DB, rollerSpeed)
	if errors.Is(err, constants.ErrDuplicateData) {
		fmt.Println("Duplicate Insertion")
		responses.ERROR(w, http.StatusConflict, err)
		return
	}
	if err != nil {
		fmt.Printf("Unable to create the record %v\n", err)
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	fmt.Printf("Below data is inserted into dB %v\n", rollerData)
	responses.JSON(w, http.StatusOK, rollerData)
}

func (server *Server) UpdateRollerSpeedData(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Inside UpdateRollerSpeedData controller")

	iD := r.Header.Get(constants.Uuid)
	roller1Speed := r.Header.Get(constants.Roller1Speed)

	uuid, err := uuid.Parse(iD)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	roller1Speedf := utils.Convert_string_float(roller1Speed)

	fmt.Printf("UUId passed is : %v and roller 1 speed is %v\n", uuid, roller1Speed)

	var rollerSpeed models.RollerSpeed7

	updatedRollerSpeedData, err := rollerSpeed.UpdateRollerSpeedDataByUUID(server.DB, uuid, roller1Speedf)

	if errors.Is(err, constants.ErrDataNotFound) {
		fmt.Println("Data is not found in dB")
		responses.ERROR(w, http.StatusConflict, err)
		return
	}

	if err != nil {
		fmt.Printf("Unable to update the record %v\n", err)
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	fmt.Printf("Updated data in dB %v\n", updatedRollerSpeedData)
	responses.JSON(w, http.StatusOK, updatedRollerSpeedData)
}

func (server *Server) DeleteRollerSpeedData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside DeleteRollerSpeedData controller")

	iD := r.Header.Get(constants.Uuid)
	uuid, err := uuid.Parse(iD)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	fmt.Printf("UUID passed is %v\n", uuid)

	var rollerSpeed models.RollerSpeed7

	err = rollerSpeed.DeleteRollerDataByUUID(server.DB, uuid)
	if err != nil {
		fmt.Printf("Unable to delete the record %v\n", err)
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, constants.SuccesfullyDeleted)
}
