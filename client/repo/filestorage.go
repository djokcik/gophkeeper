package repo

type Storage interface {
	Open() error
	Close() error
	ReadData(in interface{}) error
	WriteData(data interface{}) error
}
