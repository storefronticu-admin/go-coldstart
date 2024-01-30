# First stage: Build the binary
FROM golang:latest AS build

# Set the working directory in the container
WORKDIR /app

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Print the current directory and list files to help diagnose
RUN pwd && ls -la

RUN echo "APP_STAGE= STAGING" > app.env
# Build the binary
RUN CGO_ENABLED=0 go build -o main .

# Second stage: Copy the binary to a minimal image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/main .

COPY --from=build /app/app.env .
# Make the binary executable
RUN chmod +x main

# Expose the port that the application will run on
EXPOSE 8080

# Run the binary
ENTRYPOINT [ "./main" ]
