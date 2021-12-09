### Piping standard input to stdandard output, rate limiting.

Read standard input and rate limit to standard output if -kbps is specified.


Example:
```
dd if=/dev/zero bs=1M | ./catlimit_linux_amd64 -kbps $((1024*2)) | pv -per
```

will limit to 2MiB/s.

