package database

import (
	"saya/config"
	"saya/models"
)

func GetEventsController() (interface{}, error) {
	var event []models.Event

	if err := config.DB.Find(&event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func GetEventController(EventID uint) (interface{}, error) {
	var event models.Event
	event.ID = EventID

	if err := config.DB.First(&event).Error; err != nil {
		return nil, err
	}

	return event, nil
}

func CreateEvent(b models.Event) (interface{}, error) {
	if err := config.DB.Create(&b).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Joins("Event").Find(&b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func UpdateEvent(eventID uint, b models.Event) (interface{}, error) {
	event := models.Event{}
	event.ID = eventID
	if err := config.DB.Joins("Event").Find(&event).Error; err != nil {
		return nil, err
	}

	event.EventID = b.EventID
	event.Name = b.Name
	event.Kuota = b.Kuota
	event.Harga = b.Harga

	if err := config.DB.Save(&event).Error; err != nil {
		return nil, err
	}

	return event, nil
}

func DeleteEvent(eventID int) (interface{}, error) {
	err := config.DB.Delete(&models.Event{}, eventID).Error

	if err != nil {
		return nil, err
	}
	return eventID, nil
}
