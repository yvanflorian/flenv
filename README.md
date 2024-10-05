# Flenv: A Simple Environment Variables manager
Tool mostly created to load environment variables in my shell session depending on which which "Stage" I want to work with.
Working mostly in the terminal, need to have a way to switch to different stages (`prod`, `uat`, `dev`, `local`) and have
my environment variables values switch quickly.

All variables should be kept in a JSON file (preferably `$HOME/.flenv.json`) and have it encrypted: with `ansible-vault` or `gpg`
The file maybe should be safe to be committed in a git repo.

Example Structure of the file like [this one](./.config.json)
## Commands Proposal:
Suggested workflow:

### Stages
1. Create the stages with `flenv stage --create <stagename>`
2. List stages with `flenv stage --list`
3. Set the current Stage `flenv stage --set <stagename>`

### Configs
1. Create config with `flenv config --create <configname>` to multiply for each config
2. List configs available for a given config: `flenv config --list `

### Variables
1. Create Variables under stages > configs:
   `flenv variable --create <variablename> --config <configname>`
   Then Prompt to enter the value for as many stages as you have
2. Show variable value: `flenv variable --show <variablename> --config <configname>` (with optional stage)
3. Edit variable value: `flenv variable --edit <variablename> --config <configname> --stage <stagename>`

### Encryption
Since the file will most likely contain sensitive values, thought it best to have it encrypted.
Key is stored in the OS KeyRing and I'm using `gpg` to encrypt and decrypt the config file for each read and write.
Note: Planning Later to allow `1Password` storing of the key
