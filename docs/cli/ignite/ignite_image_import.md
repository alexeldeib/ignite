## ignite image import

Import a new base image for VMs

### Synopsis


Import an OCI image as a base image for VMs, takes in a Docker image identifier.
This importing is done automatically when the "run" or "create" commands are run.
The import step is essentially a cache for images to be used later when running VMs.


```
ignite image import <OCI image> [flags]
```

### Options

```
  -h, --help   help for import
```

### Options inherited from parent commands

```
      --log-level loglevel      Specify the loglevel for the program (default info)
      --network-plugin plugin   Network plugin to use. Available options are: [cni docker-bridge] (default docker-bridge)
  -q, --quiet                   The quiet mode allows for machine-parsable output by printing only IDs
```

### SEE ALSO

* [ignite image](ignite_image.md)	 - Manage base images for VMs

