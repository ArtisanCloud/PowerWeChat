test: test-aes test-message

test-aes:
	go test -v src/kernel/Encryptor.go src/kernel/Encryptor_test.go
	go test -v src/kernel/Support/AES.go src/kernel/Support/AES_test.go

test-message:
	go test -v test/featureUnit/main_test.go test/featureUnit/work_message_test.go

test-externalcontact:
	go test -v test/featureUnit/main_test.go test/featureUnit/work_externalcontact_add_contact_way_test.go
	go test -v test/featureUnit/main_test.go test/featureUnit/work_externalcontact_get_contact_way_test.go
	go test -v test/featureUnit/main_test.go test/featureUnit/work_externalcontact_update_contact_way_test.go
	go test -v test/featureUnit/main_test.go test/featureUnit/work_externalcontact_del_contact_way_test.go


build:
	go build