# logger
A really simple logger with four levels: Debug, Info, Warning, Error. Created this to help me learn golang.
It is a lightweight wrapper for golang's logger and intended to be a starting point for more complex loggers.

Basic usage:
```golang
	log := Logger{}
	log.Start(INFO, nil, "myproject.log")
	log.Warning("Danger Will Smith")
	log.Stop()
```