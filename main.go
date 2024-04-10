package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	bookHotel(ctx)

}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Tempo excedido para bookar o quarto")
	case <-time.After(time.Second * 5):
		fmt.Println("Quarto reservado com sucesso")
	}
}
