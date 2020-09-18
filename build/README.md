# Go build exemples

## Removing debug information included in the executable
```go build -ldflags "-w" main.go```

## To reduce even more using -s flag (strip)
```go build -ldflags "-s -w" main.go```

# Compressing binary using UPX

## Install 
```apt-get install upx```
## Using
```upx mai```

Output
```
                       Ultimate Packer for eXecutables
                          Copyright (C) 1996 - 2017
UPX 3.94        Markus Oberhumer, Laszlo Molnar & John Reiser   May 12th 2017

        File size         Ratio      Format      Name
   --------------------   ------   -----------   -----------
   2990944 ->   1157704   38.71%   linux/amd64   main                          

Packed 1 file.
```