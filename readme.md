Love Letter
===

## About

This is a project to recreate the card game "Love Letter"

The first iteration will focus on a command line game

The second on creating a HTTP API to allow multiple forms of games to be created eg. web, desktop client, other.

LoveLetter-As-A-Service - catchy.

## Structure
There's currently 4 main concepts

- cards - representing a single card in the game
- decks - represent the different groups cards can be in, available, burnt, discarded
- players - each player in the round, the cards they have in their hand
- rounds - roughly the game state for a single round

I didn't really design this with a pattern in mind, but it closely represents a behaviour 

## Techincal stuff

tests: `go test`

build: `go build`

## Contributing
- create a feature branch
- commit to the branch
- run tests (hopefully you've been doing this all along)
- submit pull request
- ???
- profit?



