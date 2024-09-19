# Problem 6: Transaction Broadcaster Service
This solution treats the EVM-compatible blockchain as a black box.

The system consists of:
  * API that interacts with users
  * Main broadcaster service logic (3 layers)
    1. Layer that interacts with API
    2. Layer that interacts with blockchain nodes
    3. Data storage layer
  * Admin interface

---

![Illustration of layers](layers.png)

---
# Detailed description of layers
## API layer:
  * Receives broadcast requests (`message_type`, `data`)
  * Forwards broadcast requests to the main broadcaster service
  * Upon successful forward, returns HTTP `200`
  * Otherwise, returns HTTP `4xx`-`5xx`

## Broadcaster service
### Transaction processing
  * Receives broadcast requests from API layer
  * Signs `data`
  * Outputs `signed transaction` to blockchain-side layer
### Blockchain-side
  * Receives `signed transaction` from tx processor
  * Send PRC request to a blockchain node
  * Also sends PRC request on command
  * All PRC requests are logged in data layer
### Data layer
  * Maintains a list of transactions + status (`pending`, `success`, `failure`) + time of PRC request
  * On `failure`, retry (calls blockchain-side layer to retry)
  * If `pending` for more than 30s, retry (calls blockchain-side layer to retry)
  * Upon unexpected restart, retry all `pending` and `failure` broadcasts

## Admin interface
  * Access to [data layer](#data-layer)

---
# Flow of events
1. API receives valid request  
2. API forwards request to tx processor  
   2.1 If tx processor fails to respond after 30s, API returns HTTP `4xx`-`5xx`  
   2.2 If error occurs, API returns error code  
3. Upon receiving the request, tx processor sends a `success` message back  
   3.1. API returns HTTP `200`  
4. Tx processor signs the message  
5. Tx processor sends signed message to blockchain-side layer  
   5.1 Tx processor logs signed message in data layer  
6. Blockchain-side layer receives signed message from tx processor
7. Blockchain-side layer sends a PRC request to a blockchain node  
   7.1 Blockchain-side layer logs the request in data layer  
   7.2 If no response is received after 30s, blockchain-side layer sends the request again  
   7.3 If error occurs, blockchain-side layer sends the request again  
8. Blockchain-side layer receives success confirmation from blockchain node  
   8.1 Blockchain-side layer logs the successful request in data layer  

In case of system failure and unexpected restart:
1. Blockchain-side layer resends all `pending` and `failure` transactions in data layer

---


ChatGPT was consulted in the making of this solution.
