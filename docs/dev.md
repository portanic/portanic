# Portanic Development Documentation

## Stack

Portanic is build using the following technologies:

* Go
* Echo
* Templ
* HTMX
* PostgreSQL

## Running portanic locally

At the moment the entire project can be setup with

* docker-compose build
* docker-compose up

These commands are also wrapped in the justfile, so you can instead do

* just build
* just up

When changes are made to a templ file we additionally need to run

* templ generate