package category

func (m Morphism) Compose2(middleware Morphism) Morphism {
	return func(ctx interface{}) interface{} {
		return m(middleware(ctx));
	}
}

func (m Morphism) Compose(middlewares ...Morphism) Morphism {
	return func(ctx interface{}) interface{} {
		res := m;
		for i := 0; i < len(middlewares); i++ {
			res = res.Compose2(middlewares[i])
		}
		return res(ctx);
	}
}
