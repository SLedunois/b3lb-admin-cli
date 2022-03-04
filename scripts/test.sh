#!/bin/bash

go test -race -covermode=atomic -coverprofile=coverage.out \
    github.com/SLedunois/b3lb-admin-cli/pkg/cmd/root \
    github.com/SLedunois/b3lb-admin-cli/pkg/cmd/instances \
    github.com/SLedunois/b3lb-admin-cli/pkg/admin \
    github.com/SLedunois/b3lb-admin-cli/pkg/config \
    github.com/SLedunois/b3lb-admin-cli/pkg/system