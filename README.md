# Distributed-Banking-System

**gRPC Written Report**

**Problem Statement**

Design and implement a distributed banking application usiing gRPC as the communication mechanism implementing a “read-your-writes” consistency mechanism.

**Goal**

The goal is to implement a distributed banking application that uses gRPC for the communication mechanism while ensuring a good level of consistency accross all the replicas. The idea is to implement “read-your-writes” consistency model and ensure consistency.

**Setup**

Languages used are: Golang as the programming language  
gRPC and Protocol Buffers as the communication mechanism

**Setup**:  
Installation and setup of Go on Ubuntu:

**Go version used: 1.21.3**

1.  Go to the link and follow the steps
    [<u>https://go.dev/doc/install</u>](https://go.dev/doc/install) 

> Step 1 is to download the Go package for Linux from
> [<u>https://go.dev/dl/</u>](https://go.dev/dl/) 
```
1.   sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.3.linux-arm64.tar.gz

2.  export PATH=$PATH:/usr/local/go/bin

3.  go version
```
**Implementation Processes**

There are three main code files which cover the entire functionality

1.  **branch.go** it contains all the implementations of the methods
    like Deposit, Withdraw, Propagate_Withdraw, Propagate_Deposit.

2.  **branch.proto** contains the gRPC protobuffer related stuff for the
    services and return types.

3.  **customer_service.go**: It reads the input data json file and
    processes the customer’s events sequentially.

4.  **start_branch_servers.go**: This code spawns the branch servers
    after reading the branches data from the input file.

    First step is to run start_branch_servers.go, it expects the input
    data file name to be the **command line argument**
    
    [Link to input_data.json](input_data.json)
    
    **Eg usage:**
    ```
    go run start_branch_servers.go input_data.json
    ```

    In the second terminal tab, run the customers
  ```
    cd customer_service

    go run customer_service.go ../input_data.json
```

I have implemented the “read-your-writes” consistency model by following the steps below:

1.Event unique token generation:  When the customer initiates an event/transaction processing, this unique “token” which in my implementation is the unique ID of the current event is sent to the Write operation which is “DEPOSIT” or “WITHDRAW” operation of a branch.


2.Token Propagation: The passed token from the customer process to the branch process will be used to send to the respective peer branches while sending the propagation requests. This way each branch process stores the information in a map called writeEventsReceived, and this is used to determine whether to block an incoming read request from the customer.


The following code represents the part where the process blocks the read request from a customer until it knows that the previous write operation has managed to propagate to this current branch.

```

func (s *BranchServer) QueryBalance(ctx context.Context, request *branch.QueryBalanceRequest) (*branch.QueryBalanceResponse, error) {

var lastWriteEventID int32 = request.LastWriteEventID

if lastWriteEventID == -1 {
// Return the current balance, since it's the first query
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

```

