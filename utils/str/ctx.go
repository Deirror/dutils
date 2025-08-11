package str

import (
	"context"
	"errors"
	"fmt"
)

func CtxVal(ctx context.Context, key string) (string, error) {
	val := ctx.Value(key)

	valStr, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("cannot convert to string: %v", val)
	}

	if valStr == "" {
		return "", errors.New("val is empty")
	}

	return valStr, nil
}
