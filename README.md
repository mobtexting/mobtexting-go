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

        var repsonse = client.Send(to, from, body, service, access_token)

	fmt.Println(repsonse)
}

```

