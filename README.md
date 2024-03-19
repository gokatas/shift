LEVEL: intermediate

TOPICS: design, cryptography, testing

PROBLEM: Alice and Bob want to exchange messages. They don't want Eve to be able to read the messages.

SOLUTION: Messages have to be encrypted when transferred. We'll use simple shift cipher to do this. It encrypts by adding a key to each plaintext message byte. It decrypts by substracting a key from each ciphertext message byte.

IMPLEMENTATION: Start implementing the solution by writing tests first (TDD).

SOURCE: Adapted from https://github.com/bitfield/eg-crypto.