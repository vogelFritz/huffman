# Huffman

## Execution

huffman {-c|-d} original result

- -d flag to decompress.
- -c flag to compress.

## Algorithm

The huffman encoding algorithm creates a binary tree that stores the source alphabet and the corresponding code word for each symbol.
To get the code word for a specific symbol, accumulate all the code symbols in the path to the source symbol.

The usual implementation uses a heap to create the tree.

This tree must be stored in the compressed file to be able to decompress it later.
