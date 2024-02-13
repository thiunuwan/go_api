GET:
   curl localhost:8080/books 

POST:
    curl localhost:8080/books --include --header "Content-TYpe: application/json" -d @newBookData.json --request "POST"