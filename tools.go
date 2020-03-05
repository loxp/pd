// +build tools

package tools

import (
	_ "github.com/go-playground/overalls"                   // 22ec1a223b7c9a2e56355bd500b539cba3784238
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint" // 1.17.1
	_ "github.com/kevinburke/go-bindata/go-bindata"         // c7c189834d8e8384b8aa7428c02e4deee6df6b17
	_ "github.com/mgechev/revive"                           // 7773f47324c2bf1c8f7a5500aff2b6c01d3ed73b
	_ "github.com/pingcap/failpoint/failpoint-ctl"          // 30cc7431d99c6a7f2836387d4bb255a3bd6a5e0a
	_ "golang.org/x/tools/cmd/goimports"                    // 04b5d21e00f1f47bd824a6ade581e7189bacde87
)
