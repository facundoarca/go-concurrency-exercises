package main

import (
	"context"
	"fmt"
)

type database map[string]bool

var db database = database{
	"jane": true,
}

type uuid string

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	processRequest(ctx, "jane")
}

func processRequest(ctx context.Context, userid string) {
	// TODO: send userID information to checkMemberShip through context for
	// map lookup.
	ctxWithValue := context.WithValue(ctx, uuid("123456"), userid)
	ch := checkMemberShip(ctxWithValue)
	status := <-ch
	fmt.Printf("membership status of userid : %s : %v\n", userid, status)
}

// checkMemberShip - takes context as input.
// extracts the user id information from context.
// spins a goroutine to do map lookup
// sends the result on the returned channel.
func checkMemberShip(ctx context.Context) <-chan bool {
	ch := make(chan bool)
	go func() {
		defer close(ch)
		// do some database lookup
		userId := ctx.Value(uuid("123456")).(string)
		status := db[userId]
		ch <- status
	}()
	return ch
}
