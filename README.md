# Ebanx API Challenge
My implementation for the Ebanx API challenge

## Balance Package Documentation
The balance package provides a HTTP handler for retrieving the balance of an account in a financial application. It uses the Gin web framework for handling HTTP requests and responses.

Functions
```BalanceHandler(c *gin.Context)```
This function is a HTTP handler that retrieves the balance of an account. It takes a ```*gin.Context``` object as an argument, which contains all the HTTP request information.

The function works as follows:

Extracts the ```account_id``` from the query parameters of the HTTP request.
Calls the ```GetBalance``` function from the ```ebanx.api/account/domain/use_cases``` package with the ```account_id``` as an argument. This function retrieves the balance for the given account ID.
If there is an error (i.e., if ```balance["error"]``` is not nil), it logs the error message and sends a HTTP response with status code 404 (Not Found) and the error message as the response body.
If there is no error, it sends a HTTP response with status code 200 (OK) and the balance amount as the response body.

## Event Package
The event package is responsible for handling different types of events that can occur in the system. It uses the ```gin-gonic/gin``` package to handle HTTP requests and responses.

### EventHandler
```EventHandler``` is the main function that handles incoming events. It takes a ```gin.Context```, binds the incoming ```JSON``` to an ```EventRequestBody``` struct, and calls the appropriate handler based on the event type. If the event type is not recognized, it responds with a 400 Bad Request status.

### DepositHandler
```DepositHandler``` handles deposit events. It takes an ```EventRequestBody``` and a ```gin.Context```, extracts the destination and amount from the body, calls the ```ChangeFunds``` use case, and responds with a 201 Created status and a ```JSON``` object containing the destination id and balance. If an error occurs, it responds with a 404 Not Found status.

### WithdrawHandler
```WithdrawHandler``` handles withdraw events. It takes an ```EventRequestBody``` and a ```gin.Context```, extracts the origin and amount from the body, calls the ```WithdrawFunds``` use case, and responds with a 201 Created status and a ```JSON object``` containing the origin id and balance. If an error occurs, it responds with a 404 Not Found status.

### TransferHandler
```TransferHandler``` handles transfer events. It takes an ```EventRequestBody``` and a ```gin.Context```, extracts the origin, destination, and amount from the body, calls the ```TransferFunds``` use case, and responds with a 201 Created status and a JSON object containing the origin and destination id and balance. If an error occurs, it responds with a 404 Not Found status.


---

# Automated Test Results

```
✅ Reset state before starting tests
POST /reset
Expected: 200 OK
Got:      200 OK

✅ Get balance for non-existing account
GET /balance?account_id=1234
Expected: 404 0
Got:      404 0

✅ Create account with initial balance
POST /event {"type":"deposit", "destination":"100", "amount":10}
Expected: 201 {"destination": {"id":"100", "balance":10}}
Got:      201 {"destination":{"balance":10,"id":"100"}}

✅ Deposit into existing account
POST /event {"type":"deposit", "destination":"100", "amount":10}
Expected: 201 {"destination": {"id":"100", "balance":20}}
Got:      201 {"destination":{"balance":20,"id":"100"}}

✅ Get balance for existing account
GET /balance?account_id=100
Expected: 200 20
Got:      200 20

✅ Withdraw from non-existing account
POST /event {"type":"withdraw", "origin":"200", "amount":10}
Expected: 404 0
Got:      404 0

✅ Withdraw from existing account
POST /event {"type":"withdraw", "origin":"100", "amount":5}
Expected: 201 {"origin": {"id":"100", "balance":15}}
Got:      201 {"origin":{"balance":15,"id":"100"}}

✅ Transfer from existing account
POST /event {"type":"transfer", "origin":"100", "amount":15, "destination":"300"}
Expected: 201 {"origin": {"id":"100", "balance":0}, "destination": {"id":"300", "balance":15}}
Got:      201 {"destination":{"balance":15,"id":"300"},"origin":{"balance":0,"id":"100"}}

✅ Transfer from non-existing account
POST /event {"type":"transfer", "origin":"200", "amount":15, "destination":"300"}
Expected: 404 0
Got:      404 0

```