package main

import (
  "fmt"
  "github.com/tendermint/tendermint/abci/types"
  "github.com/tendermint/tendermint/abci/server"
)

func main(){
  app := types.NewBaseApplication()
  svr,err := server.NewServer(":26658","socket",app)
  if err != nil { panic(err) }
  svr.Start()
  defer svr.Stop()
  fmt.Println("abci server started.")
  select {}
}
