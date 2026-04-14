package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/okunix/stash-sdk/stash/v1"
)

func getStashClient(ctx context.Context) (*stash.Client, error) {
	val := ctx.Value("stash-client")
	client, ok := val.(*stash.Client)
	if !ok {
		return nil, errors.New("no client object supplied")
	}
	return client, nil
}

func mustGetStashClient(ctx context.Context) *stash.Client {
	client, err := getStashClient(ctx)
	if err != nil {
		errExit(err)
	}
	return client
}

func errExit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	os.Exit(1)
}
