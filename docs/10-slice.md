# Slice

* Slice is a referenced data type, as it references an underlying array.
* It uses array underneath, its a abstract layer on top array.
* The name slice is given as its a slice from an underlying array.
* Size of slice can be increased or decreased dynamically.
* Slice always reference an array (its stores address of an array). Slice also contain the length to be used and the capacity.
* Length should be less than equal to capacity.
* **Length** is the number of element the slice have access to.
* **Capacity** is the number of element the slice can grow.
* Slice itself are very small in size, and it can passed to function easily as pass by value. Unlike array slice itself do not carry any data.
* Multiple slice can reference to same array and may have same or different view. This is based on the creation of the slice.
* Declaration and initialization of slice is very similar. Refer live code.
* Append function will create a new slice if current slice capacity is exhausted. of double size.

> NOTE: as you can see multiple slice can refer to same array and can modify the data. Take extra care to recreate the slice when required.

## Refer

009-slice
