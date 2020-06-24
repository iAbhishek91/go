# types

Go static-type language, hence ever variable should have type.

## Built-in types

* Built-in data types are int, float32, string, boolean, byte etc...They are shipped with the programming language.
* NOTE: Characteristic exhibited by this types are known as **Primitive nature**. Any type which do not follow these characteristic are  referred as **Non primitive nature**.
* All the build in types when we add or remove something from the value, a new value should be created. *This characteristics are known as Primitive nature*. Hence, we should pass by value(copy of the variable) to function or method for primitive type. All system libraries have been implemented in this way

## Reference types

* They are also built in the programming language, they are NOT basic types. All reference type points to underlying data structure.
* Reference type in Go: slice, map, channel, interface and function.
* Common in all these types are:
  * value they contain are almost of same size - called header
  * their value is never the actual value, its only reference to the actual value and some parameter to describe the type.
  * they are always passed as a copy and never as a pointer. Hence reference type are also **Primitive in nature**
* Refer slice, map, channel interface and function notes and live example.

## User defined types

* Only two ways to define a user defined type using: **struct** or **primitive types**.
* with each new types you can create new variables.
* you can perform data convert to new user-type type (IMPORTANT). Refer live code.


## Primitive and Non-Primitive nature

* When we change/update a value of any type (user defined), we need to understand **do we need to make a new copy of the type?** or **can we mutate the existing value?**

## Refer

04-variable-constant-type
