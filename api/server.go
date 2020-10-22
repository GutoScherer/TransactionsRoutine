package api

// Server represents any type of server expose the API
type Server interface {
	ListenAndServe(port int)
}
