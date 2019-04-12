package query

type Resolver struct{}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (_ *Resolver) Hello() string {
	return "Hello, world!"
}
