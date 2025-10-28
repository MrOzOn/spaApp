# SPA_WEB

This is example app - single page application with go backend implementation.

**Frontend** 
- Framework [Vite](https://github.com/vitejs/vite) with pattern REACT+TS

**Backend**
- Go 1.25 (no need any dependencies)
```shell
 go get -u github.com/gin-gonic/gin
 go get -u github.com/gin-contrib/cors
```

## IMPORTANT
!!! Before build native executing app you need build production bundle for javascript !!!

Bundle located in folder `frontend/dist`

## Some useful commands

1. create frontend template using vite (in `root` dir)
```shell
npm create vite@latest frontend -- --template react-ts
```

2. install nodejs dependencies (in `frontend` dir)
```shell
npm i
```

3. run frontend for coding with hot reload (in `frontend` dir)
```shell
npm run dev
```

4. run frontend for production (in `frontend` dir)
```shell
npm run build
```

5. build native program
```shell
go build -o spa_web
```

6. build native program for selected OS (for example, Window x64)
```shell
GOOS=windows GOARCH=amd64 go build -o spa_web.exe
```