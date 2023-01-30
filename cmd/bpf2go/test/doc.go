// Package test checks that the code generated by bpf2go conforms to a
// specific API.
package test

// $BPF_CLANG and $BPF_CFLAGS are set by the Makefile.
//go:generate go run github.com/Felixxxlz/ebpf/cmd/bpf2go -cc $BPF_CLANG test ../testdata/minimal.c
