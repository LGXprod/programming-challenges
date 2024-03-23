# Problem 1: Multiples of 3 or 5

[Source](https://projecteuler.net/problem=1)

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

### Versions

- [V1](./v1): Version one is a simple, synchronous approach where the multiples of 3 and 5 are summed sequentially until the multiples exceed 1000 ✅

Mean execution time: 1150 nanoseconds

- [V2](./v2): Version two uses multiple threads to calculate the multiples asynchronously ❌

Mean execution time: NA
