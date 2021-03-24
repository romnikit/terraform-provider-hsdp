# hsdp_function

Define function-as-a-service using various backends. Currently
only `iron` is supported. 

## Example usage

```hcl
resource "hsdp_function" "rds_backup" {
  name = "streaming-backup"
  
  # The docker packaged function business logic
  docker_image = var.streaming_backup_image
  docker_credentails = {
    username = var.docker_username
    password = var.docker_password
  }
  
  # Environment variables available in the container
  environment = {
    db_name = "hsdp_pg"
    db_host = "rds.aws.com"
    db_username = var.db_username
    db_password = var.db_password
    
    s3_access_key = "AAA"
    s3_secret_key = "BBB"
    s3_bucket = "cf-s3-xxx"
    s3_prefix = "/backups"
  }

  # Run every day at 4am
  schedule {
    start = "2021-01-01T04:00:00Z"
    run_every = "1d"
  }

  backend {
    type = "iron"
    credentials = module.iron_backend.credentials
  }  
}
```

## Argument reference
The following arguments are supported:

* `name` - (Required) The name of the function
* `docker_image` - (Required) The docker image that contains the logic of the function
* `docker_credentials` - (Optional) The docker registry credentials
  * `username` - (Required) The registry username
  * `password` - (Required) The registry password  
* `command` - (Optional) The command to execute in the container. Default is `/app/server`
* `environment` - (Optional, map) The environment variables to set in the docker container before executing the function
* `schedule` - (Optional) Schedule the function. When not set, the function becomes a task.
  * `start` - (Optional, RFC3339) When to start the schedule. Example: `2021-01-01T04:00:00Z`. Default is a date in the past.
  Setting the start argument allows you to control the specific time within the day when your task will be scheduled. 
  * `run_every` - (Required) Run the function every `{value}{unit}` period. Supported units are `s`, `m`, `h`, `d` for second, minute, hours, days respectively.
    Example: a value of `"20m"` would run the function every 20 minutes.
* `backend` - (Required) The backend to use for scheduling your functions.
  * `type` - (Required) The backend type. Only `iron` is supported at this time.
  * `credentials` - (Required, map) The backend credentials. Must be iron configuration details at this time.
    
## Attribute reference