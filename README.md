# greenbone-task

A little coding challenge in go. 

## Approach

While designing the solution for this, I wanted to follow entirely the API first approach which aims to first design 
the API specification using the OpenAPI standard and use this as source of truth for generating the models and handlers 
for the service. The reasoning behind this approach are nicely explained [here](https://medium.com/@bhanu.pratap/embracing-the-api-first-approach-development-practices-and-tools-for-efficient-product-development-6f0f0cc73049).

## How to use this

1. Ensure you have a local postgresql instance running and know the port and connection settings for it
2. Create a database within your postgres instance which can be used for this service
3. Ensure you have the admin notification service running. Do this via:
   ```shell script
   docker pull greenbone/exercise-admin-notification
   docker run -d -p 8080:8080 --name {your-container-name} greenbone/exercise-admin-notification
   ```
4. Update the file `./app.env` to config your service:
    - Adjust `DATABASE` vars to point to the correct database (as mentioned in step 1)
    - Set `SERVICE` vars to your needs
    - Ensure `NOTIFICATION_URL` is set to `http://localhost:8080/api/notify`
5. Setup the service dependencies with:
   ```shell script
   $ go mod tidy
   ```

## Build and run the application

1. **Build**

```shell script
make build
```

2. **Run**

```shell script
make run
```

3. **Test**

```shell script
make test
```

## Limitations

Due to time constraints, several tasks which I consider production ready
have been skipped for this task. Therefor the following items are missing in this solution:

* no tests (unit/integration/e2e)
* no auth for API
* improvable structure
* no docker build
* no CI/CD pipeline setup

## Frameworks and tools

1. Golang >= 18
2. [`gin`](https://github.com/gin-gonic/gin) for REST APIs
3. [`gorm`](https://gorm.io) as database object relation model
4. [`viper`](https://github.com/spf13/viper) for `.env` file configuration
5. [`zap`](https://github.com/uber-go/zap) for logging
6. [`oapi-codegen`](github.com/deepmap/oapi-codegen) for enabling API first development approach
