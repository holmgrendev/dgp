# DGP
**Template for building a GO/PostgreSQL web server with Docker.**

This template is intended for my personal use, but i am sharing it if someone else finds it useful.

This script creates two containers, one for the Go program and one for the PostgreSQL database.

# How to
Create your GO software and use the folder structure as suggested.

Rename the file `.env-sample` in `build/docker/` to `.env` and edit the file to define name of database, username and password.

Build and start the docker container: `./build/docker.sh -s`

Exit and remove the docker container: `./build/docker.sh -k`

Restart and rebuild the docker container: `./build/docker.sh -r`

# Examples
*This repo contains examples.*

All packages in `internal/` are example packages.

There is example code in the `main.go` file.

All files in `public/assets/` and `public/templates/` are examples.

You can try the code before doing anything by executing the build and start command `./build/docker.sh -s`

# Folder structure
dgp/
├─ build/
│  ├─ docker/
│  │  ├─ .env-sample
│  ├─ docker.sh
├─ cmd/
│  ├─ main/
│  │  ├─ main.go
├─ internal/
│  ├─ (package).../
│  │  ├─ (package.go)
├─ public/
│  ├─ assets/
│  │  ├─ images/
│  │  ├─ scripts/
│  │  ├─ styles/
│  ├─ templates/
│  │  ├─ elements/
│  │  ├─ pages/

## Information about files and folders

`build/docker.sh`
This is the build script

`cmd/main/main.go`
Write your main code here.

`internal/`
This is where your packages should be located.

`public/`
Everything in this folder will be copied to the final docker container

`public/assets/`
This is where the public static content is located.
All files here can be referenced from the html. `public/` is the root folder.

Example:
```html
<img src="/assets/images/hello_dgp.png">
```

`public/templates/`
This is where the templates are located.