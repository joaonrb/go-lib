[![Quality](https://github.com/joaonrb/go-lib/actions/workflows/quality.yml/badge.svg)](https://github.com/joaonrb/go-lib/actions/workflows/quality.yml)
# go-lib
Helpful go library. It contains the most common methods used by me in most of 
my projects. I included this in this library, so I don't have to re-write them
every time I work on a new project.

## Modules

### types
A modules with support types, most notably, my implementation of the **result**
and the **option** monads in go. 

### convert and converto
Functions to convert values to pointers, ponters to values or conversion between
monads.

### errors
Wrappers to add stack to errors.

### log
Simple log framework that allows to customise format and input.

### rand
Fast string generator.

### queue
A first in first out thread safe queue