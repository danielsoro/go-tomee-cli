# Apache TomEE-CLI

This command line tool helps system administrators and developers to manage a instance of TomEE server.

### Start server:

    tomee-cli start --path [path-to-tomee]      // Without TOMEE_HOME variable defined.
    tomee-cli start                             // With TOMEE_HOME variable defined.

### Stop server:

    tomee-cli stop --path [path-to-tomee]       // Without TOMEE_HOME variable defined.
    tomee-cli stop                              // With TOMEE_HOME variable defined.

### Restart server:

    tomee-cli restart --path [path-to-tomee]    // Without TOMEE_HOME variable defined.
    tomee-cli restart                           // With TOMEE_HOME variable defined.


### Deploy application:

    tomee-cli deploy --path [path-to-tomee] [path-to-war/ear-file]
    tomee-cli deploy [path-to-war/ear-file]

### Undeploy application:

    tomee-cli undeploy --path [path-to-tomee] [war/ear-file]
    tomee-cli undeploy [war/ear-file]


## License

Copyright 2015 Daniel Cunha (soro) and/or its affiliates and other contributors as indicated by the @authors tag. All rights reserved.

Distributed under the Apache License V2.
