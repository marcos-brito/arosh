# Overview

The most basic pipeline for a shell is: `read input` -> `parse` -> `execute`. Arosh does that, but with some extra steps
along the way. The entry point for interacting with this shell is throught the line editor. This is a concept borrowed from [zsh](https://www.zsh.org/). It
provides easy ways to extend functionallity and change behaviour. Highlights, completion, keybindings, command previews and other, are implemented using the
line editor API and are called `widgets`.

After reading input, the data pass throught lexical analysis, expansion, parsing and then execution. The syntax follows the definition of the [Shell Command Language](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html#tag_18), but
not every step is made as described in the document.

# Modules

- **lexer:** Lexical analysis. Here, the raw text inserted by the user is tokenized in something easier to work with. It's pretty straightforward and things
  just go wrong if some unexpected character is found.
- **parsing:**
- **expansion:**
- **env:**
- **builtins:** Implementaion for all the builtin commands (e.g `cd`, `pwd`, `set`).
- **lineEditor:** It's a TUI implemented with ncurses, it also defines an API for widgets and some internal use. The API have methods for moving around, changing text, add or overwrite keybings and so forth.
- **widgets:** This is the home for some builtin widgets. They are all implemented using the line editor API and can be easily replaced for external implemenations.
  - **highlights:**
  - **history:**
  - **git:**
  - **completion:**
  - **preview:**
