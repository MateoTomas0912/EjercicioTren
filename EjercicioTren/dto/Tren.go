package dto

type Tren struct {
	Id         int
	Nombre     string
	Locomotora Locomotora
	Vagones    []Vagon
}
