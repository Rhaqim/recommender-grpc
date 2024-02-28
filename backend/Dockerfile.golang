FROM golangci/build-runner

# Set the working directory
WORKDIR /APP

# Copy the current directory contents into the container at /APP
COPY . /APP

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT [ "./main" ]