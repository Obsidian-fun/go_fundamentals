
package main

import (
	"fmt"
	"context"
)

type keyType struct{}

var key = &keyType{};

func WithKey(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, key, value);
}

func GetKeyWithValue(ctx context.Context) (string, bool) {
	v := ctx.Value(key);
	if v == nil {
		return "", false;
	}

	return v.(string), true;
}


func main() {
	ctx := WithKey(context.Background(), "1234");
	fmt.Println(GetKeyWithValue(ctx));
}


