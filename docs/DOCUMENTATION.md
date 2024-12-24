# Obsidian CLI Documentation

## Overview
This project is a command-line interface (CLI) for managing notes in an Obsidian vault.

## Directory Structure
```
/home/davdev/Projects/go/obsidian-cli/
├── cmd/
│   └── create.go        # Command for creating notes in the vault
├── Makefile             # Build and install commands
├── main.go              # Entry point of the application
├── main_test.go         # Tests for the application
└── docs/                # Documentation files
```

## Command Documentation
### cmd/create.go
- **Purpose**: This file contains the command for creating notes in the Obsidian vault.

- **Key Functions**:
  - `createNoteCmd`: The command used to create a note. It accepts the following flags:
    - `--vault`, `-v`: Specify the vault name.
    - `--open`: Option to open the created note immediately.
    - `--content`, `-c`: Text content to add to the note.
    - `--append`, `-a`: Option to append to an existing note.
    - `--overwrite`, `-o`: Option to overwrite an existing note.

- **Functionality**: The command takes a note name as an argument and creates a note in the specified vault. It uses the `actions.CreateNote` function to perform the creation, handling any errors that may occur during the process.

### cmd/delete.go
- **Purpose**: This file contains the command for deleting notes in the Obsidian vault.

- **Key Functions**:
  - `deleteCmd`: The command used to delete a note. It accepts the following arguments:
    - `<note_path>`: The path of the note to delete.

### cmd/move.go
- **Purpose**: This file contains the command for moving or renaming notes in the Obsidian vault.

- **Key Functions**:
  - `moveCmd`: The command used to move or rename a note. It accepts the following arguments:
    - `<current_name>`: The current name of the note.
    - `<new_name>`: The new name for the note.

### cmd/open.go
- **Purpose**: This file contains the command for opening notes in the Obsidian vault.

- **Key Functions**:
  - `OpenVaultCmd`: The command used to open a note. It accepts the following argument:
    - `<note_name>`: The name of the note to open.

### cmd/print_default.go
- **Purpose**: This file contains the command for printing the default vault name and path.

- **Key Functions**:
  - `printDefaultCmd`: The command that prints the default vault name and path.

### cmd/search.go
- **Purpose**: This file contains the command for searching notes in the Obsidian vault.

- **Key Functions**:
  - `searchCmd`: The command used to search for a note. It accepts the following argument:
    - `<search_text>`: The text to search for in the notes.

### cmd/set_default.go
- **Purpose**: This file contains the command for setting the default vault.

- **Key Functions**:
  - `setDefaultCmd`: The command used to set the default vault. It accepts the following argument:
    - `<vault_name>`: The name of the vault to set as default.

## Usage
To create a note, use the following command:
```bash
obsidian-cli create --vault <vault_name> --content "<note_content>" <note_name>
```

## Build and Installation
Use the Makefile to build and install the application.

## Testing
Run tests using:
```bash
go test ./...
```

## Additional Notes
- Ensure that the required dependencies are installed before running the application.
