package models

import "time"

type Playlist struct {
	Id         int
	CreateDate time.Time
	Title      string
	UserId     int
	Songs      []Song
	ImagePath string
}

type Artist struct {
	Id    int
	Name  string
	Songs []Song
}

type Song struct {
	Id         int
	Title   string
	CreateDate time.Time
	Artist Artist
}

type PlaylistSongs struct {
	PlaylistId int
	Songs []Song
}