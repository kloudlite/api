package common

import (
	fn "kloudlite.io/pkg/functions"
)

func ResolveLabels(labels map[string]string) (map[string]any, error) {
	m := make(map[string]any, len(labels))
	if labels == nil {
		return nil, nil
	}
	if err := fn.JsonConversion(labels, &m); err != nil {
		return nil, err
	}
	return m, nil
}
