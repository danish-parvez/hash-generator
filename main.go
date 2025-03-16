package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Function to generate a random alphanumeric string
func generateRandomHash(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"        // Characters to use for the hash
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano())) // Seed the random number generator

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))] // Pick a random character from the charset
	}
	return string(b)
}

// Function to print the hash and schedule the next call
func getHashNow() {
	randomHash := generateRandomHash(6) // Generate a 6-character random hash
	fmt.Println(randomHash)             // Print the hash to the console

	// Schedule the next call after 5 seconds
	time.AfterFunc(5*time.Second, getHashNow)
}

func main() {
	getHashNow() // Start the process

	// Keep the program running indefinitely
	select {}
}
