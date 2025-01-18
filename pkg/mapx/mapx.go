package mapx

import "reflect"

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func HasKeys[K comparable, V any](m map[K]V, ks ...K) bool {
	for _, k := range ks {
		if _, ok := m[k]; !ok {
			return false
		}
	}
	return true
}

func MapByFunc[T any, K comparable, V any](s []T, fun func(item T) (K, V)) map[K]V {
	m := make(map[K]V, len(s))
	for i := range s {
		k, v := fun(s[i])
		m[k] = v
	}
	return m
}

func isZeroValue(v reflect.Value) bool {
	// 针对 bool 类型，false 是有效的，不应视为零值
	if v.Kind() == reflect.Bool {
		return false
	}
	// 对其他类型进行零值判断
	return v.IsZero()
}

func MapUpdatesFromObject(obj any) map[string]interface{} {
	updates := make(map[string]interface{})

	// 获取 struct 的值和字段
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	// 遍历所有字段
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name

		// 如果字段值不为空且非零值，则更新该字段
		if field.IsValid() && !isZeroValue(field) {
			updates[fieldName] = field.Interface()
		}
	}
	return updates
}
