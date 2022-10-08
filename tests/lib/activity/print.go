package activity

import (
	"context"
	"fmt"
)

func LogString(ctx context.Context, log string) error {
 fmt.Println(log)
 return nil
}
