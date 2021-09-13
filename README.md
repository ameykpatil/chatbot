# Chatbot
Chatbot to provide smart replies to messages 

## Summary 
This service exposes a single endpoint accepts a bot identifier and a visitor written message. 
It returns a single reply corresponding to the highest predicted intent above the confidence threshold. 
In order to retrieve the list of predicted intents for a given message publicly available AI API has been used.

This application is written in Golang. 
The database used is MongoDB.

The code is written keeping in mind following qualities: 
- Readable 
- Testable 
- Maintainable 
- Extensible 

Apart from this -
- Basic unit tests are written to demonstrate testing using Go's basic testing framework
- In order to run & check regardless of the environment the support for docker is provided   
- Setup section in Readme is provided for all the setup & run instructions

## Setup & Run
Clone the repository  
`git clone https://github.com/ameykpatil/chatbot`

The application can be run using a Docker container images built & setup using the `Makefile`.

#### Run the application
1. Ensure you are in the root directory of the repository.
2. Run `make docker-up`, this will start the required containers (`http`, `mongodb`)
3. Check if server is running by checking `/ping` (you should get `pong` as response)  
   `curl --request GET 'http://localhost:8000/ping'` 
4. Check the `/replies` api with several messages  
```bash
curl --request GET 'http://localhost:8000/replies?bot_id=5f74865056d7bb000fcd39ff&message=Hello'
curl --request GET 'http://localhost:8000/replies?bot_id=5f74865056d7bb000fcd39ff&message=I can not access the app'
curl --request GET 'http://localhost:8000/replies?bot_id=5f74865056d7bb000fcd39ff&message=Okay'
curl --request GET 'http://localhost:8000/replies?bot_id=5f74865056d7bb000fcd39ff&message=It worked, thank you'
curl --request GET 'http://localhost:8000/replies?bot_id=5f74865056d7bb000fcd39ff&message=Bye'
```

#### Run the tests
Unit tests can be run using `make`
```bash
make test
```

#### Run the linter
Linter can be run using `make`  
But `golangci-lint` will be needed to run it
```bash
make lint
```

## Application Design

- Application has been designed & structured in a layered format. Following diagram should help to visualise the flow through different layers.  
<img width="500" alt="Screenshot 2021-09-13 at 13 28 36" src="https://user-images.githubusercontent.com/3050421/133075958-b64458c8-ee18-4f01-a5a7-9b8f3c2ce4c9.png">  

**Handler** : Handler defines how to handle specific request & return response. It makes use of Service layer for getting required processed details.      

**Service** : This layer provides a service to get data from different sources & combine them in a more meaningful way.   

**API Client** : It provides a way to access external api to get intents.

**Storage** : This layer provides a way to access database. In this case MongoDB.  
_Note: ReplyWriter is written just to load initial data, more recommended way is to implement properly the way Reader is implemented & then provide a POST api to create the data._    

## Enhancements
Using additional time we can add following things to an application
- `POST /replies` API to add replies to a database instead of current static set of replies  
- Integration tests with mocked DB & third-party API 
- If we know more specific way to choose intent from received intents, implementing that logic. 
- If we know more specific way to choose reply, implement that logic. 