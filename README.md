# greenbone-task
A little coding challenge in go

## How to use this

1. Ensure you have a local postgresql instance running and know the port and connection settings for it
2. Create a database within your postgres instance which can be used for this service
3. Update the file `./config/service.env` to config your service:
    - Adjust `DATABASE` vars to point to the correct database (as mentioned in step 1)
    - Set `SERVICE` vars to your needs
4. Setup the service dependencies with:
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