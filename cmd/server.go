package main

import (
  "github.com/doodles526/replCodeExec/server"
  "github.com/doodles525/replCodeExec/executor"
)

func main() {
  eArgs := executor.PoolArgs {
  }
  executor := executor.Executor(eArgs)

  args := server.Args {
    ListenAddr: "localhost:8080",
    Executor: executor,
  }

  server.ServerCodeExecution(&args)
}