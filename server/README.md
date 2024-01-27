# Server

## Go Rest Api Starter

Just a minimal golang rest api starter with some basic stuff I always seem to add.

### Included

Version go 1.21.5

- dependency injection with Fx
- config using Viper
- ping / account / auth controllers
- password hasher / matcher
- paseto token maker + middleware
- app.env, .air.toml, .gitignore
- usefull Makefile

### Packages

- github.com/aead/chacha20poly1305
- github.com/brianvoe/gofakeit/v6
- github.com/gin-gonic/gin
- github.com/google/uuid
- github.com/o1egl/paseto
- github.com/spf13/viper
- github.com/stretchr/testify
- go.mongodb.org/mongo-driver

### Usage

- See `make help` for available commands.

### Resources

https://mholt.github.io/json-to-go/
https://github.com/brianvoe/gofakeit
https://pkg.go.dev/github.com/go-playground/validator/v10
https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
https://chidiwilliams.com/posts/partially-updating-an-embedded-mongo-document-in-go
https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Baked_In_Validators_and_Tags

### Usage

Just clone it, do a search and replace for `oluijks/golang-starter/server`, a `make air` and you're good to go.
