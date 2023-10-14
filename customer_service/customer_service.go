package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"branch_service/branch"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Customer struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	Events []struct {
		ID        int    `json:"id"`
		Interface string `json:"interface"`
		Money     int    `json:"money,omitempty"`
	} `json:"events"`
}

type OutputEvent struct {
	Interface string `json:"interface"`
	Result    string `json:"result"`
	Money     int    `json:"money,omitempty"`
}

type OutputData struct {
	ID   int           `json:"id"`
	Recv []OutputEvent `json:"recv"`
}

func main() {
	// Read customer data from JSON file
	if len(os.Args) < 2 {
		fmt.Println("Usage: programName filename")
		return
	}
	inputFilename := os.Args[1]
	customerData, err := readCustomerDataFromFile(inputFilename)
	if err != nil {
		log.Fatalf("Error reading customer data from file %s : %v", inputFilename, err)
	}
	// Create a map to store customer clients
	customerClients := make(map[int]*branch.BranchServiceClient)
	outputFilename := "../output.json"
	// Open the output file in append mode
	outputFile, err := os.OpenFile(outputFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Printf("Error opening output file: %v", err)
		return
	}
	defer outputFile.Close()

	encoder := json.NewEncoder(outputFile)
	outputFile.WriteString("[") // Add the '[' at the beginning

	// Connect to branch servers and establish client connections
	for i, customer := range customerData {

		// Get the customer's ID
		customerID := customer.ID

		// Get the address of the branch server corresponding to the customer's ID
		address := fmt.Sprintf("localhost:%d", 8080+customerID-1)

		// Create a gRPC connection to the branch server
		client, err := createBranchClient(address)
		if err != nil {
			log.Fatalf("Error creating a branch client for customer %d: %v", customerID, err)
		}
		customerClients[customerID] = client

		// Process customer events and collect results
		var results []OutputEvent

		for _, event := range customer.Events {
			result := processCustomerEvent(*client, customerID, event)
			log.Printf("result for customer %d and event id is %d, result %v\n", customer.ID, event.ID, result)
			results = append(results, result)
		}

		// Write the results in the specified format
		outputData := OutputData{
			ID:   customerID,
			Recv: results,
		}

		// Use the JSON encoder to write the outputData to the output file
		if err := encoder.Encode(outputData); err != nil {
			log.Printf("Error encoding and writing output data for customer %d: %v", customerID, err)
			return
		}
		// Add a comma after each object except the last one
		if i < len(customerData)-1 {
			outputFile.WriteString(",")
		}
	}
	outputFile.WriteString("]")
}
func readCustomerDataFromFile(filename string) ([]Customer, error) {
	var data []map[string]interface{}

	// Read JSON file and populate 'data' with customer information
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading customer data file: %v", err)
	}

	err = json.Unmarshal(fileContents, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling customer data: %v", err)
	}

	var customers []Customer

	for _, entry := range data {
		if entryType, ok := entry["type"].(string); ok {
			if entryType == "customer" {
				if id, ok := entry["id"].(float64); ok {
					var events []struct {
						ID        int    `json:"id"`
						Interface string `json:"interface"`
						Money     int    `json:"money,omitempty"`
					}
					if eventsData, ok := entry["events"].([]interface{}); ok {
						for _, eventData := range eventsData {
							event, ok := eventData.(map[string]interface{})
							if ok {
								eventID, _ := event["id"].(float64)
								eventInterface, _ := event["interface"].(string)
								eventMoney, _ := event["money"].(float64)
								events = append(events, struct {
									ID        int    `json:"id"`
									Interface string `json:"interface"`
									Money     int    `json:"money,omitempty"`
								}{
									ID:        int(eventID),
									Interface: eventInterface,
									Money:     int(eventMoney),
								})
							}
						}
					}

					customer := Customer{
						ID:     int(id),
						Type:   entryType,
						Events: events,
					}
					customers = append(customers, customer)
				}
			}
		}
	}

	return customers, nil
}

func createBranchClient(address string) (*branch.BranchServiceClient, error) {
	// Create a gRPC connection to the branch server
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to branch server: %v", err)
	}

	// Create a branch client
	client := branch.NewBranchServiceClient(conn)
	return &client, nil
}

func processCustomerEvent(client branch.BranchServiceClient, customerID int, event struct {
	ID        int    `json:"id"`
	Interface string `json:"interface"`
	Money     int    `json:"money,omitempty"`
}) OutputEvent {
	switch event.Interface {
	case "query":
		// Process query event
		queryResponse, err := client.QueryBalance(context.Background(), &branch.QueryBalanceRequest{})
		if err != nil {
			log.Printf("Error querying balance for customer %d: %v", customerID, err)
			return OutputEvent{Interface: "query", Result: "error"}
		}
		return OutputEvent{Interface: "query", Result: "success", Money: int(queryResponse.Balance)}

	case "deposit":
		// Process deposit event
		_, err := client.Deposit(context.Background(), &branch.DepositRequest{Amount: float32(event.Money)})
		if err != nil {
			log.Printf("Error depositing money for customer %d: %v", customerID, err)
			return OutputEvent{Interface: "deposit", Result: "error"}
		}
		return OutputEvent{Interface: "deposit", Result: "success"}

	case "withdraw":
		// Process withdraw event
		_, err := client.Withdraw(context.Background(), &branch.WithdrawRequest{Amount: float32(event.Money)})
		if err != nil {
			log.Printf("Error withdrawing money for customer %d: %v", customerID, err)
			return OutputEvent{Interface: "withdraw", Result: "error"}
		}
		return OutputEvent{Interface: "withdraw", Result: "success"}
	}

	log.Printf("Unknown event type for customer ID %d: %s\n", customerID, event.Interface)
	return OutputEvent{} // Default empty result
}
