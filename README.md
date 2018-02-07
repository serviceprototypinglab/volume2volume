# volume2volume

`volume2volume` is a command line interface tool which helps `openshift` users to migrate their data
 between different `openshift` clusters or providers.
 
 `volume2volume` will be connected to two clusters (old cluster and new cluster) and with simple commands 
 will identify all the volumes in both clusters, pairs these volumes and make the backup the volumes of 
 the old cluster and recovery of these data in the the new cluster.
  `volume2volume` uses the command line client of `openshift`, `oc`, to communicate between the `openshift` clusters 
  and `stash` for the backup and recovery of the data. 

 A unique advantage of `volume2volume` is that it combines multiple steps into a convenient workflow for the migration
  of data volumes, including backup and recovery of the data.

 `volume2volume` is a prototype from the Cloud-Native Applications research initiative of the Service Prototyping Lab
  at Zurich University of Applied Sciences. Use with care, things may break. We will share our findings
  on cloud application migration at a later point in time.

## Use Case

Use the command `help` to see the list of command and the description of them.

With `findVolumes`, you will identify the volumes in the two clusters and make the pairs.
With `backup`, you will create the restic objects to do the backup of the volumes. (in the old cluster).
With `up`, you will create the recovery objects to upload the data to the volumes. (in the new cluster).
With `migrate`, you will combine the 3 before commands.

The configuration can be added in `~/.volume2volume.yaml` or directly using the flags.
Use help to see all the flags.
If no configuration is given to the tool, it will take the default values.

```
    volume2volume help
    volume2volume findVolumes
    volume2volume backup
    volume2volume up
    volume2volume ...
```

## Installation

### Install `oc`

- https://docs.openshift.org/latest/cli_reference/get_started_cli.html#installing-the-cli

### From binary

#### Install the binary: `volume2volume`

Download the binary from /binaries/< your operative system> and run:

```
    chmod +x volume2volume
    sudo mv ./volume2volume /usr/local/bin/volume2volume
```

### From source

Note: We are in the progress of converting to a default Go project structure.
Until then, please use the following commands to compile and install:
 
```
    git clone <this repository>
    mv volume2volume $GOPATH/src/volume2volume
    go get github.com/mitchellh/go-homedir
    go get github.com/spf13/cobra
    go get github.com/spf13/viper
    cd $GOPATH/src/volume2volume
    go build volume2volume
    chmod +x volume2volume
    sudo mv ./volume2volume /usr/local/bin/volume2volume
```

## First steps

This small example shows how to migrate the data in the volumes of an OpenShift application from 
a local OpenShift development cluster to APPUiO, the Swiss Container Platform.

```
    volume2volume migrate \
          --clusterFrom https://127.0.0.1:8443 --clusterTo https://console.appuio.ch:443 \
          --projectFrom test --projectTo test \
          --usernameFrom user --usernameTo user \
          --passwordFrom pass --passwordTo pass
```

Considering the large number of options, it is advised to use the configuration file `~/.volume2volume.yaml`
to store all parameters (in YAML syntax).