# canonical
# Exercise 1. qemu folder
Solution of quemu exercise.
It downloads busybox, a linux image, compiles the kerne, makes an initrd filesystem and launches the qemu.
Qemu crashes when shutting down, I did not have time to fix it. After shutting down, the script would remove all the linux and temporary files
It assumes that wget and qemu-system-x86_64 is installed 

# Exercise 2. go folder: 
go solution exercise. It assumes that go framework is installed in the machine. I have developed the application and some test skeleton.
The test cases would be:
- Error on file open fail
- Error on file write fail
- Error on file write an incorrect number of times
- Error on file deletion
- Error on non-random data
- Success if everything passes

I have only implemented the first test case. os.File type should be mocked in order to fulfill the rest of the cases

The application expects one argument, which is the file to be shredded.

The use cases for this function would be to:
- check that a filesystem behaves correctly and writes the data to a persistent media
- Check that the data is really random. I.E: randomizer is "secure".
- 

