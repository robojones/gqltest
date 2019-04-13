package request

type ErrorLocation struct {
	Row    int
	Column int
}

type ErrorExtensions = map[string]interface{}

type Error struct {
	Message    string
	Locations  []*ErrorLocation
	Extensions ErrorExtensions
}

type Result struct {
	Data   interface{}
	Errors []*Error
}
