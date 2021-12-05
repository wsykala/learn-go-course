# Testing

```
godoc -http=:8080

go test
go test -bench .
go test -cover
go test -coverprofile c.out
go tool cover -html=c.out
```