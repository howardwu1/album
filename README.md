# album

##Install Go

Go to this link to get the instructions to install the latest version of Go
https://golang.org/doc/install


##Run/Test Application

Run the console application
```bash
go build album.go
```

Test the console application
```bash
go test
```

##Using the Application

Wait for the prompt to ask for the photo-album number and then enter a positive integer. Here's an example with 10:
```bash
> photo-album 10

[451] dolorem accusantium corrupti incidunt quas ex est

[452] mollitia dolorem qui

...
```

If you enter an album that doesn't exist, you will get a message that tells you:
```bash
> photo-album 10000
Album not found or is empty
```

Negative integers and non-integer strings will trigger a message as well:
```base
> photo-album hello
Please re-enter a valid positive integer for the photo-album
```

