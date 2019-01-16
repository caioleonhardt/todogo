This is is a repository to test a simple go programming using 
- go modules
- multistage builds

## Build the container with the code
    docker build -t todo:0.1 .

## removes all the intermediate builder images
    docker rmi -f $(docker images -q -f label=stage=intermediate)

## Run the app
    docker run --rm -v ${PWD}:/tmp  todo:0.1