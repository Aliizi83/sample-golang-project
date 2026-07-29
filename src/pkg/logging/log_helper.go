package logging

func ConvertMapToInterface(extra map[ExtraKey]interface{}) []interface{} {
	var pairs []interface{}
	for key, value := range extra {
		pairs = append(pairs, string(key), value)
	}
	return pairs
}
