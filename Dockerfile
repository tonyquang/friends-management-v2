FROM golang

WORKDIR /friends_management_v2

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 3000

RUN go build

CMD ["./friends_management_v2"]