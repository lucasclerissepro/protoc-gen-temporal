package activity

import (
	"context"
	"fmt"
	"time"
)

func LogString(ctx context.Context, log string) error {
	time.Sleep(time.Second * 10)
	fmt.Println(log)
	return nil
}
