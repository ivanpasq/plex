/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/ivanpasq/plex/cmd"

func main() {
	rootCmd := cmd.Root()
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
