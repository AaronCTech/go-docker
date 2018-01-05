# go-docker
An example (that I would have loved early on...) showing some of the most common use-cases for Go and Docker.

## Prerequisites

 * (https://golang.org/dl) Go is a little necessary :P
 * (https://store.docker.com/editions/community/docker-ce-desktop-windows) Docker is the way, the truth, and the light.

## Getting Started

run "git clone https://github.com/AaronCTech/go-docker.git" in your favorite project directory on your computer to download the project. GitHub will also allow a direct download via a zip file from the project.

## Running the application

To run the application, open your favorite command line tool (CMD/Powershell/Bash), enter the go-api directory, and run "go run main.go". This will invoke Go's built in run tool. Visit localhost:8285 in your browser, and behold!

## Building an executable

The first thing we need to do to build an ultra-small container, containing ONLY our executable, is to set a couple of flags in Go's environment to indicate that we want to build for Linux.

 * set GOOS=linux
 * set CGO_ENABLED=0

After setting our environment variables, we can now build the binary! The command to do this without CGO is slightly more complicated than a vanilla build, but if you script this command, you'll never have to think about it again :)

From the go-docker directory, run:

 * go build -a -installsuffix cgo -o main .

You will see a "main" file appear. This is the file that we will place into our container using the Dockerfile!

## Building the docker container

Now we need to create the docker container. Run the following command from within the go-docker directory:

 * docker build -t "go-docker" .

This will create a new docker image, tagging it as "go-docker".

## Deploying the docker container

We're ready to deploy this container onto our system! To do that, we will use docker-compose, which will load the docker-compose.yml file, and run our container.

From the go-docker directory, run:

 * docker-compose up -d

This will bring the container to life, expose our port 8285 to the world, and detach us from the shell so it can run without our interaction. To attach to the shell for logs, etc, run the command without "-d".

## Testing it out

Visit localhost:8285 from your browser, and you should see some text welcoming you to your new Go Docker app!