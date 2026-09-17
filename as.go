package builder

type Aliased struct {
	table any
	alias string
}

func As(table any, alias string) *Aliased {
	return &Aliased{table, alias}
}
