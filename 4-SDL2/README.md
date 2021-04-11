To use SDL2

1. Download SDL2 software toolkit from libsdl.org

2. Download mingw-w64 gcc compiler

3. install the compiler

4. open the tar.gz sdl2 file and migrate to the folder where you have the bin, lib, include, and share folder @ SDL2-devel-2.0.14-mingw.tar.gz\SDL2-2.0.14\x86_64-w64-mingw32 - TAR+GZIP archive, unpacked size 82,110,616 bytes.

5. copy the four folders and paste in the bin directory for the gcc compiler @ path-to-mingw-w64\x86_64-8.1.0-posix-seh-rt_v6-rev0\mingw64\x86_64-w64-mingw32

6. set two new system environmental variables 
    1. path-to-mingw-w64\x86_64-8.1.0-posix-seh-rt_v6-rev0\mingw64\x86_64-w64-mingw32\bin
    2. path-to-mingw-w64\\x86_64-8.1.0-posix-seh-rt_v6-rev0\mingw64\bin

    7. run "go get github.com/veandco/go-sdl2/sdl" to get the sdl2 bindings for Go

    8. step 7 will not work unless the once before it are done correctly