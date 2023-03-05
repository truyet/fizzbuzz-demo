# fizzbuzz-demo

## Prerequire
- Golang sdk
- Nodejs

## Run
 
### Backend

Backend will run on port `8888`

```shell

cd backend/fizzbuzz
go mod tidy
go run fizzbuzz.go -f etc/fizzbuzz-api.yaml

```

### Frontend

Frontend will run on port `3000`

```shell

cd frontend
yarn
yarn start
```

Open browser with url `http://localhost:3000/`

Demo video included `demo.mov`