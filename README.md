# go-rest-api

REST API to support read/write/search application metadata as yaml payloads with a integrated in-mem db 

- [validators library](https://github.com/go-playground/validator) is used to validate data received for CreateAppMetadata
- GET and POST to create and get application metadata
- Application supports yaml format
- [Wire](https://github.com/google/wire) is used for Dependency Injections
- [go-memdb library](https://github.com/hashicorp/go-memdb) is used to store data within memory.
- Dockerfile is used to create an image that can be used to run our application on any host machine.
- [Mux](https://github.com/gorilla/mux) is used for creating & routing requests.
- [Zap](https://github.com/uber-go/zap) logger is used to create log and it creates application log file
- [Testify](https://github.com/stretchr/testify) is used to write unit test cases for the application

# How to run the application

## On your machine.
You must have golang installed on your machine

Checkout the reposority
```
git clone https://github.com/bhagvatiB/go-rest-yaml.git
```

Go to the folder
```
cd go-rest-yaml/
```

To run the application, build application with.  
```
go build
```

The execute with
```
./payloadrest
```

From a Browser or Postman try this URL
```
http://localhost:8000/api/v1/createappmetadata

Payload

title: Valid App 1
maintainers:
 - name: firstmaintainer app1
   email: firstmaintainer@hotmail.com
 - name: firstmaintainer app2
   email: firstmaintaine2r@hotmail.com
company: Random Inc.
website: https://website.com
source: https://github.com/random/repo
license: Apache-2.0
description: Interesting Title Some application content, and description
```
```
http://localhost:8000/api/v1/searchappmetadata
```
## Using Docker Container
Checkout the reposority

To build the image with the application run
```
docker build -t go-rest-yaml .
```
To run the image execute
```
docker run -p 8080:8000 go-rest-yaml
```
Try these URLs
```
http://localhost:8080/api/v1/createappmetadata

Payload

title: Valid App 1
maintainers:
 - name: firstmaintainer app1
   email: firstmaintainer@hotmail.com
 - name: firstmaintainer app2
   email: firstmaintaine2r@hotmail.com
company: Random Inc.
website: https://website.com
source: https://github.com/random/repo
license: Apache-2.0
description: Interesting Title Some application content, and description
```
```
http://localhost:8080/api/v1/searchappmetadata
```

Project structure:

- ## cmd/

    This is where main function is defined and all the components
	are being utilized and integrated. main func simple performs;
	
	- initialze logger
   	- initialize server application
	- invoking handlers
	- listen and serve

    Also application log & wire_gen file is generated in this same package.
    Wire is used for dependency injections to initialize the server and return AppMetadataController.

    - ##  controller/

    This Package contains all the end points for this rest API 

    - setup the listeners
    - initialize the mux router 
    - create the get & post endpints
    - create function to create AppMetadata
    - search function to search AppMetadata based on query params

    - ## bizlogic/

    This is where business logic for creating & searching AppMetadata is defined. Right now this doesn't have much code but it can be used when the 	       application evolves and theer is more business logic than just read and write into database.

    - ## dao/

    It is a simple in-memory strorage to store application metadata. This storage is based on Radix tree. More details in [go-memdb].       (https://github.com/hashicorp/go-memdb)
	Supports Insert and Read methods
	[go-memdb library](https://github.com/hashicorp/go-memdb) provides Atomicity, Consistency & isolation properties. This DB will be able to handle concurrent requests.
	Key for storing AppMetadata is "<Title>_<version>"

    - #### entity/

    This has schema for table that stores AppMetadata sent by Client. 

    - ## logger/

    This is where we initialise our logger to be used through out the application to maintain a log file for trouble shooting if necessry.  
    Application logger file will be stored at ${GOPATH}/application.log

    - ## model/

    This package stores all the data Objects that are needed in the service.
    - #### request/

    This has AppMetadata request object structure

    - #### response/

    This has AppMetadata response object structure

    - ## util/

    This is where util files for dao & converter; 

## API Details

Server provides a simple enpoint for GET and POST operations.  
**POST - /api/v1/createappmetadata**
**GET - /api/v1/searchappmetadata**

**POST**  operation is used to create application metadata.   
Yaml is supported as payload

**GET**  operation is used to read application metadata records. Search parameters are being passed via URL.
Supported URL query search parameters are:  

- { version, title, company, website, source, license and description}  

**Note that** Only Prefix search is allowed on any fields. Search with part of the value other than prefix will yield no resuts from the qquery.

## POST OPERATION  

Creates a application metadata. Accepts **yaml** payload. All fields are required. Valid email addresses and URL's are required otherwise returns error. 


Sample POST request is:

```
http://localhost:8000/api/v1/createappmetadata"  
```

Sample application metadata payload is :

```yaml
title: My valid app
version: 1.0.8
company: Ecaglar Inc.
website: https://xyz.com
source: https://github.com/bhagvatiB/
license: Apache-2.1
maintainers:
  - name: Firstname Lastname
    email: bhag@hotmail.com
description: |
    ### blob of markdown
    More markdown
```

## GET OPERATIONS

Changing the URL query parameters, you can query different records.

Samples
**GET - /api/v1/searchappmetadata**.  
Returns all the records from the database.

**GET - /api/v1/searchappmetadata**.  
Returns all records

**GET - /api/v1/searchappmetadata?version=1.0.0**.  
Returns the record with version 1.0.0

**GET - /api/v1/searchappmetadata?version=1.0.0&title=my%20app**.  
Returns the record with version 1.0.0 **OR** title my app

**GET - /api/v1/searchappmetadata?company=mycompany.com&title=my%20app**.  
Returns record(s) with company name "mycompany.com" **OR** title contains "my app"

**GET - /api/v1/searchappmetadata?description=latest**.  
Returns record(s) with description with latest in the prefix

**GET - /api/v1/searchappmetadata?maintainers.email=bill@c.com&license=Apache-2.1**.  
Returns record(s) which have/has maintainers email "bill@hotmail.com" **OR** licence "Apache-2.1"

# Follow Ups
Some more features need to be added or improved
* Add support to search based on Maintainer name and Email
* Add support for search based on any part of the search string
* Add Support for Batch Creation of AppMetadata. Right now API supports only single record creation
* Add support for Update and delete
* Add log rotation to create new log every hour for better trouble shooting.
