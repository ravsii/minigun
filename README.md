# Minigun

Just a fun project. `Minigun` because it _supposted_ to be fast, not sure if it will be lol.

It's not usable at the moment, but I hope it will be one day...

## WIP

Implementation order is "when needed"

- [ ] Debug info
  - [ ] `.log` file
  - [x] Some on-screen output (Command mode)
    - [x] Error messages
    - [x] Info messages
- [ ] Modes
  - [x] View
    - [x] `hjkl` movements
  - [ ] Insert Mode
    - [ ] `i`, `I`, `a`, `A` will do for now
  - [x] Command Mode (`:` from vim)
    - [x] Basic input
    - [x] Command handling
      - [x] ~~You can `:q`, do you need anything else?~~
- [x] Status bar
  - [x] Current Mode
  - [x] Cursor Line / Position
- [ ] Commands
  - [x] Make all integrations with the minigun via commands, no hardcoded actions
  - [x] Opening files (`:o`, `:open`)
  - [ ] Saving file
- [ ] UI Library - We need some kind of a UI components library, because rn everything is hardcoded so its not good.
  - [ ] Must Have
    - [ ] Padding
    - [ ] Margin
- [ ] File / Project tree
- [ ] Config options
  - [ ] Global / per project config
  - [ ] Tab size
  - [ ] Line Numbers
  - [x] Keybinds
    - Technically all input is getting handled by the keybinds already, I just need to move them into config
- [ ] Multiple windows / tabs
- [ ] Tree sitter
- [ ] LSP (multiple for the same lang)
- [ ] Formatters
- [ ] Plugins (ideally no lua)

## Preview

![](./demo/base.gif)
