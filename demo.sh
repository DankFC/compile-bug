#!/bin/bash

unset GOPATH
#run plugin directly to see correct result
go run plugin/plugin.go


#compile as plugin 
go build -buildmode=plugin -o /tmp/plugin.so ../compiler_bug/plugin/

#running plugin will trigger error
go run runplugin/main.go


#However, if the plugin is compiled with generic tag (To disable asm code, the bug doesn't occur), Some register corruption occurs due to wrong code being generated.
#go build -buildmode=plugin -tags generic -o /tmp/plugin.so ../compiler_bug/plugin/

#running plugin will not cause an error
#go run runplugin/main.go

