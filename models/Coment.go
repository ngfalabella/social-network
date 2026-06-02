package models

type Coment struct {
	ID        int    `json:"id"`
	Texto     string `json:"texto"`
	IPUsuario string `json:"-"`
}
