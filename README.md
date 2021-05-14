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

## Tips
Read the [replay](https://github.com/corverroos/replay) docs first (can skip *Under the hood*)

Read each exercise from top to bottom (including comments).

Running the programs can be done in IDE's supporting `go:generate` or from the command line:
```
go run github.com/corverroos/replaytutorial/00_helloworld
# Exit early with Ctrl-C to test restart robustness
```

Install the `typedreplay` code generation tool:
```
go install github.com/corverroos/replay/typedreplay/cmd/typedreplay@latest
```

It assumes a mysql DB is accessible at `mysql://root@unix(/tmp/mysql.sock)/`. 
If not specify the connection url via `db_url`. Note it creates a long-lived 
database called `replay_tutorial`. Use the flag `db_refresh` to clear the database.

Experiment!

