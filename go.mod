module github.com/r-usenko/go-migrator

go 1.22

//exclude (
//	github.com/r-usenko/go-migrator/drivers/postgres latest
//)

replace (
	github.com/r-usenko/go-migrator/drivers/postgres => ./drivers/postgres
	github.com/r-usenko/go-migrator/drivers/processor => ./drivers/processor
)

require (
	github.com/r-usenko/go-migrator/drivers/postgres v0.0.0-00010101000000-000000000000
	github.com/r-usenko/go-migrator/drivers/processor v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
