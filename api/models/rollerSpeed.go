package models

import (
	"crudAPIs/api/constants"
	"errors"
	"fmt"

	"github.com/google/uuid"
	gorm "github.com/jinzhu/gorm"
)

type RollerSpeed7 struct {
	UUID            uuid.UUID `gorm:"type:uuid;not null" json:"uuid"`
	EventTimestamp  int64     `gorm:"not null; primary key" json:"event_timestamp"`
	RollerTimestamp string    `gorm:"not null" json:"roller_timestamp"`
	Roller1Speed    float64   `gorm:"column:roller_1_speed" json:"roller_1_speed"`
	Roller2Speed    float64   `gorm:"column:roller_2_speed" json:"roller_2_speed"`
	Roller3Speed    float64   `gorm:"column:roller_3_speed" json:"roller_3_speed"`
	Roller4Speed    float64   `gorm:"column:roller_4_speed" json:"roller_4_speed"`
	Roller5Speed    float64   `gorm:"column:roller_5_speed" json:"roller_5_speed"`
	Roller6Speed    float64   `gorm:"column:roller_6_speed" json:"roller_6_speed"`
	Roller7Speed    float64   `gorm:"column:roller_7_speed" json:"roller_7_speed"`
	Roller8Speed    float64   `gorm:"column:roller_8_speed" json:"roller_8_speed"`
	Roller9Speed    float64   `gorm:"column:roller_9_speed" json:"roller_9_speed"`
	Roller10Speed   float64   `gorm:"column:roller_10_speed" json:"roller_10_speed"`
	Roller11Speed   float64   `gorm:"column:roller_11_speed" json:"roller_11_speed"`
	Roller12Speed   float64   `gorm:"column:roller_12_speed" json:"roller_12_speed"`
	Roller13Speed   float64   `gorm:"column:roller_13_speed" json:"roller_13_speed"`
	Roller14Speed   float64   `gorm:"column:roller_14_speed" json:"roller_14_speed"`
	RecipeID        string    `gorm:"not null" json:"recipe_id"`
	EquipmentID     string    `gorm:"not null" json:"equipment_id"`
	MotherRollID    string    `gorm:"not null" json:"mother_roll_id"`
}

func (r *RollerSpeed7) TableName() string {
	return "roller_speed_7"
}

func (r RollerSpeed7) FindRollerDataByID(db *gorm.DB, uuid uuid.UUID) (RollerSpeed7, error) {
	fmt.Println("Entered function FindRollerDataByID")

	err := db.Debug().Model(RollerSpeed7{}).Where("uuid = ?", uuid).Take(&r).Error
	//to select particular columns specify it into .Select(string..... column_name) before .Where in above line of code
	if err != nil {
		fmt.Printf("error here is %v\n", err)
		return RollerSpeed7{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return RollerSpeed7{}, errors.New("Roller Speed data Not Found")
	}

	return r, err
}

func (r RollerSpeed7) FindRollerDataByRollerData(db *gorm.DB, rollerSpeed7 RollerSpeed7) (RollerSpeed7, error) {
	fmt.Println("Entered the function FindRollerDataByRollerData")

	err := db.Debug().Model(RollerSpeed7{}).Where("event_timestamp = ? AND roller_timestamp = ? AND roller_1_speed = ? AND roller_2_speed = ? AND roller_3_speed = ? AND roller_4_speed = ? AND roller_5_speed = ? AND roller_6_speed = ? AND roller_7_speed = ? AND roller_8_speed = ? AND roller_9_speed = ? AND roller_10_speed = ? AND roller_11_speed = ? AND roller_12_speed = ? AND roller_13_speed = ? AND roller_14_speed = ? AND recipe_id = ? AND equipment_id = ? AND mother_roll_id = ?",
		rollerSpeed7.EventTimestamp, rollerSpeed7.RollerTimestamp, rollerSpeed7.Roller1Speed, rollerSpeed7.Roller2Speed, rollerSpeed7.Roller3Speed, rollerSpeed7.Roller4Speed, rollerSpeed7.Roller5Speed, rollerSpeed7.Roller6Speed, rollerSpeed7.Roller7Speed, rollerSpeed7.Roller8Speed, rollerSpeed7.Roller9Speed, rollerSpeed7.Roller10Speed, rollerSpeed7.Roller11Speed, rollerSpeed7.Roller12Speed, rollerSpeed7.Roller13Speed, rollerSpeed7.Roller14Speed, rollerSpeed7.RecipeID, rollerSpeed7.EquipmentID, rollerSpeed7.MotherRollID).Take(&r).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return RollerSpeed7{}, nil
		}
		fmt.Printf("error here is %v\n", err)
		return RollerSpeed7{}, err
	}

	return r, nil
}

func (r RollerSpeed7) CreateRollerData(db *gorm.DB, rollerSpeed7 RollerSpeed7) (RollerSpeed7, error) {
	fmt.Println("Entered the function CreateRollerData")
	uuid := uuid.New()

	existingRollerData, err := r.FindRollerDataByRollerData(db, rollerSpeed7)

	if err != nil {
		return RollerSpeed7{}, err
	}
	if (existingRollerData != RollerSpeed7{}) {
		fmt.Printf("Duplicate data \n")
		return existingRollerData, constants.ErrDuplicateData
	}

	fmt.Printf("RollerSpeed data passed to insert is %v\n", rollerSpeed7)
	rollerSpeed7.UUID = uuid

	if err := db.Debug().Model(&RollerSpeed7{}).Create(&rollerSpeed7).Error; err != nil {
		fmt.Printf("Unable to create a record in dB")
		return RollerSpeed7{}, err
	}

	return rollerSpeed7, nil
}

func (r RollerSpeed7) UpdateRollerSpeedDataByUUID(db *gorm.DB, uuid uuid.UUID, roller_1_speed float64) (RollerSpeed7, error) {
	fmt.Println("Entered the function UpdateRollerDataByUUID")

	existingRollerData, err := r.FindRollerDataByID(db, uuid)

	if err != nil {
		return RollerSpeed7{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return RollerSpeed7{}, constants.ErrDuplicateData
	}

	fmt.Printf("\n found record %v\n", existingRollerData)

	err = db.Debug().Model(&RollerSpeed7{}).Where("uuid = ?", existingRollerData.UUID).Updates(map[string]interface{}{"roller_1_speed": roller_1_speed}).Error
	if err != nil {
		return RollerSpeed7{}, err
	}

	return existingRollerData, nil

}

func (r RollerSpeed7) DeleteRollerDataByUUID(db *gorm.DB, uuid uuid.UUID) error {
	fmt.Println("Entered into function DeleteRollerDataByUUID")

	existingRollerData, err := r.FindRollerDataByID(db, uuid)

	if err != nil {
		return err
	}

	if gorm.IsRecordNotFoundError(err) {
		return constants.ErrDataNotFound
	}

	fmt.Printf("\n found record %v\n", existingRollerData)

	err = db.Debug().Delete(&RollerSpeed7{}, "uuid = ?", existingRollerData.UUID).Error
	if err != nil {
		return err
	}

	return nil
}
