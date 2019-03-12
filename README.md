# golang plugin for mobtexting

This package makes it easy to send [Mobtexting notifications](https://mobtexting.com).

## Installation

You can install the package using go get :

``` bash
go get github.com/mobtexting/mobtexting-go
```

## Send SMS Usage

```go
package main
import "fmt"
import client "github.com/mobtexting/mobtexting-go"

func main() {
        var to = "1234567890"
        var from = "MobTxt"
        var body = "hello from go"
        var service = "P"
        var access_token = "xxxxxxxxxxxxxxxxxxxxxxxxxxxx"

        var response = client.Send(to, from, body, service, access_token)

	    fmt.Println(response)
}

```

## Verify Usage
### Send
```go
var to = "1234567890"
var access_token = "xxxxxxxxxxxxxxxxx"
var response = client.VerifySend(to, access_token)
fmt.Println(response)
```

### Check
```go
var id = "b51be650-fdb2-4633-b101-d450e8d9ec64", // id received while sending
var token = "123456" // token entered by user
var access_token = "xxxxxxxxxxxxxxxxx"
var response = client.VerifyCheck(id, token, access_token)
fmt.Println(response)
```

### Cancel
```go
var id = "b51be650-fdb2-4633-b101-d450e8d9ec64", // id received while sending
var access_token = "xxxxxxxxxxxxxxxxx"
var response = client.VerifyCancel(id, access_token)
fmt.Println(response)
```
