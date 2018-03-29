# Task Manager RESTful API

This TaskManager RESTful API will be a JSON-based API in which the JSON data format will be used for sending and receiving messages among client applications and the RESTful API server.

## Third-Party Packages

The following third-party packages will be used for building the TaskManager application:

* [gopkg.in/mgo.v2](gopkg.in/mgo.v2): The mgo package provides a MongoDB driver for Go.
* [gopkg.in/mgo.v2/bson](gopkg.in/mgo.v2/bson): A subpackage of mgo, it provides implementation of the
  BSON specification for Go.
* [github.com/gorilla/mux](github.com/gorilla/mux): The mux package implements a request router and dispatcher.
* [github.com/dgrijalva/jwt-go](github.com/dgrijalva/jwt-go): The jwt-go package implements helper functions for working with JSON Web Tokens (JWT).
* [github.com/codegangsta/negroni](github.com/codegangsta/negroni): The Negroni package provides an idiomatic approach to HTTP middleware.

## Application Structure

The RESTful API application has been divided into the following packages:

* **common**: Implements some utility functions and provides initialization logic for the application
* **controllers**: Implements the application’s application handlers
* **data**: Implements the persistence logic with the MongoDB database
* **models**: Describes the data model of the application
* **routers**: Implements the HTTP request routers for the RESTful API
