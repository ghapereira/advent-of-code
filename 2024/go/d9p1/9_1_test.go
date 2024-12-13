package main

import "testing"

func TestChecksum(t *testing.T) {
	p := new(Problem)
	p.Initialize("2333133121414131402")
	p.Defrag()
	have := p.Checksum()
	want := int64(1928)

	if have != want {
		t.Fatalf("Checksum = %v, want %v", have, want)
	}
}
