package api

const baseURL = "/api/v1"

func (a *API) routes() {
	a.router.HandleFunc(baseURL+"/info", contentTypeApplicationJSONMiddleware(a.handleInfo()))
}
