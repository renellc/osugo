package osugo

type query interface {
	constructQuery(key string) (string, error)
	validateQuery() error
}
