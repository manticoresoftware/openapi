#!/usr/bin/env python3
"""Generate manticore_no_nullables.yml from manticore.yml for the C# client."""

import re
import sys


def main() -> None:
    src, dst = sys.argv[1], sys.argv[2]
    content = open(src, encoding="utf-8").read()
    content = content.replace("type: {}", "type: object")
    content = re.sub(
        r"^([ \t]+)(\S+:\n)\1  oneOf:\n\1    - type: null\n\1    - (\$ref: [^\n]+)$",
        r"\1\2\1  \3",
        content,
        flags=re.MULTILINE,
    )
    open(dst, "w", encoding="utf-8").write(content)


if __name__ == "__main__":
    if len(sys.argv) != 3:
        sys.exit(f"usage: {sys.argv[0]} <input.yml> <output.yml>")
    main()
