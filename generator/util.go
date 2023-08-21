package generator

func intP(i *int32) *int {
	if i == nil {
		return nil
	}

	ii := int(*i)
	return &ii
}
