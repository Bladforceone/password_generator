# Password Generator

Password Generator is a simple Go utility for generating random passwords with customizable length and character types (letters, numbers, special characters).

## Features

- Generate passwords of a specified length.
- Support for letters, numbers, and special characters.
- Option to choose which types of characters to include in the password.

## Command Line Arguments
| Flag     | Description                   | Type  | Example      |
| -------- | ----------------------------- | ----- | ------------ |
| `-l`     | Length of the password         | int   | `-l 12`      |
| `--let`  | Use letters                    | bool  | `--let=false`|
| `--numb` | Use numbers                    | bool  | `--numb=true`|
| `--spec` | Use special characters         | bool  | `--spec=true`|
| `--count`| Number of passwords to generate| int   | `--count 5`  |

## Usage Examples

By default, the generator creates an 8-character password containing only letters:
```bash
password_generator
```

Generate a 10-character password using only special characters:
```bash
password_generator -l 10 --let=false --spec=true
```

Generate a 16-character password using letters, numbers, and special characters:
```bash
password_generator -l 16 --let=true --numb=true --spec=true
```