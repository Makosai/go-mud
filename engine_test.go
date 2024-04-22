package main

import (
	"testing"
)

var name string = "Mako"

// NewClient should create a local client with an ID of 1.
func TestLocalClient(t *testing.T) {
	client := NewClient(name)
	if client.id == 1 {
		t.Log("Client ID is 1")
	} else {
		t.Error("Client ID is not 1")
	}
}

// NewClientRemote should set the ID properly.
func TestRemoteClient(t *testing.T) {
	client := NewClientRemote(2, name)
	if client.id == 2 {
		t.Log("Client ID is 2")
	} else {
		t.Error("Client ID is not 2")
	}
}

// NewClientRemote should panic if the ID is 1.
func TestRemoteClientPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		} else {
			t.Log("The code panicked")
		}
	}()
	NewClientRemote(1, name)
}
