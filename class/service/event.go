package service

import (
	"20241114/class/model"
	"20241114/class/repository"
)

type EventService struct {
	EventRepo repository.Event
}

func InitEventService(repo repository.Event) *EventService {
	return &EventService{EventRepo: repo}
}

func (eventService EventService) All(date string, page int, sort string) (int, int, []model.Event, error) {
	return eventService.EventRepo.All(date, page, sort)
}

func (eventService EventService) Get(event *model.Event) error {
	return eventService.EventRepo.Get(event)
}
