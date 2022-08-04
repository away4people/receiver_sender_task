package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"math/rand"
	"time"
)

type message struct {
	time time.Time
	text string
}

func sender(ctx context.Context, ch chan message, duration time.Duration) {
	ticker := time.Tick(duration)
	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		case <-ticker:
			ch <- message{time: time.Now(), text: randString(rand.Intn(32))}
		}
	}
}

func receiver(ctx context.Context, ch chan message) {
	ticker := time.Tick(5 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker:
			message, ok := <-ch
			if ok {
				fmt.Printf("Receiver got message. Time: %v, text: %s \n", message.time, message.text)
			}
		}
	}
}

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("config loading failed")
	}
	sendInterval := viper.GetInt("SEND_MESSAGE_INTERVAL")

	ch := make(chan message)
	waitCh := make(chan struct{})
	defer close(waitCh)
	ctx, cancel := context.WithCancel(context.Background())

	go sender(ctx, ch, time.Second*time.Duration(sendInterval))
	go receiver(ctx, ch)
	go func() {
		select {
		case <-time.After(60 * time.Second):
			cancel()
			waitCh <- struct{}{}
		}
	}()
	<-waitCh
}

func randString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
