alias c := compile
alias r := run
alias t := test

# default
default:
    just --list

# run the program
run:
    go run ./

# compile the program
compile:
    go build ./

# test the program
test:
    go test ./...
