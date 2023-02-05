### What is it?
A CLI tool to store secrets. Stores key, value pairs of which value is encrypted.
Value can be accessed by supplying a decryption key.

Keys can form a group called vault. 




### Learnings
- Using AES and md5
- Filling random data in empty slice
- Using `io.Reader` and `io.Writer` to simplify and future proof operations. Look at `vault.load()` and `vault.save()`.

