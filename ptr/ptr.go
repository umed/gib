package ptr

func Of[T any](v T) *T {
	return &v
}

func ValueOr[T any](p *T, value T) T {
	if p == nil {
		return value
	}
	return *p
}

func Value[T any](p *T) T {
	if p == nil {
		var result T
		return result
	}
	return *p
}
