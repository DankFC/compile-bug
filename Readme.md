**Compiler Bug

A bug which occurs only in plugins with assembly code.  
The package consists of a bn256 crypto implementation which has amd64 assembly code and also a generic implementation.  
The assembly implementation is default and generic implementation can be enabled using the "generic" tag.  
The plugin when compiled as a binary runs **OK** with both assembly and go implementation.  
The plugin when compiled as a plugin  **FAILS** with assembly code  but runs OK with generic implementation.  

As far as my understanding goes, everything should be running okay.  

**Please see attached source code and test on x86_64 linux system.**  
`go version : 1.15.7` ( However go version 1.14 and 1.16beta1 also show same error behaviour.)  
`OS: linux x86_64`  

My initial guess is some sort of corruption ( Or wrong code generation) by the go compiler/assembler.  

**TESTING:**  

```
#Run plugin directly to see correct result.
go run plugin/plugin.go

#Compile as plugin.
go build -buildmode=plugin -o /tmp/plugin.so ../compiler_bug/plugin/

#Running plugin will trigger error
go run runplugin/main.go
```


However, if the plugin is compiled with generic tag (To disable asm code, the bug doesn't occur), Some register corruption occurs due to wrong code being generated.  
`go build -buildmode=plugin -tags generic -o /tmp/plugin.so ../compiler_bug/plugin/`

**Running plugin will not cause an error.**  
`go run runplugin/main.go`


Maybe to detect such case, test cases can also be run in plugin mode to detect such cases.  
This can be done by adding another argument to go test for plugin mode testing.  


**Actual behaviour:**  

```
./demo.sh 
Actual a01f9bcc1208dee302769931ad378a4c0c4b2c21b0cfb3e752607e12d2b6fa6425 
Expected a01f9bcc1208dee302769931ad378a4c0c4b2c21b0cfb3e752607e12d2b6fa6425 
running plugin 
panic: Decode point err bn256: malformed point err 
goroutine 1 [running]: 
_/tmp/compiler_bug/plugin.Export() 
/tmp/compiler_bug/plugin/plugin.go:23 +0x2d7 
main.main() 
/tmp/compiler_bug/runplugin/main.go:24 +0x107 
exit status 2 
```

**Expected behaviour:**  
```
./demo.sh 
Actual a01f9bcc1208dee302769931ad378a4c0c4b2c21b0cfb3e752607e12d2b6fa6425 
Expected a01f9bcc1208dee302769931ad378a4c0c4b2c21b0cfb3e752607e12d2b6fa6425 
running plugin 
Actual a01f9bcc1208dee302769931ad378a4c0c4b2c21b0cfb3e752607e12d2b6fa6425 
Expected a01f9bcc1208dee302769931ad378a4c0c4b2c21b0cfb3e752607e12d2b6fa6425 
```
**demo.sh is included in the repo to recreate the bug easily.**
