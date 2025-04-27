BINARY := bin/server
CMD     := ./cmd/server

.PHONY: build run clean

# Build the server binary into ./bin/server
build:
	@mkdir -p bin
	go build -o $(BINARY) $(CMD)

# Rebuild (if needed) then run
run: build
	@echo "→ Running $(BINARY)"
	./$(BINARY)

# Remove the build artifacts
clean:
	@echo "→ Cleaning up"
	rm -rf bin
