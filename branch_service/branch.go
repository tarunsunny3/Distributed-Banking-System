package branch_service

/*

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"


protoc \
--go_out=branch \
--go_opt=paths=source_relative \
--go-grpc_out=branch \
--go-grpc_opt=paths=source_relative \
branch.proto
*/

import (
	"branch_service/branch"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type BranchServer struct {
	branch.UnimplementedBranchServiceServer
	ID                  int32
	Balance             float32 // Balance property for the branch server
	port                int32
	peers               map[int32]branch.BranchServiceClient
	writeEventsReceived map[int32]bool
}

func NewBranchServer(id int32, balance float32, port int32) *BranchServer {
	return &BranchServer{
		ID:                  id,
		Balance:             balance,
		port:                port,
		peers:               make(map[int32]branch.BranchServiceClient),
		writeEventsReceived: make(map[int32]bool),
	}
}

func (s *BranchServer) AddEventID(id int32) {
	s.writeEventsReceived[id] = true
}

func (s *BranchServer) IsEventIDExists(id int32) bool {
	_, exists := s.writeEventsReceived[id]
	return exists
}
func (s *BranchServer) QueryBalance(ctx context.Context, request *branch.QueryBalanceRequest) (*branch.QueryBalanceResponse, error) {

	var lastWriteEventID int32 = request.LastWriteEventID

	if lastWriteEventID == -1 {
		// Return the current balance, since it's first query
		// represented by EventID = -1
		return &branch.QueryBalanceResponse{
			Balance: s.Balance,
		}, nil
	}
	for !s.IsEventIDExists(lastWriteEventID) {
		time.Sleep(100 * time.Millisecond) // Wait for a short duration
	}

	// Return the current balance
	return &branch.QueryBalanceResponse{
		Balance: s.Balance,
	}, nil

}

func (s *BranchServer) Deposit(ctx context.Context, request *branch.DepositRequest) (*branch.DepositResponse, error) {

	// Add the deposited amount to the balance
	s.Balance += request.Amount
	s.AddEventID(request.WriteEventID)
	for peerID, client := range s.peers {
		response, err := client.PropagateDeposit(context.Background(), &branch.PropagateDepositRequest{
			Balance:      request.Amount,
			WriteEventID: request.WriteEventID,
		})
		if err != nil || !response.Success {
			log.Printf("Failed to propagate withdrawal to peer %d: %v", peerID, err)
		}
	}
	// Return the updated balance
	return &branch.DepositResponse{
		NewBalance: s.Balance,
	}, nil
}

func (s *BranchServer) Withdraw(ctx context.Context, request *branch.WithdrawRequest) (*branch.WithdrawResponse, error) {

	// Check if there's enough balance to withdraw
	if s.Balance < request.Amount {
		return nil, fmt.Errorf("insufficient balance")
	}

	// Deduct the amount from the balance
	s.Balance -= request.Amount
	s.AddEventID(request.WriteEventID)
	for peerID, client := range s.peers {
		response, err := client.PropagateWithdraw(context.Background(), &branch.PropagateWithdrawRequest{
			Balance:      request.Amount,
			WriteEventID: request.WriteEventID,
		})
		if err != nil || !response.Success {
			log.Printf("Failed to propagate withdrawal to peer %d: %v", peerID, err)
		}
	}
	return &branch.WithdrawResponse{
		NewBalance: s.Balance,
	}, nil
}

func (s *BranchServer) StartBranchServer() {
	go func() {
		listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		server := grpc.NewServer()
		branch.RegisterBranchServiceServer(server, s)

		// log.Printf("Branch server is running on port %d...\n", s.port)
		if err := server.Serve(listen); err != nil {
			log.Fatalf("Failed to serve branch server: %v", err)
		}
	}()
}

// RegisterPeer registers a peer's gRPC client connection.
func (s *BranchServer) RegisterPeer(peerID int32, client branch.BranchServiceClient) {

	// Store the peer client in the peers map.
	s.peers[peerID] = client
}

// UpdateBalance updates the balance of a specific branch in the data map.
func (s *BranchServer) PropagateWithdraw(ctx context.Context, request *branch.PropagateWithdrawRequest) (*branch.PropagateWithdrawResponse, error) {

	s.Balance -= request.Balance
	s.AddEventID(request.WriteEventID)
	return &branch.PropagateWithdrawResponse{
		Success: true,
	}, nil
}

// UpdateBalance updates the balance of a specific branch in the data map.
func (s *BranchServer) PropagateDeposit(ctx context.Context, request *branch.PropagateDepositRequest) (*branch.PropagateDepositResponse, error) {

	s.Balance += request.Balance
	s.AddEventID(request.WriteEventID)
	return &branch.PropagateDepositResponse{
		Success: true,
	}, nil
}
