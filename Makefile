Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")
SERVICES=$(shell ls -1 api | grep \.proto | sed s/\.proto//)


.PHONY: bin
bin: $(info $(M) install bin)
	GOBIN=$(CURDIR)/bin go install -mod=mod github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc


.PHONY: gen
gen: bin; $(info $(M) protoc gen)
	$(Q) for srv in $(SERVICES); do \
		echo "Generate $$srv" && \
		mkdir -p $(CURDIR)/cmd/$$srv && \
		mkdir -p $(CURDIR)/internal/$$srv && \
		mkdir -p $(CURDIR)/pkg/$$srv && \
		mkdir -p $(CURDIR)/docs/swagger && \
        protoc --plugin=protoc-gen-grpc-gateway=$(CURDIR)/bin/protoc-gen-grpc-gateway \
            --plugin=protoc-gen-openapiv2=$(CURDIR)/bin/protoc-gen-openapiv2 \
            --plugin=pprotoc-gen-go-grpc=$(CURDIR)/bin/protoc-gen-go-grpc \
            -I$(CURDIR)/api:$(CURDIR)/vendor.pb \
            --go_out=$(CURDIR)/pkg \
            --go-grpc_out=$(CURDIR)/pkg \
            --grpc-gateway_out=$(CURDIR)/pkg \
            --grpc-gateway_opt=logtostderr=true \
            --openapiv2_out=$(CURDIR)/docs/swagger \
            --openapiv2_opt=logtostderr=true \
            --openapiv2_opt=use_go_templates=true \
            $(CURDIR)/api/$$srv.proto ; \
	done


.PHONY: mockgen
mockgen:
	    go generate ./...

.PHONY: run
run:
	go run cmd/echo_service/main.go
