package models

type Playlist struct {
	Id int
	CreateDate string
	Title string
	User User
	Artists []Artist

}

type Artist struct {
	Id int
	Name string
	Songs []Song
}

type Song struct {
	Id int
	CreateDate string
}