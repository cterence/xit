## xit completion powershell

Generate the autocompletion script for powershell

### Synopsis

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	xit completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
xit completion powershell [flags]
```

### Options

```
  -h, --help              help for powershell
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