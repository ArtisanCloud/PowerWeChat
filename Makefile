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

	go test -v test/featureUnit/main_test.go test/featureUnit/work_externalcontact_get_test.go

test-externalcontact-message:
	go test -v test/featureUnit/main_test.go test/featureUnit/work_externalcontact_add_msg_template_test.go
	go test -v test/featureUnit/main_test.go test/featureUnit/work_externalcontact_get_group_msg_result_test.go

test-media:
	go test -v test/featureUnit/main_test.go test/featureUnit/work_media_test.go

test-payment:
	go test -v test/featureUnit/main_test.go test/featureUnit/payment_order_test.go
	go test -v test/featureUnit/main_test.go test/featureUnit/payment_jssdk_test.go
	go test -v test/featureUnit/main_test.go test/featureUnit/payment_redpack_test.go

build:
	go build