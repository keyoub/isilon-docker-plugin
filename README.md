EMC Isilon Docker Plugin
======================
The volume-driver plugin for Docker containers.

## Description
The Isilon volume-driver plugin works alongside the Docker engine to provide volume support for containers. The volumes
are hosted on an Isilon cluster therefore making all Isilon data management tools(snapshots, backups, ...) instantly available to the container volume.


## Installation
- Install and setup Go environment: http://golang.org/doc/install
- Install EMC Isilon volume-driver plugin for Docker containers
```bash
$ go get github.com/bkeyoumarsi/isilon-docker-plugin
$ cd $GOPATH/src/github.com/bkeyoumarsi/isilon-docker-plugin
$ make install
```
## Usage Instructions
To use the driver a bit of preparation on the Isilon cluster is needed.

On the Isilon cluster, run the following commands with root privileges:
```bash
$ mkdir -p /ifs/data/docker/volumes
$ chown nobody:nobody /ifs/data/docker/volumes
```

To start the plugin run the commands below on Docker host where plugin is installed:
```bash
$ sudo $GOPATH/bin/isilon-docker-plugin -cluster-ip <x.x.x.x> -username=<root-user> -password=<password>
```

To use the plugin with your containers pass ```--volume-driver=isilon``` option to the docker run command.
```bash
$ docker run -it --volume-driver=isilon -v test_volume:/data ubuntu /bin/bash
```
## Future
- Add better testing
- Add more protocol support (smb, swift, ...)

## Contribution
Create a fork of the project into your own reposity. Make all your necessary changes and create a pull request with a description on what was added or removed and details explaining the changes in lines of code. If approved, project owners will merge it.

Licensing
---------
**EMC CODE does not provide legal guidance on which open source license should be used in projects. We do expect that all projects and contributions will have a valid open source license, or align to the appropriate license for the project/contribution**

“The MIT License (MIT)
Copyright (c) [Year], [Company Name (e.g., EMC Corporation)]
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.”


Support
-------
Please file bugs and issues at the Github issues page. For more general discussions you can contact the EMC Code team at <a href="https://groups.google.com/forum/#!forum/emccode-users">Google Groups</a> or tagged with **EMC** on <a href="https://stackoverflow.com">Stackoverflow.com</a>. The code and documentation are released with no warranties or SLAs and are intended to be supported through a community driven process.
