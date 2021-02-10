package rest_api_gin

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type NoteList struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type NoteItem struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Delete  bool   `json:"delete"`
}

type ListItem struct {
	Id     int
	ListId int
	UserId int
}
