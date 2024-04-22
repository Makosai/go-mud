package main

// An array of clients with a size of uint8
var clients = [256]*Client{}

// Find a client by ID.
func FindClient(id uint8) *Client {
	return clients[id]
}
