# https
**HTTP** **S**tream CLI

## Purpose

The purpose of `https` command is to
* Use HTTP/HTTPS streamingly
* Use HTTPS by default unlike `curl`

The last `s` of `https` has both meanings of **s**tream and HTTP**S** for usability of the flow and higher security.

The design policy is using the same option names as `curl`'s to relieve users' burden not to remember new things and not have confusion as much as possible.

## Usage

```bash
# GET
https get example.com
```

```bash
# POST
echo hello | https post example.com
```

### Same options as `curl`

Yon can access an HTTPS server disabling certification check like `curl -k`.

```bash
https get -k example.com
```

```bash
https get --insecure example.com
```

### Use HTTP

`https` command uses HTTPS by default. You can specify "http://" as follows to use HTTP.

```bash
# GET
https get http://example.com
```
