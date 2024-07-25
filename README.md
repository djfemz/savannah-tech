
# Project Title
SavannahTech


## Environment Variables

To run this project, you will need to add the following environment variables to your example.env file

`DATABASE_HOST`
`DATABASE_USERNAME`
`DATABASE_PASSWORD`
`DATABASE_NAME`
`DATABASE_PORT`
`REPO_OWNER`
`REPO_NAME`



## Getting Started
To Run the project, navigate to the projects root directory and execute the following commands:
1. 'docker-compose up --build' to build the application.


## Testing
- The test suites for all the components of the application are located in the controllers, repositories and services directories in the api directory. To run the tests, run the command:
- go test .\services .\repositories .\controllers
## Documentation
1. The project is documented using swagger.io and can be accessed by running the application and visiting <app_base_url:PORT>/swagger/index.html in a web browser e.g localhost:8080/swagger/index.html.
[Documentation](https://<app_base_url>/swagger/index.html)

