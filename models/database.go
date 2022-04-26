package models

type StorageData struct {
	User             GophUser
	LoginPasswordMap map[string]string
}
