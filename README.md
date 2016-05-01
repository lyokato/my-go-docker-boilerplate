## Change Username

```
vi ./build/web.sh
```

change **lyokato** to your username on DockerHub
```
...
docker build -t lyokato/myapp-web:0.1 .
...
```

## Change ProjectName

In default, project name is **myapp**,
replace it to your project name.

You can do easily with 

```
cd ./helper/project_setting/
vi main.go
```

modify, **yourapp**
```
var origin = "myapp"
var modified = "yourapp"
```

And execute

```
go run main.go
```

## Development

working in ./src/proj
and don't forget to resolve dependencies, (exp: "go get" or "glide up")
before you run container

```
/develop/web/run.sh
/develop/web/stop.sh
```

## Build

```
./build/web.sh
```
