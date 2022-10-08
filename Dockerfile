FROM golang:1.19-alpine
WORKDIR /zup

# Download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Build app
COPY . .
RUN go build -o /zup-message-service

# Set run command
EXPOSE 8081
CMD [ "/zup-message-service" ]