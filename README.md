# multihash

Multihash is a console utility to calculate hashes recursively.

# Output Example

`multihash .`

```
# ./multihash.go
size:1925 mtime:21:55:32 2022-09-04  ./multihash.go
MD5:658a3e7770006d647c47505f95b0b14c  ./multihash.go
SHA1:04e1d580433ab43149400f51f439927e3dd1af90  ./multihash.go
SHA256:a18104b8c06fb74c16d0e59040fb39f6ed5cf66f53fe65b9dbb992a5480b0e66  ./multihash.go
SHA512:e79769dcc314e9ad32e5b3d4c97e50532a1e845bca004d031c492c1bd655052cec640d4b3b97d1dfb6cd39e16e41bd66b3032760ec282e7a908480f84941587d  ./multihash.go

# ./multihash_test.go
size:1236 mtime:20:54:21 2022-09-04  ./multihash_test.go
MD5:69574fc6a501dc0842e19e6148438592  ./multihash_test.go
SHA1:02d88773eec877f41f1e1b9fcd9d2e46967bc493  ./multihash_test.go
SHA256:ac4461c0f4036e8214c1dc66949399d61558520a4c7b4b6bbe76422303d36639  ./multihash_test.go
SHA512:e85d93693bbc9154a9c4283336a01a1e7555aec875a18ebe6b6c5e672336100ead31fa5c93b458006d9841d3c7c7e662594718d93c2e5d4281e6e26e090dec14  ./multihash_test.go

# ./README.md
size:2293 mtime:21:32:13 2022-09-04  ./README.md
MD5:4cf5d9f4e6386eea7f73e07a6608e3be  ./README.md
SHA1:db789c2a9d93f67871ec7524a889664e238a4530  ./README.md
SHA256:511da0b8fa85f140447329089151b9f52feed9560bb7b91fd74c43b2413d7c25  ./README.md
SHA512:d85a56444da9a8102122a8ff9c8d2d4b56042377414e2e86efb123538e9477f8b9cf395b402cc4c261c1e0017c33a34ef28c11cf56af79a226468622c56cf386  ./README.md

# ./go.mod
size:45 mtime:21:29:15 2022-09-04  ./go.mod
MD5:fd5dfa28c810cc22b3443f963f37817e  ./go.mod
SHA1:0f077c393ef8bd3f0b8ba8aa9f45dadfcbc4a694  ./go.mod
SHA256:6058db165e4efc4e5d1901a7800c3fb6d6f02d1608b111eaf257a4a35a7b76e5  ./go.mod
SHA512:b81ed65ff0832a7a15a08d4a12b0a2e13aa69320cbc92a0fa51b647d9d80c4347e9e6988cb66ceeaa0d1ef24dc6c04fcd1089b42c3b21ef8db55a383ffaf8df4  ./go.mod

# ./COPYING
size:35147 mtime:20:52:53 2022-09-04  ./COPYING
MD5:d32239bcb673463ab874e80d47fae504  ./COPYING
SHA1:8624bcdae55baeef00cd11d5dfcfa60f68710a02  ./COPYING
SHA256:8ceb4b9ee5adedde47b31e975c1d90c73ad27b6b165a1dcd80c7c545eb65b903  ./COPYING
SHA512:7633623b66b5e686bb94dd96a7cdb5a7e5ee00e87004fab416a5610d59c62badaf512a2e26e34e2455b7ed6b76690d2cd47464836d7d85d78b51d50f7e933d5c  ./COPYING

# 443.909µs runtime
```

# Verify
You can wrote the output to a file and can do something like this
`sed -n -e 's/^SHA512://p' YOUR_FILE | sha512sum -c`