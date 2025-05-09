module blueberry_homework

go 1.24.2

require github.com/go-chi/cors v1.2.1

require (
	github.com/go-chi/chi/v5 v5.2.1
	github.com/gocql/gocql v1.7.0
	github.com/google/uuid v1.6.0
)

require (
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)

replace github.com/gocql/gocql => github.com/scylladb/gocql v1.14.5
