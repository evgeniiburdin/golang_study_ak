package main

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"golang.org/x/sync/errgroup"
)

func monitorGoroutines(ctx context.Context, prevGoroutines int) {
	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			currentGoroutines := runtime.NumGoroutine()
			fmt.Printf("Текущее количество горутин: %d\n", currentGoroutines)

			if prevGoroutines == 0 {
				prevGoroutines = currentGoroutines
				continue
			}

			diff := float64(currentGoroutines-prevGoroutines) / float64(prevGoroutines) * 100
			if diff >= 20 {
				fmt.Printf("⚠  Предупреждение: Количество горутин увеличилось более чем на 20%%!\n")
			} else if diff <= -20 {
				fmt.Printf("⚠  Предупреждение: Количество горутин уменьшилось более чем на 20%%!\n")
			}

			prevGoroutines = currentGoroutines
		}
	}
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		monitorGoroutines(ctx, runtime.NumGoroutine())
		return nil
	})

	for i := 0; i < 64; i++ {
		i := i
		g.Go(func() error {
			time.Sleep(5 * time.Second)
			if i == 32 { // условие для демонстрации ошибки
				return fmt.Errorf("ошибка в горутине %d", i)
			}
			return nil
		})
		time.Sleep(80 * time.Millisecond)
	}

	if err := g.Wait(); err != nil {
		fmt.Println("Ошибка:", err)
	}
}
