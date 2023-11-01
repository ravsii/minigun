# Minigun

Just a fun project. `Minigun` because it _supposted_ to be fast, not sure if it will be lol.

It's not usable at the moment, but I hope it will be one day...

It aims to be a battaries-included terminal text editor. Ideally I'd like to switch off VSCode, but `nvim` is just not for me and `helix` is being released like once a year nowadays.

## Glossary

- `Tab` - usually what is called a buffer in other editors, but we'd like to keep it simple.

## Status

Implementation order is "when needed". Literally everything is WIP and could be changed any minute.

### Done

- [x] Debug info
  - [x] `.log` file
    - Defaults to `$HOME/.config/minigun/debug.log`, can be overwritten by using `-logfile`
- [x] Status bar - like modes and unlike nvim\helix, status bar is global and not per-tab for easier inspection of the current "state"
  - [x] Current Mode
  - [x] Cursor Line / Position
- [x] Commands
- [x] Basic Components Library - We need some kind of a UI components library, because rn everything is hardcoded so its not good.
  - [x] Box
  - [x] Text Box

### Bugs

- [ ] Cursor appears on empty lines

### Work in Progress

- [ ] Modes - Unlike in helix/nvim where mode is usually per-tab, minigun modes are global and represent the current "state" of the editor
  - [x] View
  - [ ] Insert Mode
    - [ ] `i`, `I`, `a`, `A` will do for now
  - [x] Command Mode (`:` from vim)
    - [ ] `wq`
    - [ ] Although we don't have any _errors_, exclamations mark support would be nice for forced stuff like `:wq!`
  - [x] Replace Mode (single char, `r` from vim)
  - [ ] File / Workspace mode (kind of file explorer)
  - [ ] Still thinking about select-like mode. tbh I like `helix`'s approach more. (select-action rather than `nvim`'s action-select)
- [ ] Config options
  - [ ] Global / per project config
  - [ ] Line Numbers
  - [x] Keybinds
    - Technically all input is getting handled by the keybinds already, I just need to move them into config

### "Actually usable"-level features

- [ ] Themes
  - [ ] Nerd font icons or something
- [ ] Multiple windows / tabs
  - Probably the easiest of all these
- [ ] Tree sitter
- [ ] LSP (multiple for the same lang)
  - `helix` kind of has the best LSP integration out of the box, but you still need to configure it sometimes, which is very painful if you haven't done it before.
- [ ] Formatters
- [ ] Plugins (ideally no lua)

## Preview (current WIP version)

![demo gif](./demo/base.gif)

## List of Commands

[here](https://github.com/ravsii/minigun/blob/main/internal/command/commands.go)
