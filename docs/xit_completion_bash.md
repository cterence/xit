## xit completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(xit completion bash)

To load completions for every new session, execute once:

#### Linux:

	xit completion bash > /etc/bash_completion.d/xit

#### macOS:

	xit completion bash > $(brew --prefix)/etc/bash_completion.d/xit

You will need to start a new shell for this setup to take effect.


```
xit completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.xit.yaml)
      --dry-run         Dry run the command
```

### SEE ALSO

* [xit completion](xit_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 10-Jun-2023