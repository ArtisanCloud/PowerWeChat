test: test-aes test-rsa test-signer

test-aes:
#	go test -v src/kernel/Encryptor.go src/kernel/Encryptor_test.go
	go test -v src/kernel/support/aes.go src/kernel/support/aes_test.go

test-rsa:
	go test -v src/kernel/support/rsa_oaep.go src/kernel/support/rsa_oaep_test.go

test-signer:
	go test -v src/kernel/support/signer.go src/kernel/support/signer_test.go

test-message:
	go test -v test/featureUnit/main_test.go test/featureUnit/work_message_test.go

test-external-contact:
	go test -v test/featureUnit/main_test.go test/featureUnit/work_externalcontact_add_contact_way_test.go
	go test -v test/featureUnit/main_test.go test/featureUnit/work_externalcontact_get_contact_way_test.go
	go test -v test/featureUnit/main_test.go test/featureUnit/work_externalcontact_update_contact_way_test.go
	go test -v test/featureUnit/main_test.go test/featureUnit/work_externalcontact_del_contact_way_test.go

	go test -v test/featureUnit/main_test.go test/featureUnit/work_externalcontact_get_test.go

test-external-contact-message:
	go test -v test/featureUnit/main_test.go test/featureUnit/work_externalcontact_add_msg_template_test.go
	go test -v test/featureUnit/main_test.go test/featureUnit/work_externalcontact_get_group_msg_result_test.go

test-media:
	go test -v test/featureUnit/main_test.go test/featureUnit/work_media_test.go

test-payment:
	go test -v test/featureUnit/main_test.go test/featureUnit/payment_order_test.go
	go test -v test/featureUnit/main_test.go test/featureUnit/payment_jssdk_test.go
	go test -v test/featureUnit/main_test.go test/featureUnit/payment_redpack_test.go

test-kernel-support:
	go test -v src/kernel/support/str.go src/kernel/support/str_test.go

build:
	go build