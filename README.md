# multihash

Multihash is a console utility to calculate hashes recursively.

# Output Example

`multihash .`

```
# ./multihash.go
size:2104 mtime:22:13:25 2022-09-04  ./multihash.go
CRC32:2493743490  ./multihash.go
MD5:58061db01a4670cddafea9686773d757  ./multihash.go
SHA1:91719bd0bb13c4b6a7a46c09e1499ecb74ffa151  ./multihash.go
SHA256:b8907c9cd6cebb9f654e39b7c5c2f1d91368c17a823650b58a4e9d83af101da2  ./multihash.go
SHA512:2c5b72a3ee3573be218fbfd5dcc6f37201e0197181a2e9d957ca72e4cb89cd490355e02b6952ea8b85e02d0d812b0447e8e0caa2911d95b08fa63096c043558d  ./multihash.go

# ./multihash_test.go
size:1476 mtime:22:13:05 2022-09-04  ./multihash_test.go
CRC32:1201092340  ./multihash_test.go
MD5:cbad743dfa2202574153fe3f9266a5c4  ./multihash_test.go
SHA1:4a5e3ebf767f2a06e7ffb442b19280cc1d54b167  ./multihash_test.go
SHA256:2570b9a0f112c506ca7381171b85f6c9a9332a70fe1b22b645570a9f31538a08  ./multihash_test.go
SHA512:43a0a6143364a7beefde49be4ba633dbaa2e47e17eaca2368b3921d0bc4463458f0f22a6e7f19d144d8b4ab9814d2425c9049dee233a98d164876661e447bbae  ./multihash_test.go

# ./README.md
size:2464 mtime:22:13:51 2022-09-04  ./README.md
CRC32:1137677913  ./README.md
MD5:ee913a20538cb6230985c86956c871c2  ./README.md
SHA1:50c332e0862a2976f22bf17b7e2b35729c7db3cc  ./README.md
SHA256:80f6d3038f575049dfaf7d3eed6d2e2440a701e66970552f1bc34445075bcb8b  ./README.md
SHA512:3dbf34e1dc54274d24f0034b1f3a0b580eabd4f8ceec4729cc6f9a8e9c2267e06f36836570dc3dc128cc701a3e65962743dbbdfe98df8c3afe730980eac9d214  ./README.md

# ./go.mod
size:45 mtime:21:29:15 2022-09-04  ./go.mod
CRC32:4264533554  ./go.mod
MD5:fd5dfa28c810cc22b3443f963f37817e  ./go.mod
SHA1:0f077c393ef8bd3f0b8ba8aa9f45dadfcbc4a694  ./go.mod
SHA256:6058db165e4efc4e5d1901a7800c3fb6d6f02d1608b111eaf257a4a35a7b76e5  ./go.mod
SHA512:b81ed65ff0832a7a15a08d4a12b0a2e13aa69320cbc92a0fa51b647d9d80c4347e9e6988cb66ceeaa0d1ef24dc6c04fcd1089b42c3b21ef8db55a383ffaf8df4  ./go.mod

# ./COPYING
size:35147 mtime:20:52:53 2022-09-04  ./COPYING
CRC32:1719137660  ./COPYING
MD5:d32239bcb673463ab874e80d47fae504  ./COPYING
SHA1:8624bcdae55baeef00cd11d5dfcfa60f68710a02  ./COPYING
SHA256:8ceb4b9ee5adedde47b31e975c1d90c73ad27b6b165a1dcd80c7c545eb65b903  ./COPYING
SHA512:7633623b66b5e686bb94dd96a7cdb5a7e5ee00e87004fab416a5610d59c62badaf512a2e26e34e2455b7ed6b76690d2cd47464836d7d85d78b51d50f7e933d5c  ./COPYING
```

# Verify
You can wrote the output to a file and can do something like this
`sed -n -e 's/^SHA512://p' YOUR_FILE | sha512sum -c`