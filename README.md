# go-datastructures

### augmentedtree

This is an implemnetation of an interval tree for collision in n-dimensional ranges. It is a fork of [augmentedtree package](https://github.com/Workiva/go-datastructures) from Workiva with one modification: it supports intervals with integers of arbitrary size using Go's [math/big](https://golang.org/pkg/math/big/) package.

This was created to handle IPv4 and IPv6 range lookups for [DNSAuth](https://github.com/Packet-Clearing-House/DNSAuth).
