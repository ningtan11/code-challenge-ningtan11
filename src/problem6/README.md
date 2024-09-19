# Problem 6: Transaction Broadcaster Service
This solution treats the EVM-compatible blockchain as a black box.
## API layer:
  * Receives broadcast requests (`message_type`, `data`)
  * Returns HTTP `200` or HTTP `4xx`-`5xx`

## Broadcaster service
### Transaction processing
  * Signs `data`
  * Outputs `signed transaction`
### Blockchain-side
  * Receives `signed transaction`
  * PRC request to a blockchain node
  * PRC Request on command
### Data layer
  * Maintains a list of transactions + status (`pending`, `success`, `failure`) + time of PRC request
  * On `failure`, retry (calls blockchain-side layer to retry)
  * If `pending` for more than 30s, retry

## Admin interface
  * Access to [data layer](#data-layer)
