test: test-aes

test-aes:
	go test -v src/kernel/Encryptor.go src/kernel/Encryptor_test.go
	go test -v src/kernel/Support/AES.go src/kernel/Support/AES_test.go

build:
	go build