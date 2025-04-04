package main

import (
	"context"
	initApp "testTrood/cmd/app/init"
	"testTrood/internal/common/types/user_type"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	init := initApp.NewInit()
	init.MustInit()
	init.GetAskHandler().Handle(ctx, user_type.UserMessage("я хочу изменить пароль"))
}
