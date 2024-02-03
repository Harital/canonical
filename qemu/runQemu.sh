#sudo apt-get install grub-pc-bin xorriso
#!/bin/bash

# Create a directory for the filesystem
mkdir -p myfs
cd myfs
mkdir -p bin dev etc lib mnt proc sbin sys tmp var usr/bin usr/sbin

# Download and extract BusyBox
wget https://www.busybox.net/downloads/binaries/1.35.0-x86_64-linux-musl/busybox -O bin/busybox
chmod +x bin/busybox
cd bin
ln -s busybox sh
ln -s busybox mount
cd ..

# Create init script
cat >>init << EOF
#!/bin/sh

mount -t devtmpfs  devtmpfs  /dev
mount -t proc      proc      /proc
mount -t sysfs     sysfs     /sys
mount -t tmpfs     tmpfs     /tmp

echo "Hello World!"
sh
EOF

chmod +x init

# Create filesystem image
find . | cpio -ov -H newc | gzip > ../rootfs.cpio.gz
cd ..

build kernel
wget https://cdn.kernel.org/pub/linux/kernel/v6.x/linux-6.7.2.tar.xz
tar -xvf linux-6.7.2.tar.xz
cd linux-6.7.2
make ARCH=x86_64 defconfig
make -j `nproc`
cd ..

# Run QEMU with the created filesystem
qemu-system-x86_64 -kernel linux-6.7.2/arch/x86/boot/bzImage -initrd rootfs.cpio.gz  -nographic -append "console=ttyS0"


