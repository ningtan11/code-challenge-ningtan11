# Problem 6: Transaction Broadcaster Service
## API layer:
  * Receives broadcast requests (`message_type`, `data`)
  * Returns HTTP `200` or HTTP `4xx`-`5xx`

## Broadcaster service
### Layer 1
  * Signs `data`
  * Outputs `signed transaction`
### Layer 2
  * Receives `signed transaction`
  * PRC request to a blockchain node
### Layer 3
  * Maintains a list of transactions + status (pending, success, failure)

## Blockchain
### 

