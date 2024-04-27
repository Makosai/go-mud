package ui

import (
	"embed"
)

//go:embed images/*
var images embed.FS

type Images string

const (
	Merchant Images = "images/merchant.jpg"
)
