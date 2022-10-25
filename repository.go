package redishoneypot

type Repository interface {
	Save(*Command) error
}
