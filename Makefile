all: clean cointool

clean:
	rm -f cointool

cointool:
	go build -o cointool ./cmd/cli