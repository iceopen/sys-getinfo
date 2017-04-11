# 接口文档说明

## swap 区使用情况
### 地址： 
> ip:port/swap/info

### 请求方法：
GET

### 返回参数：
```javascript
{
  "total": 1073741824,
  "used": 4456448,
  "free": 1069285376,
  "usedPercent": 0.4150390625,
  "sin": 0,
  "sout": 0
}
```
##  内存使用情况
### 地址： 
> ip:port/vmem/info

### 请求方法：
GET

### 返回参数：
```javascript
{
  "total": 17179869184,
  "available": 6403690496,
  "used": 10776178688,
  "usedPercent": 62.72561550140381,
  "free": 135561216,
  "active": 7000592384,
  "inactive": 6268129280,
  "wired": 2920157184,
  "buffers": 0,
  "cached": 0,
  "writeback": 0,
  "dirty": 0,
  "writebacktmp": 0,
  "shared": 0,
  "slab": 0,
  "pagetables": 0,
  "swapcached": 0
}
```
##  CPU使用情况
### 地址： 
> ip:port/cpu/time

### 请求方法：
GET

### 返回参数：
```javascript
[
  {
    "cpu": "cpu0",
    "user": 1557.59375,
    "system": 1236,
    "idle": 10911.109375,
    "nice": 0,
    "iowait": 0,
    "irq": 0,
    "softirq": 0,
    "steal": 0,
    "guest": 0,
    "guestNice": 0,
    "stolen": 0
  },
  {
    "cpu": "cpu1",
    "user": 123.109375,
    "system": 75.015625,
    "idle": 13505.8203125,
    "nice": 0,
    "iowait": 0,
    "irq": 0,
    "softirq": 0,
    "steal": 0,
    "guest": 0,
    "guestNice": 0,
    "stolen": 0
  },
  {
    "cpu": "cpu2",
    "user": 1406.7734375,
    "system": 688.3984375,
    "idle": 11608.9296875,
    "nice": 0,
    "iowait": 0,
    "irq": 0,
    "softirq": 0,
    "steal": 0,
    "guest": 0,
    "guestNice": 0,
    "stolen": 0
  }
]
```

##  CPU信息
### 地址： 
> ip:port/cpu/info

### 请求方法：
GET

### 返回参数：
```javascript
[
  {
    "cpu": 0,
    "vendorId": "GenuineIntel",
    "family": "6",
    "model": "70",
    "stepping": 1,
    "physicalId": "",
    "coreId": "",
    "cores": 4,
    "modelName": "Intel(R) Core(TM) i7-4850HQ CPU @ 2.30GHz",
    "mhz": 2300,
    "cacheSize": 256,
    "flags": [
      "syscall",
      "xd",
      "1gbpage",
      "em64t",
      "lahf",
      "lzcnt",
      "rdtscp",
      "tsci",
      "smep",
      "erms",
      "rdwrfsgs",
      "tsc_thread_offset",
      "bmi1",
      "hle",
      "avx2",
      "bmi2",
      "invpcid",
      "rtm",
      "fpu_csds",
      "fpu",
      "vme",
      "de",
      "pse",
      "tsc",
      "msr",
      "pae",
      "mce",
      "cx8",
      "apic",
      "sep",
      "mtrr",
      "pge",
      "mca",
      "cmov",
      "pat",
      "pse36",
      "clfsh",
      "ds",
      "acpi",
      "mmx",
      "fxsr",
      "sse",
      "sse2",
      "ss",
      "htt",
      "tm",
      "pbe",
      "sse3",
      "pclmulqdq",
      "dtes64",
      "mon",
      "dscpl",
      "vmx",
      "smx",
      "est",
      "tm2",
      "ssse3",
      "fma",
      "cx16",
      "tpr",
      "pdcm",
      "sse4.1",
      "sse4.2",
      "x2apic",
      "movbe",
      "popcnt",
      "aes",
      "pcid",
      "xsave",
      "osxsave",
      "seglim64",
      "tsctmr",
      "avx1.0",
      "rdrand",
      "f16c"
    ],
    "microcode": ""
  }
]
```

##  根目录磁盘使用情况
### 地址： 
> ip:port/disk/usage

### 请求方法：
GET

### 返回参数：
```javascript
{
  "path": "/",
  "fstype": "hfs",
  "total": 499418034176,
  "free": 361791827968,
  "used": 137364062208,
  "usedPercent": 27.504826179262782,
  "inodesTotal": 4294967279,
  "inodesUsed": 2316801,
  "inodesFree": 4292650478,
  "inodesUsedPercent": 0.053942227018302744
}
```

##  磁盘列表
mountpoint 参数为 根据传参目录磁盘使用情况 使用目录
### 地址： 
> ip:port/disk/list

### 请求方法：
GET

### 返回参数：
linux 返回案例：
```javascript
[
  {
    "device": "sysfs",
    "mountpoint": "/sys",
    "fstype": "sysfs",
    "opts": "rw,nosuid,nodev,noexec,relatime"
  },
  {
    "device": "proc",
    "mountpoint": "/proc",
    "fstype": "proc",
    "opts": "rw,nosuid,nodev,noexec,relatime"
  },
  {
    "device": "udev",
    "mountpoint": "/dev",
    "fstype": "devtmpfs",
    "opts": "rw,nosuid,relatime,size=4066816k,nr_inodes=1016704,mode=755"
  },
  {
    "device": "devpts",
    "mountpoint": "/dev/pts",
    "fstype": "devpts",
    "opts": "rw,nosuid,noexec,relatime,gid=5,mode=620,ptmxmode=000"
  },
  {
    "device": "tmpfs",
    "mountpoint": "/run",
    "fstype": "tmpfs",
    "opts": "rw,nosuid,noexec,relatime,size=817136k,mode=755"
  },
  {
    "device": "/dev/mapper/ubuntu--134--vg-root",
    "mountpoint": "/",
    "fstype": "ext4",
    "opts": "rw,relatime,errors=remount-ro,data=ordered"
  },
  {
    "device": "securityfs",
    "mountpoint": "/sys/kernel/security",
    "fstype": "securityfs",
    "opts": "rw,nosuid,nodev,noexec,relatime"
  },
  {
    "device": "tmpfs",
    "mountpoint": "/dev/shm",
    "fstype": "tmpfs",
    "opts": "rw,nosuid,nodev"
  },
  {
    "device": "tmpfs",
    "mountpoint": "/run/lock",
    "fstype": "tmpfs",
    "opts": "rw,nosuid,nodev,noexec,relatime,size=5120k"
  },
  {
    "device": "tmpfs",
    "mountpoint": "/sys/fs/cgroup",
    "fstype": "tmpfs",
    "opts": "rw,mode=755"
  },
  {
    "device": "cgroup",
    "mountpoint": "/sys/fs/cgroup/systemd",
    "fstype": "cgroup",
    "opts": "rw,nosuid,nodev,noexec,relatime,xattr,release_agent=/lib/systemd/systemd-cgroups-agent,name=systemd"
  },
  {
    "device": "pstore",
    "mountpoint": "/sys/fs/pstore",
    "fstype": "pstore",
    "opts": "rw,nosuid,nodev,noexec,relatime"
  },
  {
    "device": "cgroup",
    "mountpoint": "/sys/fs/cgroup/perf_event",
    "fstype": "cgroup",
    "opts": "rw,nosuid,nodev,noexec,relatime,perf_event,release_agent=/run/cgmanager/agents/cgm-release-agent.perf_event"
  },
  {
    "device": "cgroup",
    "mountpoint": "/sys/fs/cgroup/cpuset",
    "fstype": "cgroup",
    "opts": "rw,nosuid,nodev,noexec,relatime,cpuset,clone_children"
  },
  {
    "device": "cgroup",
    "mountpoint": "/sys/fs/cgroup/blkio",
    "fstype": "cgroup",
    "opts": "rw,nosuid,nodev,noexec,relatime,blkio"
  },
  {
    "device": "cgroup",
    "mountpoint": "/sys/fs/cgroup/pids",
    "fstype": "cgroup",
    "opts": "rw,nosuid,nodev,noexec,relatime,pids,release_agent=/run/cgmanager/agents/cgm-release-agent.pids"
  },
  {
    "device": "cgroup",
    "mountpoint": "/sys/fs/cgroup/cpu,cpuacct",
    "fstype": "cgroup",
    "opts": "rw,nosuid,nodev,noexec,relatime,cpu,cpuacct"
  },
  {
    "device": "cgroup",
    "mountpoint": "/sys/fs/cgroup/hugetlb",
    "fstype": "cgroup",
    "opts": "rw,nosuid,nodev,noexec,relatime,hugetlb,release_agent=/run/cgmanager/agents/cgm-release-agent.hugetlb"
  },
  {
    "device": "cgroup",
    "mountpoint": "/sys/fs/cgroup/memory",
    "fstype": "cgroup",
    "opts": "rw,nosuid,nodev,noexec,relatime,memory"
  },
  {
    "device": "cgroup",
    "mountpoint": "/sys/fs/cgroup/freezer",
    "fstype": "cgroup",
    "opts": "rw,nosuid,nodev,noexec,relatime,freezer"
  },
  {
    "device": "cgroup",
    "mountpoint": "/sys/fs/cgroup/net_cls,net_prio",
    "fstype": "cgroup",
    "opts": "rw,nosuid,nodev,noexec,relatime,net_cls,net_prio"
  },
  {
    "device": "cgroup",
    "mountpoint": "/sys/fs/cgroup/devices",
    "fstype": "cgroup",
    "opts": "rw,nosuid,nodev,noexec,relatime,devices"
  },
  {
    "device": "systemd-1",
    "mountpoint": "/proc/sys/fs/binfmt_misc",
    "fstype": "autofs",
    "opts": "rw,relatime,fd=24,pgrp=1,timeout=0,minproto=5,maxproto=5,direct"
  },
  {
    "device": "debugfs",
    "mountpoint": "/sys/kernel/debug",
    "fstype": "debugfs",
    "opts": "rw,relatime"
  },
  {
    "device": "mqueue",
    "mountpoint": "/dev/mqueue",
    "fstype": "mqueue",
    "opts": "rw,relatime"
  },
  {
    "device": "hugetlbfs",
    "mountpoint": "/dev/hugepages",
    "fstype": "hugetlbfs",
    "opts": "rw,relatime"
  },
  {
    "device": "sunrpc",
    "mountpoint": "/run/rpc_pipefs",
    "fstype": "rpc_pipefs",
    "opts": "rw,relatime"
  },
  {
    "device": "fusectl",
    "mountpoint": "/sys/fs/fuse/connections",
    "fstype": "fusectl",
    "opts": "rw,relatime"
  },
  {
    "device": "/dev/sda1",
    "mountpoint": "/boot",
    "fstype": "ext2",
    "opts": "rw,relatime,block_validity,barrier,user_xattr,acl"
  },
  {
    "device": "cgmfs",
    "mountpoint": "/run/cgmanager/fs",
    "fstype": "tmpfs",
    "opts": "rw,relatime,size=100k,mode=755"
  },
  {
    "device": "tmpfs",
    "mountpoint": "/run/user/0",
    "fstype": "tmpfs",
    "opts": "rw,nosuid,nodev,relatime,size=817136k,mode=700"
  },
  {
    "device": "172.16.136.143:/exports/10000sq",
    "mountpoint": "/data/file",
    "fstype": "nfs",
    "opts": "rw,relatime,vers=3,rsize=32768,wsize=32768,namlen=255,hard,proto=tcp,timeo=600,retrans=2,sec=sys,mountaddr=172.16.136.143,mountvers=3,mountport=597,mountproto=udp,local_lock=none,addr=172.16.136.143"
  }
]
```

window 返回案例：
```javascript
[
  {
    "device": "C:",
    "mountpoint": "C:",
    "fstype": "NTFS",
    "opts": "rw.compress"
  },
  {
    "device": "E:",
    "mountpoint": "E:",
    "fstype": "NTFS",
    "opts": "rw.compress"
  }
]
```

##  根据传参目录磁盘使用情况
### 地址： 
> ip:port/disk/path

### 请求方法：
POST

### 参数说明：
path : 目录地址，参数案例 /data/file

### 返回参数：
```javascript
{
  "path": "/data/file",
  "fstype": "nfs",
  "total": 536870912000,
  "free": 420079468544,
  "used": 116791443456,
  "usedPercent": 21.7541015625,
  "inodesTotal": 827138037,
  "inodesUsed": 969507,
  "inodesFree": 826168530,
  "inodesUsedPercent": 0.11721223745390395
}
```