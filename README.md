# canonical
# Go folder: 
go solution exercise
The test cases would be:
- Error on file open fail
- Error on file write fail
- Error on file write an incorrect number of times
- Error on file deletion
- Success if everything passes

I have only implemented the first test case. os.File type should be mocked in order to fulfill the rest of the cases

# qemu folder
Solution of quemu exercise.
Qemu crashes when shutting down, I did not have time to fix it. After shutting down, the script would remove all the linux and temporary files
It assumes that wget and qemu-system-x86_64 is installed 
