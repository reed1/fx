# Fork Information
This is a fork of [fx](https://github.com/antonmedv/fx) - a terminal JSON viewer.

## Purpose

This fork adds useful enhancements to improve the fx terminal JSON viewing experience.

## Changes Made

### 1. Transform JSON Feature (Keybinding: "T")

Added the ability to parse a JSON string value and open it in a new fx instance.

**Use Case:** When viewing JSON data that contains stringified JSON (common in API responses, logs, etc.), you can now parse and explore that nested JSON directly.

**How it works:**
- Position cursor on a JSON string value that contains valid JSON
- Press `T` to transform
- The string is parsed as JSON and opens in a new fx instance
- Uses a temporary file that auto-cleans after closing

**Implementation:**
- Added `reed_commands.go` with `transformToJSON()` method
- Modified `keymap.go` to bind the "T" key
- Modified `main.go` to integrate the command

**Example:**
```json
{
  "data": "{\"name\":\"John\",\"age\":30}"
}
```

Pressing `T` while on the `"data"` value will parse and open:
```json
{
  "name": "John",
  "age": 30
}
```

### 2. Print Keybindings as JSON (Flag: `--print-keybindings`)

Added a CLI flag to output all keybindings in JSON format.

**Use Case:** Useful for scripting, documentation, or quickly checking what keys do what without launching the TUI.

**How it works:**
```bash
fx --print-keybindings
```

**Output format:**
```json
{
  "Quit": {
    "keys": ["q", "ctrl+c", "esc"],
    "description": "exit program"
  },
  "Down": {
    "keys": ["down", "j"],
    "description": "down"
  },
  ...
}
```

**Implementation:**
- Added `--keybindings` flag in `main.go`
- Added `printKeybindingsJSON()` function in `help.go`

### 3. Search Input Auto-Clear

Modified search behavior to automatically clear the input when entering search mode, instead of retaining the previous search value.

**Implementation:**
- Modified `main.go` to clear search input value when entering search mode

## Upstream

This fork is based on fx and may be merged with upstream changes periodically.

**Original repository:** https://github.com/antonmedv/fx
**Original author:** Anton Medvedev
**License:** MIT