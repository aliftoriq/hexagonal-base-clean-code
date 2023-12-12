### Generate Unit Test
 `mockery --all`

### Generate Coverage Unit Test
`go test -coverpkg=./internal/core/usecase/... -coverprofile=coverage ./internal/tests/...  `

### Open in html Coverage
`go tool cover -html=coverage; `