package main

import (
	// Update import path

	"branch_service"
	"branch_service/branch"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func readBranchDataFromFile(filename string) ([]*branch.Branch, error) {
	var data []map[string]interface{}

	// Read JSON file and populate 'data' with branch information
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading branch data file: %v", err)
	}

	err = json.Unmarshal(fileContents, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling branch data: %v", err)
	}

	var branches []*branch.Branch

	for _, entry := range data {
		if entryType, ok := entry["type"].(string); ok {
			if entryType == "branch" {
				if balance, ok := entry["balance"].(float64); ok {
					id, _ := entry["id"].(float64)
					branch := &branch.Branch{
						Id:      int32(id),
						Balance: float32(balance),
					}
					branches = append(branches, branch)
				}
			}
		}
	}

	return branches, nil
}

func createBranchClient(address string) (branch.BranchServiceClient, error) {
	// Create a gRPC connection to the branch server
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to branch server: %v", err)
	}

	// Create a branch client
	client := branch.NewBranchServiceClient(conn)
	return client, nil
}

func main() {
	// Read branch data from JSON file
	log.SetOutput(os.Stdout)
	if len(os.Args) < 2 {
		fmt.Println("Usage: programName filename")
		return
	}
	inputFilename := os.Args[1]
	branchData, err := readBranchDataFromFile(inputFilename)
	fmt.Print(branchData)
	if err != nil {
		log.Fatalf("Error reading branch data: %v", err)
	}
	// Create a map to store branch servers and their clients
	branchServers := make(map[int32]*branch_service.BranchServer)
	branchClients := make(map[int32]branch.BranchServiceClient)
	// Use a wait group to ensure all servers and clients are initialized
	var wg sync.WaitGroup

	for _, data := range branchData {
		wg.Add(1) // Increment the wait group counter
		port := 8080 + data.Id - 1
		// Start the branch server
		server := branch_service.NewBranchServer(data.Id, data.Balance, port)
		go func(data *branch.Branch, server *branch_service.BranchServer, port int32) {
			defer wg.Done() // Decrement the wait group counter when done
			fmt.Printf("Starting branch server for ID: %d, Initial Balance: %.2f on port: %d\n", data.Id, data.Balance, port)
			server.StartBranchServer()
			if err != nil {
				log.Printf("Error starting branch server: %v", err)
			}
		}(data, server, port)

		// Register the branch server
		branchServers[data.Id] = server

		// Create a client for the branch
		client, err := createBranchClient(fmt.Sprintf("localhost:%d", port))
		if err != nil {
			log.Fatalf("Error creating a branch client for the  branch: %v", err)
		}
		branchClients[data.Id] = client
		// Increment the port for the next branch server
		port++
	}
	// Wait for all branch servers and clients to be initialized
	wg.Wait()

	// Register peers and establish connections between branches
	for id, server := range branchServers {
		for peerID, client := range branchClients {
			if id != peerID {
				server.RegisterPeer(peerID, client)
			}
		}
	}

	// Block to keep the servers running
	select {}
}
