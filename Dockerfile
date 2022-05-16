FROM golang:1.17.9

ENV NODE=production
ENV PORT=8080
ENV CONTEXT_PATH=/feedback-processor-api
ENV DB_HOST=postgres
ENV DB_NAME=unisuncmsdevdb
ENV DB_USER=urquhmotrdhwqg
ENV DB_PASS=efad4bb2169e67ddaa17c21aba5c76efc6a9daa6a06310949eba9a006bf258da
ENV DB_PORT=5432
ENV DB_SSL=disable
ENV DB_TIMEZONE=Asia/Bangkok
ENV GIN_MODE=release

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app .

EXPOSE 8080

CMD ["app"]