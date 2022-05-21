FROM golang:1.18-alpine AS build

WORKDIR /app
COPY go.* ./
RUN go mod tidy

COPY . .
RUN go build -o bot

FROM alpine
COPY --from=build ./app/bot ./
ENV TRANSLIT_TOKEN=""
CMD ["./bot"]