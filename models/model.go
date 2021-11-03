package models

type Model interface {
	ToMap() ([]byte, error)
}
