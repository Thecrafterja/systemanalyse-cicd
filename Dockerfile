FROM golang:1.24

# sets current directory in container's filesystem
WORKDIR /app

COPY . .

# At buildtime
RUN CGO_ENABLED=0 GOOS=linux go build -o ./compilexample

# At runtime
CMD [ "./compilexample" ]