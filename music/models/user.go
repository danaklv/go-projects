package models


type User struct {
	Id int
	Name string
	Email string
	Password string
	Playlists []Playlist
}



type ResetPassword struct {
	Email string
	ErrorCode error
	SuccessCode string

}