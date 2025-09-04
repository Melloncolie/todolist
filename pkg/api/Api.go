package api

func (api *Api) New() error {
	return api.todoStorage.Import()
}
