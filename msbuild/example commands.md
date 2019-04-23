### Using rsp
```
msbuild .\DoSomething.proj @args.rsp
```
### Setting verbosity
```
msbuild .\DoSomething.proj /v:minimal

msbuild .\DoSomething.proj /v:normal

msbuild .\DoSomething.proj /v:detailed

msbuild .\DoSomething.proj /v:diagnostic
```
### Using parameters
```
msbuild .\DoSomething.proj @args.rsp /p:Name=Lisa
```
### Target Order
```
msbuild .\Lifecycle.proj /t:TargetC;TargetA
```
### SLN and *proj
SLN files are not msbuild files, but they get converted by msbuild
```
msbuild helloworld.sln
msbuild helloworld.csproj
```
