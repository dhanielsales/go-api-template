package logger

func LogString(key, value string) FieldOption {
	return func(f *Field) {
		f.Key = key
		f.Value = value
	}
}

func LogInt(key string, value int) FieldOption {
	return func(f *Field) {
		f.Key = key
		f.Value = value
	}
}

func LogBool(key string, value bool) FieldOption {
	return func(f *Field) {
		f.Key = key
		f.Value = value
	}
}

func LogAny(key string, value any) FieldOption {
	return func(f *Field) {
		f.Key = key
		f.Value = value
	}
}

func LogErr(key string, value error) FieldOption {
	return func(f *Field) {
		f.Key = key
		f.Value = value.Error()
	}
}

func formatFields(fieldOptions []FieldOption) []any {
	res := make([]any, 0)
	for _, fieldOpt := range fieldOptions {
		field := &Field{}
		fieldOpt(field)
		res = append(res, field.Key, field.Value)
	}

	return res
}
