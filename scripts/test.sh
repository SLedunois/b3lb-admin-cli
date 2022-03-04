#!/bin/bash

go test -race -covermode=atomic -coverprofile=coverage.out \
    github.com/SLedunois/b3lbctl/pkg/cmd/root \
    github.com/SLedunois/b3lbctl/pkg/cmd/instances \
    github.com/SLedunois/b3lbctl/pkg/admin \
    github.com/SLedunois/b3lbctl/pkg/config \
    github.com/SLedunois/b3lbctl/pkg/system