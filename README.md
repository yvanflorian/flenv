# Flenv: A Simple Environment Variables manager
Tool mostly created to load environment variables in my shell session depending on which which "Stage" I want to work with.
Working mostly in the terminal, I need to have a way to switch to different stages (`prod`, `uat`, `local`, etc.) and have
my environment variables values switch quickly.

For example, I've got some curl commands I frequently run which I usually only modify endpoints (and maybe also credentials) to point to another "stage".
`flenv` can now serve as the tool to store these variables and help switch easily.

As I'm using [starship] as the prompt for my shell, you may add below custom module to display which 'stage' you're currently working with.
```toml
[custom.flenv_cli_load]
command = "echo $FLENV_STAGE_ENVIRONMENT"
when = "test -n \"$FLENV_STAGE_ENVIRONMENT\""
style = "bold yellow"
format = "[<flenv:$output>]($style) "

```

Example Structure of the file like [this one](./.config.json)
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
