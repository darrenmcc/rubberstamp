.PHONY: build install clean

emojis_generated.go: all_emojis.txt generate_emojis.go
	go generate

build: emojis_generated.go
	go build -o stamp

install: emojis_generated.go
	go install

clean:
	rm -f rubberstamp emojis_generated.go
