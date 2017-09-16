# ntry-backend

Setup `go 1.8+`. Install required system packages

Start Maria DB docker container with below command. Update `ntryapp.yml` according to your customization in below command.
```
docker run -e MYSQL_ROOT_PASSWORD=toor -e MYSQL_DATABASE=ntry -p3306:3306 -it mariadb
```

Download ntry-backend, and install dependencies using `glide`:
```
go get github.com/NTRYPlatform/ntry-backend
cd $GOPATH/src/github.com/NTRYPlatform/ntry-backend
curl https://glide.sh/get | sh
glide install
glide update
```
Build NTRY server:
```
go build app/server.go
```
To run/test NTRY server:
```
./server -c .notaryconf/ntryapp.yml -l /var/log/notary.log
```
