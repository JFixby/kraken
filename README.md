# kraken 
Release 0.0.1

This is not an application. This is a plugable go-lang component (library).

Packages:
### input 
Reads input data and transforms it into data stream
sending events to the ```DataListener```.

See ```input_test.go``` for usage example.

### output 
Reads output TEST data and keeps it for the following testing.

Warning: Make sure scenario names in output file match the names in the input file.
For example ```scenario 1``` in input corresponds to ```scenario 1``` and not to
```scenario  1``` (two spaces)

See ```output_test.go``` for usage example.

### util
Contains ```SkipList``` for log N data access

### install
Contains setup, deployment and installation scripts. No need to use docker.
Simple ```git``` is enough.

### orderbook
Contains the main class ```Book```.
```Book``` is a plugable go-lang component (library).
Book listens to input order events and spawns
output events sent to ```BookListener```

Usage example and the main test is located in the ```orderbook_test.go```
The test 
1) reads expected output values.
2) reads input values and redirects it to order book
3) outputs spawned by order book are checked against test output

input -> test setup -> ```Book``` -> test setup, checked against expected output