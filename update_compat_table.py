#!/usr/bin/env python3
"""Update the top 2x2 block of a client README compatibility table.

Inserts two rows for the new client/Manticore versions and rewrites the three
affected prior rows, preserving markdown column widths.

Usage:
  update_compat_table.py --client-version 11.2.0 --manticore-version 18.0.0 FILE [FILE ...]
"""

from __future__ import annotations

import argparse
import re
import sys
from pathlib import Path
from typing import List, Optional, Sequence, Tuple


OR_NEWER_RE = re.compile(r"^(v?\d+(?:\.\d+)*)\s+or newer$", re.I)
RANGE_RE = re.compile(
    r"^(v?\d+(?:\.\d+)*)\s+to\s+(v?\d+(?:\.\d+)*)$", re.I
)
MS_OR_NEWER_RE = re.compile(r"^(\d+(?:\.\d+)*)\s+or newer$")
MS_RANGE_RE = re.compile(r"^(\d+(?:\.\d+)*)\s+to\s+(\d+(?:\.\d+)*)$")
SEP_RE = re.compile(r"^\|[\s|:-]+\|$")


class SkipUpdate(Exception):
    """Raised when the table is already up to date for this client version."""


def strip_v(version: str) -> str:
    return version[1:] if version.lower().startswith("v") and version[1:2].isdigit() else version


def parse_cells(line: str) -> List[str]:
    raw = line.strip()
    if not raw.startswith("|"):
        raise ValueError(f"not a table row: {line!r}")
    # Keep trailing empty from final | out
    parts = raw.split("|")
    # split('|a|b|') -> ['', 'a', 'b', '']
    return [p.strip() for p in parts[1:-1]]


def is_separator(line: str) -> bool:
    return bool(SEP_RE.match(line.strip())) and "-" in line


def is_header_row(cells: Sequence[str]) -> bool:
    if not cells:
        return False
    joined = " ".join(cells).lower()
    return "compatibility" in joined or cells[0].startswith("**")


def is_devel_row(cells: Sequence[str]) -> bool:
    if len(cells) < 2:
        return False
    ms = cells[1].lower()
    return "dev" in ms and "development" in ms


def column_widths_from_separator(line: str) -> List[int]:
    cells = []
    raw = line.strip()
    parts = raw.split("|")
    for part in parts[1:-1]:
        # Count dash runs; allow leading/trailing spaces around dashes
        dashes = part.count("-")
        if dashes == 0:
            dashes = max(len(part.strip()), 1)
        cells.append(dashes)
    return cells


def format_row(cells: Sequence[str], widths: Sequence[int]) -> str:
    if len(cells) != len(widths):
        raise ValueError(
            f"cell count {len(cells)} != column count {len(widths)}: {cells}"
        )
    parts = []
    for cell, width in zip(cells, widths):
        parts.append(" " + cell.ljust(width) + " ")
    return "|" + "|".join(parts) + "|"


def with_prefix(version: str, like: str) -> str:
    """Apply the same leading 'v' style as an existing version token."""
    if like.lower().startswith("v") and like[1:2].isdigit():
        bare = strip_v(version)
        return "v" + bare
    return strip_v(version)


def update_table_text(text: str, client_version: str, manticore_version: str) -> str:
    lines = text.splitlines(keepends=True)
    # Work without newlines for parsing; restore later
    plain = [ln.rstrip("\n") for ln in lines]

    # Find separator + following data rows of the first compatibility table
    sep_idx: Optional[int] = None
    for i, line in enumerate(plain):
        if not line.strip().startswith("|"):
            continue
        if is_separator(line):
            # Prefer separators that follow a compatibility header
            # Look back a few lines for header
            window = plain[max(0, i - 3) : i]
            if any(
                is_header_row(parse_cells(w))
                for w in window
                if w.strip().startswith("|") and not is_separator(w)
            ):
                sep_idx = i
                break
            if sep_idx is None:
                sep_idx = i
    if sep_idx is None:
        raise RuntimeError("compatibility table separator row not found")

    widths = column_widths_from_separator(plain[sep_idx])
    ncols = len(widths)

    # Collect contiguous data rows after separator
    data_start = sep_idx + 1
    data_indices: List[int] = []
    for i in range(data_start, len(plain)):
        line = plain[i]
        if not line.strip().startswith("|") or is_separator(line):
            break
        cells = parse_cells(line)
        if len(cells) != ncols:
            break
        if is_header_row(cells):
            break
        data_indices.append(i)

    if len(data_indices) < 5:
        raise RuntimeError(
            f"expected devel row + 4 top-block rows, found {len(data_indices)} data rows"
        )

    devel_i = data_indices[0]
    block_idxs = data_indices[1:5]
    devel_cells = parse_cells(plain[devel_i])
    if not is_devel_row(devel_cells):
        raise RuntimeError(
            f"first data row does not look like a devel row: {devel_cells}"
        )

    rows = [parse_cells(plain[i]) for i in block_idxs]
    r1, r2, r3, r4 = rows

    def require(cond: bool, msg: str) -> None:
        if not cond:
            raise RuntimeError(msg)

    require("Fully" in r1[-1], f"row1 should be Fully Compatible: {r1}")
    require("Partial" in r2[-1], f"row2 should be Partially Compatible: {r2}")
    require("Partial" in r3[-1], f"row3 should be Partially Compatible: {r3}")
    require("Fully" in r4[-1], f"row4 should be Fully Compatible: {r4}")

    m1 = OR_NEWER_RE.match(r1[0])
    m2 = OR_NEWER_RE.match(r2[0])
    require(m1 and m2 and r1[0] == r2[0], f"rows 1-2 client mismatch: {r1[0]!r} / {r2[0]!r}")
    c_cur_token = m1.group(1)

    m3 = RANGE_RE.match(r3[0])
    m4 = RANGE_RE.match(r4[0])
    require(m3 and m4 and r3[0] == r4[0], f"rows 3-4 client mismatch: {r3[0]!r} / {r4[0]!r}")
    c_prev_token, c_cur_from_range = m3.group(1), m3.group(2)
    require(
        strip_v(c_cur_from_range) == strip_v(c_cur_token),
        f"row3 upper client {c_cur_from_range!r} != row1 {c_cur_token!r}",
    )

    ms1 = MS_OR_NEWER_RE.match(r1[1])
    ms3 = MS_OR_NEWER_RE.match(r3[1])
    require(ms1 and ms3 and r1[1] == r3[1], f"rows 1/3 manticore mismatch: {r1[1]!r} / {r3[1]!r}")
    m_cur = ms1.group(1)

    ms2 = MS_RANGE_RE.match(r2[1])
    ms4 = MS_RANGE_RE.match(r4[1])
    require(ms2 and ms4 and r2[1] == r4[1], f"rows 2/4 manticore mismatch: {r2[1]!r} / {r4[1]!r}")
    m_prev, m_cur_from_range = ms2.group(1), ms2.group(2)
    require(m_cur_from_range == m_cur, f"row2 upper manticore {m_cur_from_range} != {m_cur}")

    if strip_v(c_cur_token) == strip_v(client_version):
        raise SkipUpdate(
            f"client version {client_version} already is the top 'or newer' "
            f"entry ({c_cur_token}); skipping"
        )

    c_new = with_prefix(client_version, c_cur_token)
    c_cur = with_prefix(c_cur_token, c_cur_token)
    c_prev = with_prefix(c_prev_token, c_prev_token)
    m_new = strip_v(manticore_version)

    # Preserve middle columns (language etc.) from the corresponding old rows.
    def middle(row: Sequence[str]) -> List[str]:
        return list(row[2:-1])

    fully = r1[-1]
    partial = r2[-1]

    new_r1 = [f"{c_new} or newer", f"{m_new} or newer", *middle(r1), fully]
    new_r2 = [f"{c_new} or newer", f"{m_cur} to {m_new}", *middle(r2), partial]
    new_r3 = [f"{c_cur} to {c_new}", f"{m_new} or newer", *middle(r1), partial]
    new_r4 = [f"{c_cur} to {c_new}", f"{m_cur} to {m_new}", *middle(r2), fully]
    new_r5 = [f"{c_prev} to {c_cur}", f"{m_cur} to {m_new}", *middle(r3), partial]
    # r4 (old) stays as-is at the end of the rewritten block
    kept_r6 = list(r4)

    block_rows = [new_r1, new_r2, new_r3, new_r4, new_r5, kept_r6]
    # Widen columns if new cell text exceeds the original separator width.
    new_widths = []
    for col in range(ncols):
        widest = widths[col]
        for row in block_rows:
            widest = max(widest, len(row[col]))
        new_widths.append(widest)

    formatted = [format_row(row, new_widths) for row in block_rows]
    sep_cells = ["-" * w for w in new_widths]
    new_sep = "|" + "|".join(" " + dashes + " " for dashes in sep_cells) + "|"

    # Also realign header + devel rows to the (possibly widened) columns.
    header_idx = None
    for j in range(sep_idx - 1, max(-1, sep_idx - 4), -1):
        if plain[j].strip().startswith("|") and not is_separator(plain[j]):
            header_idx = j
            break
    header_line = (
        format_row(parse_cells(plain[header_idx]), new_widths)
        if header_idx is not None
        else None
    )
    devel_line = format_row(devel_cells, new_widths)

    out_lines = list(plain)
    if header_line is not None:
        out_lines[header_idx] = header_line
    out_lines[sep_idx] = new_sep
    out_lines[devel_i] = devel_line
    # Replace the 4 old block rows with 6 new ones (indices still valid:
    # we only rewrote earlier lines in-place).
    out_lines = (
        out_lines[: block_idxs[0]] + formatted + out_lines[block_idxs[-1] + 1 :]
    )
    result = "\n".join(out_lines)
    if text.endswith("\n"):
        result += "\n"
    return result


def main(argv: Optional[Sequence[str]] = None) -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--client-version", required=True)
    parser.add_argument("--manticore-version", required=True)
    parser.add_argument("files", nargs="+", type=Path)
    args = parser.parse_args(argv)

    for path in args.files:
        if not path.is_file():
            print(f"ERROR: file not found: {path}", file=sys.stderr)
            return 1
        original = path.read_text(encoding="utf-8")
        try:
            updated = update_table_text(
                original, args.client_version, args.manticore_version
            )
        except SkipUpdate as exc:
            print(f"WARNING: {path}: {exc}", file=sys.stderr)
            continue
        except Exception as exc:  # noqa: BLE001 - surface as CLI error
            print(f"ERROR: {path}: {exc}", file=sys.stderr)
            return 1
        if updated == original:
            print(f"ERROR: {path}: update produced no changes", file=sys.stderr)
            return 1
        path.write_text(updated, encoding="utf-8")
        print(f"Updated compatibility table in {path}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
