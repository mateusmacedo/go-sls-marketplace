package adapter

type HttpMethodGuard interface {
	IsMethodAllowed(method string) bool
}

type httpMethodGuard struct {
	allowedMethods []string
}

func NewHttpMethodGuard(allowedMethods []string) HttpMethodGuard {
	return &httpMethodGuard{
		allowedMethods: allowedMethods,
	}
}

func (g *httpMethodGuard) IsMethodAllowed(method string) bool {
	for _, allowedMethod := range g.allowedMethods {
		if method == allowedMethod {
			return true
		}
	}
	return false
}
