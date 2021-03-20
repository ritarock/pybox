# pybox

## install
```
$ git clone https://github.com/ritarock/pybox.git
$ cd pybox/
$ go install
```

## Usage
```
NAME:
   pybox - Run pthon

USAGE:
   [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --file value, -f value     target file
   --version value, -v value  target version (default: 3.7)
   --help, -h                 show help (default: false)
```

## Example
Put the file under `sandbox`.
```
$ echo "print(\"sample\")" > sandbox/sample.py
$ pybox -f sample.py -v 3.7

Attaching to py37_runner_1
runner_1  | sample
py37_runner_1 exited with code 0
```
