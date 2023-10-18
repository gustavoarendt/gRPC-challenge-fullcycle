## This project was added as a challenge during the event [Imersão Full Stack_ && Full Cycle](https://imersao.fullcycle.com.br/) promoted from [FullCycle](https://fullcycle.com.br/)

### How to run the gRPC server:

- run docker compose: `docker-compose up -d --build`
- Access the app container: `docker exec -it grpc_challenge_app bash`
- Start go application: `go run main.go`

### Having the server up and running, it's possible to create or list products
- From a new terminal (recommended) start evans (already installed in app container): `evans -r repl`
- Create a product: `call CreateProduct`
  - insert name
  - insert description
  - insert price
- Listing products: `call FindProducts`