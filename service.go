package main

// RoomParams is ...
type RoomParams struct {
	Owner       string
	Description string
	Public      bool
	Expire      float64
}

// Room is ...
type Room struct {
	ID          string
	Owner       string
	Description string
	Public      bool
}

// RoomsService is ...
type RoomsService interface {
	CreateRoom(args *RoomParams) (*Room, error)
	GetRoom(roomID string) (*Room, error)
	DestroyRoom(roomID string) (interface{}, error)
}

type defaultService struct {
	rooms map[string]Room
}

// NewRoomsService is a factory for creating a rooms service.
func NewRoomsService() RoomsService {
	return &defaultService{rooms: make(map[string]Room)}
}

func (s *defaultService) CreateRoom(args *RoomParams) (*Room, error) {
	room := Room{Description: args.Description, ID: "12345", Owner: args.Owner, Public: args.Public}
	s.rooms[room.ID] = room
	return &room, nil
}

func (s *defaultService) GetRoom(roomID string) (*Room, error) {
	room := s.rooms[roomID]
	return &room, nil
}

func (s *defaultService) DestroyRoom(roomID string) (interface{}, error) {
	return nil, nil
}
