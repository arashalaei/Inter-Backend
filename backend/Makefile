SOURCE_DIR := .\contracts\src
ABI_TARGET_DIR := build\abi
BIN_TARGET_DIR := build\bin

run:
	go run ./cmd/main.go
	
build:
	for /r "$(SOURCE_DIR)" %%F in (*.sol) do $(MAKE) compile FILE="%%F"


compile:
	solcjs --optimize --abi $(FILE) -o $(ABI_TARGET_DIR)
	solcjs --optimize --bin $(FILE) -o $(BIN_TARGET_DIR)
