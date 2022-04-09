# multihash

Multihash is a console utility to calculate hashes recursively.

# Output Example

`multihash .`

```
# ./multihash.go
size:2104 mtime:22:02:30 2022-09-04  ./multihash.go
CRC32:ed215a71  ./multihash.go
MD5:7d8ce2d5c1307680b48afe89c5278aa7  ./multihash.go
SHA1:62846f4d8beca8abe513b7cea852b7eea3bbe00d  ./multihash.go
SHA256:d189b94a8579ad1e26897e8751736f2b1860e832fb1b2ac30cd585024d2dd7f5  ./multihash.go
SHA512:03dde86ee563905ceb35d6fb5794f987c911ba9d21c544736081e46a050816395b51b82414ed55c57eea610fc1146af5d27f5cffef456fd6b46b43859b6a521d  ./multihash.go

# ./multihash_test.go
size:1236 mtime:20:54:21 2022-09-04  ./multihash_test.go
CRC32:8178964d  ./multihash_test.go
MD5:69574fc6a501dc0842e19e6148438592  ./multihash_test.go
SHA1:02d88773eec877f41f1e1b9fcd9d2e46967bc493  ./multihash_test.go
SHA256:ac4461c0f4036e8214c1dc66949399d61558520a4c7b4b6bbe76422303d36639  ./multihash_test.go
SHA512:e85d93693bbc9154a9c4283336a01a1e7555aec875a18ebe6b6c5e672336100ead31fa5c93b458006d9841d3c7c7e662594718d93c2e5d4281e6e26e090dec14  ./multihash_test.go

# ./README.md
size:2338 mtime:21:57:17 2022-09-04  ./README.md
CRC32:7a6fe403  ./README.md
MD5:0c0bd139d30bdf7015ad2aaae1404f91  ./README.md
SHA1:977008c155d0827751beacdae119c7e5ffaf62e1  ./README.md
SHA256:790a694c69ea07aec86cb219892b4c0609116635c060d426e5646e2d15665670  ./README.md
SHA512:fd5a262bf107d6950b350736ddb376c06b742468fb2b25e947335486259b2e516a934d917407887a36c306cfe2e553a0d2b1804fdf25e366099389126fa4200c  ./README.md

# ./go.mod
size:45 mtime:21:29:15 2022-09-04  ./go.mod
CRC32:fe2f9e32  ./go.mod
MD5:fd5dfa28c810cc22b3443f963f37817e  ./go.mod
SHA1:0f077c393ef8bd3f0b8ba8aa9f45dadfcbc4a694  ./go.mod
SHA256:6058db165e4efc4e5d1901a7800c3fb6d6f02d1608b111eaf257a4a35a7b76e5  ./go.mod
SHA512:b81ed65ff0832a7a15a08d4a12b0a2e13aa69320cbc92a0fa51b647d9d80c4347e9e6988cb66ceeaa0d1ef24dc6c04fcd1089b42c3b21ef8db55a383ffaf8df4  ./go.mod

# ./COPYING
size:35147 mtime:20:52:53 2022-09-04  ./COPYING
CRC32:6677f57c  ./COPYING
MD5:d32239bcb673463ab874e80d47fae504  ./COPYING
SHA1:8624bcdae55baeef00cd11d5dfcfa60f68710a02  ./COPYING
SHA256:8ceb4b9ee5adedde47b31e975c1d90c73ad27b6b165a1dcd80c7c545eb65b903  ./COPYING
SHA512:7633623b66b5e686bb94dd96a7cdb5a7e5ee00e87004fab416a5610d59c62badaf512a2e26e34e2455b7ed6b76690d2cd47464836d7d85d78b51d50f7e933d5c  ./COPYING
```

# Verify
You can wrote the output to a file and can do something like this
`sed -n -e 's/^SHA512://p' YOUR_FILE | sha512sum -c`