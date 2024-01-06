package api

const baseUrl = "/api/v1"

func (a *api) routes() {
	a.router.HandleFunc(baseUrl+"/info", contentTypeApplicationJsonMiddleware(a.handleInfo()))
}
