package server

import (
	"../client"

	"log"
)
var rooms roomList

func init() {
	rooms = roomList{rooms: map[string]*Room{}}
}

type Room struct {
	id       string
	password string
	hub      *Hub
}

type roomList struct {
	rooms map[string]*Room
}

func getRoom(id string, password string) (room *Room, auth bool) {
	room, exist := rooms.rooms[id]
	if exist {
		if password == room.password {
			return room, true
		}
		return room, false
	}

	hub := newHub()
	go hub.run()
	room = &Room{id: id, password: password, hub: hub}
	rooms.rooms[id] = room
	return room, true
}

func removeRoom(id string) (exist bool) {
	room, exist := rooms.rooms[id]
	if !exist {
		return exist
	}
	log.Println("delete room :", room.id)
	room = nil
	delete(rooms.rooms, id)
	return exist
}

func (room *Room) broadcast(message []byte) {
	room.hub.broadcast <- message
}

func (room *Room) register(client *client.Client) {
	room.hub.register <- client
}

func (room *Room) unregister(client *client.Client) {
	room.hub.unregister <- client
	if len(room.hub.clients) == 0 {
		removeRoom(room.id)
	}
}