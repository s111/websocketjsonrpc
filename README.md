# websocketjsonrpc
JSON-RPC over websocket


## Example usage
```go
package main

import (
    "log"
    "net/http"
    "net/rpc"
    "net/rpc/jsonrpc"

    "github.com/gorilla/websocket"
    "github.com/s111/websocketjsonrpc"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

type Args struct {
    A int
    B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
    *reply = args.A * args.B

    return nil
}

func main() {
    rpc.Register(new(Arith))

    http.HandleFunc("/ws", serveWs)

    err := http.ListenAndServe(":3001", nil)

    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func serveWs(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Method not allowed", 405)

        return
    }

    ws, err := upgrader.Upgrade(w, r, nil)

    if err != nil {
        log.Println("Upgrade:", err)

        return
    }

    jsonrpc.ServeConn(websocketjsonrpc.NewConn(ws))
}
```
