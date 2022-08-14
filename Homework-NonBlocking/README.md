# Generic Processes
## Process
A process can do generic work in `doWork` method.  
A process has a unique ID.

## HTTP Endpoints
Processes can be created as one or as many.
Processes can be created:
- syncronous, basically waiting for the process to finish work
- asyncronous, basically the process ID is returned but the work might not have finished.

## Cache
The process output would be always taken from cache.  
The cache can delay output delivery on processes that did not yet finish work.  
This can be seen in the below sequence of commands:
```sh
curl 127.0.0.1:8080/many/5/async/20000
curl 127.0.0.1:8080/cache/all   # repeat as needed
```

## Corners cut
Some code duplication.  
HTTP verb (using GET instead of POST).

