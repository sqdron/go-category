package theory;

func (m Morphism) Compose2(middleware Morphism) Morphism {
	return func(ctx interface{}) interface{} {
		return m(middleware(ctx));
	}
}

func (m Morphism) Compose(morphisms ...Morphism) Morphism {
	return func(ctx interface{}) interface{} {
		res := m;
		for i := 0; i < len(morphisms); i++ {
			res = res.Compose2(morphisms[i])
		}
		return res(ctx);
	}
}