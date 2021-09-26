# AWSCTL

[![Go Report Card](https://goreportcard.com/badge/github.com/chloe-codes1/awsctl)](https://goreportcard.com/report/github.com/chloe-codes1/awsctl) [![Coverage Status](https://coveralls.io/repos/github/chloe-codes1/awsctl/badge.svg?branch=main)](https://coveralls.io/github/chloe-codes1/awsctl?branch=main)

> 👩🏻‍🔧 Manage your AWS resources easily and efficiently

<br>
<br>

```AsciiDoc
                         _   _ 
                        | | | |
  __ ___      _____  ___| |_| |
 / _` \ \ /\ / / __|/ __| __| |
| (_| |\ V  V /\__ \ (__| |_| |
 \__,_| \_/\_/ |___/\___|\__|_|
                               
                               
```


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
awsctl/internal/config/config.product.yml

# Development Environment
awsctl/internal/config/config.devel.yml
```

<br>

### 2. Build & Run project


#### 2-1. By using `Makefile`

Check available commands
```sh
make build run
```

Execute command

```sh
make build run cmd={COMMAND}
```
<br>

#### 2-2. By using CLI
Install the latest version of the library
```sh
go get -u github.com/chloe-codes1/awsctl
```
Run CLI
```sh
awsctl {COMMAND}
```

ex)
```sh
awsctl get-lambda-with-vpc
```

<br>  