# Road Load

## Requirements

A section of road can sustain a given maximum weight and a number of two vehicles.  
If weight is too big the vehicle needs to return.  

## Task

Given a weight and a number of vehicle weights provide the number of returned vehicles.

## Example

Weight 9. Vehicle weights []uint{5, 3, 8, 1, 8, 7, 7, 6}. Returns 4.
Weight 7. Vehicle weights []uint{7, 6, 5, 2, 7, 4, 5, 4}. Returns 5.  
Weight 7. Vehicle weights []uint{3, 4, 3, 1}. Returns 0.  
