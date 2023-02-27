# Go Api
---  

Simple rest api from the 2. task from the list
To access this API use `POST` method on the `http://localhost:7000/upload` url with request body of the type like this:  
`{
    "operation":"deduplicate",
    "data": [1, 1, 1, 2, 4, 5, 5, 6, 7, 7, 8]
}`

---

To start the Docker container just use   
`docker compose up` command.
