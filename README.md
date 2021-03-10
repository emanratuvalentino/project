# API Docudigital (Simpeg)

## Pre-installation
You have to install all of this package before running.
- `go get -u -v golang.org/x/lint/golint`
- `go get -u -v github.com/labstack/echo/v4`
- `go get -u -v github.com/jinzhu/gorm`
- `go get -u -v github.com/tkanos/gonfig`
- `go get -u gorm.io/gorm `



## Configurations
Copy `config.json.example` to `config.json` as your own config file.
Configuration used gonfig. All configs are declared in `config/config.json`

## Architecture
| Folder | Details |
| --- | ---|
| api | Holds the api endpoints |
| db | Database Initializer and DB manager |
| helper | Commonly used global function |
| route | router setup |
| model | Models|


## Run 
`go run server.go`


