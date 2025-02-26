# Simple Blockchain in Go

This project is a basic blockchain implementation in Go for learning purposes. It demonstrates core blockchain concepts like blocks, chains, and proof-of-work.

## Blockchain Principles Demonstrated

This simple project touches upon several fundamental blockchain ideas:

*   **Blocks:** Data is stored in containers called blocks. Each block in this example holds simple text data.
*   **Chaining:** Blocks are linked together in a chain. Each block contains the hash of the *previous* block, creating a chronological and tamper-evident history.
*   **Hashing:**  Cryptographic hashes (SHA-256 in this case) are used to uniquely identify each block and ensure data integrity. If any data within a block is changed, its hash will change, breaking the chain.
*   **Immutability (Simplified):** Because blocks are chained and rely on hashes, altering a block becomes computationally difficult (due to Proof-of-Work and needing to recalculate subsequent block hashes). This demonstrates the concept of immutability, although in a very simplified form here.
*   **Proof-of-Work (PoW):**  This project implements a basic Proof-of-Work system.  Adding a new block requires some computational effort (finding a valid nonce). This simulates the idea that adding blocks to a real blockchain requires work, helping to secure the chain.

## Project Structure

```bash
└── felix101930-blockchain-go/
    ├── go.mod
    ├── main.go
    └── blockchain/
        ├── block.go
        ├── blockchain.go
        └── proof.go
```

## How to Run

1.  **Ensure you have Go installed.** You can download Go from the [official Go website](https://go.dev/dl/).

2.  **Navigate to the project directory:**

    ```bash
    cd felix101930-blockchain-go
    ```

3.  **Run the `main.go` file:**

    ```bash
    go run main.go
    ```

    This will execute the `main` function in `main.go`, which will:
    *   Initialize a blockchain.
    *   Add three blocks to it.
    *   Print the details of each block (Previous Hash, Data, Hash, and PoW validation status) to the console.

## TODO

- [ ] **Persistence (Database):** Implement a database to store the blockchain data so it persists between runs.
- [ ] **Transactions:** Move beyond simple text data and implement a transaction structure to represent transfers of value.
- [ ] **Wallet Module:** Create a wallet module to manage user keys and addresses.
- [ ] **Digital Signatures:** Implement digital signatures for transactions to ensure authenticity and prevent tampering.


## Disclaimer

This project is a simplified blockchain implementation for educational purposes.
It serves as a starting point to understand the fundamental concepts behind blockchain technology.
