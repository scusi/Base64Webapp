# Base64Webapp
Base64Webapp is a Webapp - written in go - to encode and decode base64.

## Usage

You have to set an environment variable in order to set the local port to be used for the webserver. How to do this depends on your shell environment on most unix systems the following should work.

    export PORT=80

The above command sets the port to 80, the default webserver port. However you can specify any TCP port.

To run the programm one can simply execute the following command

    go run Base64Webserver.go
    
or if you prefer to build a binary you can do some thing like this

    go build Base64Webserver.go
    export PORT=80
    ./Base64Webserver

However you choose to do it, after you are done point your browser to http://127.0.0.1, or if you have choosen another port than 80 you have to add the port like this http://127.0.0.1:<PORT>/. 

## Use with docker

If you prefer to run it in docker, try this

    docker pull scusi/base64webapp
    docker run --publish 80:80 -d scusi/base64webapp

To find your docker IP adress when useing boot2docker (e.g. on MAC) use the following command

    boot2docker ip
