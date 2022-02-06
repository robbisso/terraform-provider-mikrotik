.PHONY: import testacc testclient test

TIMEOUT ?= 40m
ifdef TEST
    TEST := ./... -run $(TEST)
else
    TEST := ./...
endif

ifdef TF_LOG
    TF_LOG := TF_LOG=$(TF_LOG)
endif

build:
	go build -o terraform-provider-mikrotik

clean:
	rm dist/*

plan: build
	terraform init
	terraform plan

apply:
	terraform apply

test: testclient testacc

testclient:
	cd client; go test $(TEST) -race -v -count 1

testacc:
	TF_ACC=1 $(TF_LOG) go test $(TEST) -v -count 1 -timeout $(TIMEOUT)

install: build
	mkdir -p ~/.terraform.d/plugins/terraform.local/local/mikrotik/1.0.0/linux_amd64
	cp terraform-provider-mikrotik ~/.terraform.d/plugins/terraform.local/local/mikrotik/1.0.0/linux_amd64/terraform-provider-mikrotik

