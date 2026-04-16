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
	fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
	os.Exit(1)
}

func printStashResponse(resp *stash.StashResponse) error {
	fmt.Printf("ID:\t%s\n", resp.ID)
	fmt.Printf("Name:\t%s\n", resp.Name)
	fmt.Printf("MaintainerID:\t%s\n", resp.MaintainerID)
	fmt.Printf("CreatedAt:\t%s\n", resp.CreatedAt)
	fmt.Printf("Locked:\t%v\n", resp.Locked)
	if resp.Description != nil {
		fmt.Printf("Description:\t%s\n", *resp.Description)
	}
	return nil
}

func printStashResponseList(resp []stash.StashResponse) error {
	for _, v := range resp {
		printStashResponse(&v)
		fmt.Println()
	}
	return nil
}

func printUserResponse(resp *stash.UserResponse) error {
	fmt.Printf("ID:\t%s\n", resp.ID)
	fmt.Printf("Username:\t%s\n", resp.Username)
	fmt.Printf("Locked:\t%v\n", resp.Locked)
	fmt.Printf("CreatedAt:\t%s\n", resp.CreatedAt)
	if resp.ExpiredAt != nil {
		fmt.Printf("ExpiredAt:\t%s\n", resp.CreatedAt)
	}
	return nil
}

func printUserResponseList(resp []*stash.UserResponse) error {
	for _, v := range resp {
		printUserResponse(v)
		fmt.Println()
	}
	return nil
}
