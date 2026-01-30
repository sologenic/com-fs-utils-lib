# Models

This folder contains small proto models, mainly enums, that are used in the project, and which do not have a clearly defined place in the project structure.
Think for example about Network: This could be place of auth, role, and possibly a few others. 
By extracting these models to a separate folder, we can keep the project structure clean and organized.

Since they can be generated for both golang and typescript, the models live in the root folder of the project.

The models all have their own go.mod file for isolation and versioning.


## Building the required files

Once the proto file is updated, the following files need to be generated:

e.g. for role.proto:
* go - role.pb.go, role_grpc.pb.go
* typescript - role.ts

Generating these files can be done by running a build script with following command, which will also generate a build dir and add newly generated files and dir to git commit.

```sh
./bin/build.sh
```
