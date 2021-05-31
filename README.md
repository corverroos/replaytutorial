# replay tutorial

This repo contains a tutorial that introduces developers to the [replay](https://github.com/corverroos/replay) framework.

Each sub-folder contains a simple go program. Some are purely examples and can just be studied and run with `go run`. 
Others contain `TODO(you)` that must completed before the program will run successfully.

- [00_helloworld](./00_helloworld/main.go) 
- [01_typedhello](./01_typedhello/main.go) 
- [02_signalhello](./02_signalhello/main.go) 
- [03_outputhello](./03_outputhello/main.go) 
- [04_restarthello](./04_restarthello/main.go) 
- [05_cron](./05_cron/main.go) 
- [06_appendonlylog](./06_appendonlylog/main.go) 
- [07_makertrade](./07_makertrade/main.go) 
- [08_sparkline](./08_sparkline/main.go) 
- [09_parentjob](./09_parentjob/main.go) 
- [10_distributed](./10_distributed/main.go) 

#### TODO
- sharded activity consumers
- versioning

## Tips
- Read the [replay](https://github.com/corverroos/replay) docs first (can skip *Under the hood*)

- Read each exercise from top to bottom (including comments).

- Running the programs can be done in IDE's supporting `go:generate` or from the command line:
```
go run github.com/corverroos/replaytutorial/00_helloworld
# Exit early with Ctrl-C to test restart robustness
```

- Regularly install the latest `typedreplay` code generation tool:
```
go install github.com/corverroos/replay/typedreplay/cmd/typedreplay@latest
```

- Some exercises contain `//showme:hidden foo` directives associated with a `TODO(you)` instruction. 
Executing the following `go:generate` command will show the next 1 solution in the go file. 
Note that the `-hide` flag will hide the solutions.
```
//go:generate go run ../lib/showme 1 
``` 

- This tutorial assumes a mysql DB is accessible at `mysql://root@unix(/tmp/mysql.sock)/`. 
If not specify the connection url via `db_url`. Note it creates a long-lived 
database called `replay_tutorial`. Use the flag `db_refresh` to clear the database.

- Experiment!

