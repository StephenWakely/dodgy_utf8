# Dodgy Utf8

Simple utility to dump dodgy utf8 to various places so you can test that stuff properly handles Utf8.

Params:

- `--text <text>` - The text to dump to the place. The text is prepended with ♥ as a valid Utf8 sequence, at will have dodgy utf8 appended to the end.
- `--loop` - Add to dump one line per second to the place.
- `--stdout` - Dump to stdout.
- `--journald` - Dump to journald.
- ``--filename <filename>` - Dump to the specified file.
