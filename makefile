# Set the project name
PROJECT_NAME = vcardlist

# Set the working directory
WORKDIR = $(shell pwd)

# Create a build directory
BUILD_DIR = $(WORKDIR)/build

# Create a release directory
RELEASE_DIR = $(WORKDIR)/release

# Create a temporary directory
TEMP_DIR = $(WORKDIR)/tmp

# Build the project
build:
	go build -v -o $(BUILD_DIR)/$(PROJECT_NAME)

# Clean up
clean:
	rm -rf $(BUILD_DIR)
	rm -rf $(TEMP_DIR)

# Run the tests
test:
	go test -v ./...
