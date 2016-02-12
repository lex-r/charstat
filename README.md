# Charstat

Charstat takes text as input arguments and calculates a links between symbols.
It shows a graph as the dot language.

### Usage

```sh
$ ./charstat "A weighted graph is a graph in which a number is assigned to each edge" > graph.dot
$ dot -Tsvg graph.dot > graph.svg
```

![](https://lex-r.github.io/storage/graph.svg)
