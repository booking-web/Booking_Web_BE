# BOOKING_WEB_BE

## Overview
This is the repository for Back-End of JioHealth.

**Link Project** : [`booking.web`](https://github.com/booking-web)

**Author** : [Bill Zay 😉](https://github.com/billzayy)

## How to run
To run this project, you just need to use this command on your terminal:

```bash
docker-compose -f ./deploy/docker-compose.yml up -d
```

<!-- ***<span style="color:red;">P/S</span>*** : Before run this project, you need to contact with author or create a `.env` file inside `/internal` with this template : -->

<!-- ```properties
PORT=<port-number>
DB_PATH="postgresql://<username>:<password>@<ip>:<port>/<database-name>?sslmode=disable"
``` -->

## Go Structure
### `/api`
This folder is use for Swagger specs and JSON schema files.
### `/assets`
Other assets to go along with this project (images, logos, etc).
### `/build`
Packaging and Continuous Integration.

container (Dockerfile) package configurations and scripts are put in this directory.
### `/cmd`
Main applications for this project.

The directory name for each application should match the name of the executable you want to have (e.g.,`/cmd/main`).
### `/deploy`
`docker-compose.yml` file and some folders of each container will volume on this directory.
### `/docs`
Design and user documents [`docs`](docs/ReadMe.md)
### `/internal`
Private application and library code. This is store the project source code.

This directory is the same of `/src` in the [Front-End](https://github.com/booking-web/Booking_Web_FE) directory

## REFERENCES
* Golang Standard Project Layout : [Click here !](https://github.com/golang-standards/project-layout)
* Basic Syntax Markdown : [Click here !](https://docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax#links)

---
<br>

$${\color{lightgreen}HAPPY \space CODING \space 😉 \space !!!}$$	
