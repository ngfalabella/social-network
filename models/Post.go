package models

type Post struct {
	ID        int    `json:"id"`
	Autor     string `json:"autor"`
	Contenido string `json:"contenido"`
	Likes     int    `json:"likes"`
	IpCreador string `json:"-"`
}
