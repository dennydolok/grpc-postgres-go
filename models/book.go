package models

type Book struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	OwnerId   int       `json:"owner_id"`
	Author    string    `json:"author"`
	Desc      string    `json:"desc"`
	BookOwner BookOwner `json:"book_owner" gorm:"foreignKey:owner_id"`
}

type BookOwner struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
