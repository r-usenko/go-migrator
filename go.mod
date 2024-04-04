module github.com/r-usenko/go-migrator

go 1.22

replace (
	github.com/r-usenko/go-migrator/drivers/postgres => ./drivers/postgres
	github.com/r-usenko/go-migrator/drivers/processor => ./drivers/processor
)

require github.com/r-usenko/go-migrator/drivers/processor v0.0.0-00010101000000-000000000000
