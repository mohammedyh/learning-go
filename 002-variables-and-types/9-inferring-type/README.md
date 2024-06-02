# Inferring Type

There is a way to declare a variable without explicitly stating its type using the 'short declaration' `:=` operator. We might use the `:=` operator if we know what value we want our variable to store when creating it.

Go also offers a separate syntax to declare a variable and infer its type using the `var` keyword, but without specifying a type.

Floats created using the short declaration operator are of type `float64`. Integers created in this way are either `int32` or `int64` depending on the computers architecture (32 or 64 bit)

Computers actually have a default length for the data in their Read-Only Memory (ROM). Some newer computers may have more processing power and can store/handle bigger chunks of data. These computers might be using a 64-bit architecture, but other computers still run on 32-bit architecture.

By providing the type `int` or `uint`, Go will check to see if the computer architecture is 32-bit or 64-bit. Then it will either provide a 32-bit `int` (or `uint`) or a 64-bit one.

It’s recommended to use `int` unless there’s a reason to specify the size of the `int` (like knowing that value will be larger than the default type, or needing to optimize the amount of space used).
