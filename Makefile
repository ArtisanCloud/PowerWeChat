test: test-aes test-message

test-aes:
	go test -v src/kernel/Encryptor.go src/kernel/Encryptor_test.go
	go test -v src/kernel/Support/AES.go src/kernel/Support/AES_test.go

test-message:
    go test -v test/featureUnit/main_test.go test/featureUnit/work_message_test.go

build:
	go build