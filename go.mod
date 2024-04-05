module github.com/r-usenko/go-migrator

go 1.22

//replace (
//	github.com/r-usenko/go-migrator/drivers/postgres => ./drivers/postgres
//	github.com/r-usenko/go-migrator/drivers/processor => ./drivers/processor
//)
//
//exclude (
//	github.com/r-usenko/go-migrator/drivers/postgres v0.0.0-20240405004535-5911085a1b84
//	github.com/r-usenko/go-migrator/drivers/processor v0.0.0-20240405004535-5911085a1b84
//)

require (
	github.com/r-usenko/go-migrator/drivers/postgres v0.0.0-20240405004535-5911085a1b84
	github.com/r-usenko/go-migrator/drivers/processor v0.0.0-20240405004535-5911085a1b84
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
