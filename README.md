# Hello Goa project

A companion project for a technical blog post. 
The project demonstrates the usage of [goa.design](https://goa.design) to generate a m/s. 

It contains branches named `v1`, `v2`,... which show the project as it evolves.

## Requirements 

* Go 1.19
* Make
* ssh-keygen
* OpenSSL

## Usage 

To generate the Goa scaffolding code  
`make generate`

To execute the tests  
`make test`

To generate an RSA keypair in `/tmp`
`make keys`

To run a local HTTP server  
` make run-local-http`

To build the CLI client  
`make build-cli`