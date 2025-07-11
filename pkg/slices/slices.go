package slices

func Mapper[T any, R any](dbModels []T, mapper func(T) R) []R {
	models := []R{}

	for _, dbModel := range dbModels {
		models = append(models, mapper(dbModel))
	}

	return models
}
