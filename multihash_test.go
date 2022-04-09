package main

import (
	"encoding/hex"
	"os"
	"testing"
)

func TestCRC32sum(t *testing.T) {
	// 6677F57C
	file, err := os.ReadFile("COPYING")
	if err != nil {
		t.Fatal(err)
	}

	var want uint32 = 1719137660
	got := crc32sum(file)

	if want != got {
		t.Fatalf("want: %d got: %d", want, got)
	}
}

func TestMd5sum(t *testing.T) {
	file, err := os.ReadFile("COPYING")
	if err != nil {
		t.Fatal(err)
	}

	want := "d32239bcb673463ab874e80d47fae504"
	got := hex.EncodeToString(md5sum(file))

	if want != got {
		t.Fatalf("want: %s got: %s", want, got)
	}
}

func TestSha1Sum(t *testing.T) {
	file, err := os.ReadFile("COPYING")
	if err != nil {
		t.Fatal(err)
	}

	want := "8624bcdae55baeef00cd11d5dfcfa60f68710a02"
	got := hex.EncodeToString(sha1sum(file))

	if want != got {
		t.Fatalf("want: %s got: %s", want, got)
	}
}

func TestSha256Sum(t *testing.T) {
	file, err := os.ReadFile("COPYING")
	if err != nil {
		t.Fatal(err)
	}

	want := "8ceb4b9ee5adedde47b31e975c1d90c73ad27b6b165a1dcd80c7c545eb65b903"
	got := hex.EncodeToString(sha256sum(file))

	if want != got {
		t.Fatalf("want: %s got: %s", want, got)
	}
}

func TestSha512Sum(t *testing.T) {
	file, err := os.ReadFile("COPYING")
	if err != nil {
		t.Fatal(err)
	}

	want := "7633623b66b5e686bb94dd96a7cdb5a7e5ee00e87004fab416a5610d59c62badaf512a2e26e34e2455b7ed6b76690d2cd47464836d7d85d78b51d50f7e933d5c"
	got := hex.EncodeToString(sha512sum(file))

	if want != got {
		t.Fatalf("want: %s got: %s", want, got)
	}
}
