package models

type Paste struct {
	ID     int    `json:"item_id" ,db:"id"`
	Title  string `json:"item_title" ,db:"title"`
	Paste  string `json:"item_paste" ,db:"paste"`
	Author *User  `json:"item_author" ,db:"author"`
}
