# Distributed-Banking-System

**gRPC Written Report**

**Problem Statement**

Design and implement a distributed banking application usiing gRPC as
the communication mechanism.

**Goal**

The goal is to implement a distributed banking application which uses
gRPC for the communication mechanism and it ensures that the replica of
bank’s balance is maintained accross all its branches so that the
availability is increased.

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

    **Eg usage:**
    ```
    go run start_branch_servers.go inputData.json
    ```

    In the second terminal tab, run the customers
  ```
    cd customer_service

    go run customer_service.go ../inputData.json
```
**Results**

All the events of the customer will be processed and the ‘**Withdraw’**
Operation will withdraw the amount from one branch and then this branch
propgates this transaction to all its peer branches using
Propagate_Withdraw method, similarly the ‘**Deposit’** operation is
carried out. The ‘**query’** type returns the current balance present in
that branch.
