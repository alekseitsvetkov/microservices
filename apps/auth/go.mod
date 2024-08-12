module github.com/alekseytsvetkov/microservices/apps/auth

go 1.22.5

replace (
	github.com/alekseytsvetkov/microservices/libs/hash => ../../libs/hash
	github.com/alekseytsvetkov/microservices/libs/jwt => ../../libs/jwt
	github.com/alekseytsvetkov/microservices/proto => ../../proto
)

require (
	github.com/alekseytsvetkov/microservices/libs/hash v0.0.0-00010101000000-000000000000
	github.com/alekseytsvetkov/microservices/libs/jwt v0.0.0-00010101000000-000000000000
	github.com/alekseytsvetkov/microservices/proto v0.0.0-00010101000000-000000000000
	github.com/ilyakaznacheev/cleanenv v1.5.0
	go.uber.org/fx v1.22.1
	google.golang.org/grpc v1.65.0
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	go.uber.org/dig v1.17.1 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/crypto v0.25.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
