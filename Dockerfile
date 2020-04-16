FROM golang:buster
WORKDIR /usr/src/app
COPY . .
RUN ./installdeps.sh
RUN go build main.go
EXPOSE 8080
CMD main