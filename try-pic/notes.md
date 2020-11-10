
Notes and gotchas
-----------------

install 'pic' (was already installed?)
install 'pic2graph' (via 'groff'?)
install 'imagemagick` (was already installed)

#### pic2graph usage

Input to pic2graph needs to skip '.PS' and '.PE' tags.

#### `operation not allowed` error

ImageMagick's `convert` is configured to forbid conversion of PS files by
default, resulting in an error like this from `convert`:

```
convert-im6.q16: attempt to perform an operation not allowed by the security policy `PS' @ error/constitute.c/IsCoderAuthorized/408.
convert-im6.q16: no images defined `/tmp/pic2graph-xRIihz/pic2graph.png' @ error/convert.c/ConvertImageCommand/3258.
```

Needed to edit `/etc/ImageMagick-6/policy.xml` in order to fix that. Find a line like this:

```
<policy domain="coder" rights="none" pattern="PS" />
```

and change it to say `rights="read | write"`.

