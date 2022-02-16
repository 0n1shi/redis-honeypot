package redis

type Repository interface {
	Save(*Command) error
}
