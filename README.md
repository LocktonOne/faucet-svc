# faucet

## Description

[//]: # (TODO: add description)

## Install

  ```bash
  git clone gitlab.com/tokene/faucet
  cd faucet
  go build main.go
  export KV_VIPER_FILE=./config.yaml
  ./main run service
  ```

## Documentation

We do use openapi:json standard for API. We use swagger for documenting our API.

To open online documentation, go to [swagger editor](http://localhost:8080/swagger-editor/) here is how you can start it
```bash
  cd docs
  npm install
  npm start
```
To build documentation use `npm run build` command,
that will create open-api documentation in `web_deploy` folder.

To generate resources for Go models run `./generate.sh` script in root folder.
use `./generate.sh --help` to see all available options.


## Running from docker 
  
Make sure that docker installed.
{%_ if (handleHTTP) { _%}
use `docker run ` with `-p 8080:80` to expose port 80 to 8080

{%_ } _%}

    ```bash
    docker build -t gitlab.com/tokene/faucet .
    docker run -e KV_VIPER_FILE=/config.yaml gitlab.com/tokene/faucet
    ```

## Running from Source

* Set up environment value with config file path `KV_VIPER_FILE=./config.yaml`
* Provide valid config file
* Launch the service with `run service` command



### Third-party services


## Contact

Responsible 
The primary contact for this project is  [//]: # (TODO: place link to your telegram and email)
