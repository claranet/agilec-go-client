package models

type Model interface {
	ToJson() ([]byte, error)
}
