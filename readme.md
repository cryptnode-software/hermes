# Hermes
Hermes is an application that will take care of all of your messaging needs. Similar to an event based system it will dispatch any of the configured actions after receiving an even. If you need something as simple server that keeps track of a commenting system, Hermes will handle it. If you need something that will dispatch actions that go over multiple streams/subscriptions. Hermes will quickly and efficiently handle it for you. 

## Documentation
Although this is a work in progress and shouldn't be used in production. You can read the [documentation](https://docs.hermes.cryptnode.tech) to start playing with it.  

## Environment
Below is a list of environment variables to get Hermes up and running. At cryptnode we use envrc in order to add hook environment variables into our shell on a application level basis. You can use any method that you may prefer.
* database related
    * SUPER_USER_CONNECTION (only required to create initially hermes db)
    * DB_CONNECTION (default 127.0.0.1:3306)