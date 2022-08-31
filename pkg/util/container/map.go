package container

type Map interface {
	Get(from string) (string, bool)
	Set(from, to string)
	Remove(from string)
	Range() map[string]string
}
