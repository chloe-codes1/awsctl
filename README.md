# AWS Manager

> 👩🏻‍🔧 Manage your AWS resources easily and efficiently

<br>
<br>


## Features

### Lambda
- [x] List Functions
  - Returns only the first 50 lambdas
- [x] List Functions Pages
  - Return all lambdas
- [x] Get Function
  - Check if the function is in 
- [x] Update Function Configuration

<br>

### Auto Scaling Groups
- [x] List Auto Scaling Groups
  - Return all ASGs
- [x] List ASG in specific subnet


<br>
<br>

## How to run

### 1. Fill out config files

Config files are located in the path below

```sh
# Production Environment
asg-manager/internal/config/config.product.yml

# Development Environment
asg-manager/internal/config/config.devel.yml
```

<br>

### 2. Build & Run project
```sh
make build run
```
