# pokedex

A command-line Pokedex REPL in Go, backed by [PokéAPI](https://pokeapi.co/).

## Requirements

- Go 1.25+

## Run

```sh
go run .
```

Or build a binary:

```sh
go build -o pokedex . && ./pokedex
```

## Commands

| Command | Description |
| --- | --- |
| `help` | Show available commands |
| `map` | List the next 20 location areas |
| `mapb` | List the previous 20 location areas |
| `explore <location>` | List Pokémon found in a location area |
| `catch <pokemon>` | Try to catch a Pokémon (higher base XP = harder) |
| `inspect <pokemon>` | Show stats and types of a caught Pokémon |
| `pokedex` | List caught Pokémon |
| `exit` | Quit |

Input is case-insensitive.

## Example

```text
Pokedex > map
canalave-city-area
eterna-city-area
...
Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokemon:
 - tentacool
 - magikarp
 ...
Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
magikarp was caught!
Pokedex > inspect magikarp
```

## License

MIT
