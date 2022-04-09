# multihash

Multihash is a console utility to calculate hashes recursively.

# Output Example

`multihash .`

```
# ./multihash.go
size:1851 mtime:2022-09-04  ./multihash.go
MD5:c6796f2b4473cde241e586a5ba10ed93  ./multihash.go
SHA1:c13ebeccc09d8786f0ceafc7a57ecdb2527800f8  ./multihash.go
SHA256:ac93a985548c681a8623c2112d860ea528899b82f4e7c58bdefe5ad89caa8205  ./multihash.go
SHA512:f7d903d593a9b46b5451d82944fa0f2d52f1586c335153bcd3d06eb32ecbb118f5eba1ceba47f76f5fe16c1744861cf5d4686ec186602dacdb574a369fc76af3  ./multihash.go

# ./multihash_test.go
size:1236 mtime:2022-09-04  ./multihash_test.go
MD5:69574fc6a501dc0842e19e6148438592  ./multihash_test.go
SHA1:02d88773eec877f41f1e1b9fcd9d2e46967bc493  ./multihash_test.go
SHA256:ac4461c0f4036e8214c1dc66949399d61558520a4c7b4b6bbe76422303d36639  ./multihash_test.go
SHA512:e85d93693bbc9154a9c4283336a01a1e7555aec875a18ebe6b6c5e672336100ead31fa5c93b458006d9841d3c7c7e662594718d93c2e5d4281e6e26e090dec14  ./multihash_test.go

# ./README.md
size:301 mtime:2022-09-04  ./README.md
MD5:24c7f3349d17fd20bb50354eae5bbae4  ./README.md
SHA1:8115de61d2d12239356c61a1d5fd43d7404335f1  ./README.md
SHA256:a0be8b5b273797fe93875fa4fd6518f6ed166f37ee8b63a0b2bd242aee4c7e04  ./README.md
SHA512:a9cc228401772419989ff6cbbd953aa41b84ad2e54ec26ee2b4c21b7dc6bf1f0451c1c754b0bb4dd4f667caa0bb5f03e814f3e6b0d2ddf7291e650e94bf33496  ./README.md

# ./go.mod
size:45 mtime:2022-09-04  ./go.mod
MD5:fd5dfa28c810cc22b3443f963f37817e  ./go.mod
SHA1:0f077c393ef8bd3f0b8ba8aa9f45dadfcbc4a694  ./go.mod
SHA256:6058db165e4efc4e5d1901a7800c3fb6d6f02d1608b111eaf257a4a35a7b76e5  ./go.mod
SHA512:b81ed65ff0832a7a15a08d4a12b0a2e13aa69320cbc92a0fa51b647d9d80c4347e9e6988cb66ceeaa0d1ef24dc6c04fcd1089b42c3b21ef8db55a383ffaf8df4  ./go.mod

# ./COPYING
size:35147 mtime:2022-09-04  ./COPYING
MD5:d32239bcb673463ab874e80d47fae504  ./COPYING
SHA1:8624bcdae55baeef00cd11d5dfcfa60f68710a02  ./COPYING
SHA256:8ceb4b9ee5adedde47b31e975c1d90c73ad27b6b165a1dcd80c7c545eb65b903  ./COPYING
SHA512:7633623b66b5e686bb94dd96a7cdb5a7e5ee00e87004fab416a5610d59c62badaf512a2e26e34e2455b7ed6b76690d2cd47464836d7d85d78b51d50f7e933d5c  ./COPYING

# 443.909µs runtime

```

# Verify
You can wrote the output to a file and can do something like this
`sed -n -e 's/^SHA512://p' YOUR_FILE | sha512sum -c`