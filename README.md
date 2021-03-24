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

Use library
```
$ mkdir sandbox/sample
$ echo "simplejson" > sandbox/sample/requirements.txt
$ echo "import simplejson as json
print(json.dumps(['foo', {'bar': ('baz', None, 1.0,2)}]))" > sandbox/sample/sample.py

$ pybox -f sample.py -v 3.7 -d sample
runner_1  | Collecting simplejson
runner_1  |   Downloading simplejson-3.17.2-cp37-cp37m-manylinux2010_x86_64.whl (128 kB)
runner_1  | Installing collected packages: simplejson
runner_1  | Successfully installed simplejson-3.17.2
runner_1  | ["foo", {"bar": ["baz", null, 1.0, 2]}]
```
