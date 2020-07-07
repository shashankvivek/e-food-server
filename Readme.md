Cmd to generate new changes in swagger :

> swagger generate server -f api/swagger.yaml --default-scheme http --exclude-main

To run the repository locally :

> go run cmd\e-food-server\main.go --scheme http --port=8080
