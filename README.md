# SplitBBQ

Like Splitwise, but from the command line.

# Usage

The CLI reads spendings from stdin. Each spending should be in its own line and follow the format *"name amount"*.
The name can include spaces and the amount can either be a float or integer number.

```bash
$ splitbbq
Ava Max 30.5
Eric Clapton 59.5
Steve Vai 0
```

Or pass it the contents of a file:

```bash
$ splitbbq < spendings
```
