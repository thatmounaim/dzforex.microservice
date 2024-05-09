package storage

type Storage interface {
	UpdateData(map[string]float32)
	GetAll() map[string]float32
	Get(string) (float32, error)
}
