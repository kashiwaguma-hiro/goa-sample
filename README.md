# goa-sample


## How to create.

https://goa.design/learn/getting-started/


logs 

```bash
go mod init github.com/kashiwaguma-hiro/goa-sample
go get -u goa.design/goa/v3
go get -u goa.design/goa/v3/...

# create go code from design,move the oas3.0 docs
make gen
```


## How to running
```bash
# create the sample code from design
goa example github.com/kashiwaguma-hiro/goa-sample/design

# editing code generated ./admin.go

# Start the api
go run ./cmd/admin

# testing post
% curl -i http://localhost:8000/users/ \
  -H 'Content-Type:application/json' \
  -X POST \
  -d '{"first-name":"aaa", "last-name":"bbbb", "email":"aaa-bbb@example.com", "tel":"000-0000-00000"}'
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 31 May 2021 12:07:05 GMT
Content-Length: 3

{}
```