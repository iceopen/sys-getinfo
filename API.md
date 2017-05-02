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

##  CPU 负载查看
### 地址：
> ip:port/load/stat

### 请求方法：
GET

### 返回参数：
```javascript
{
    "load1": 1.99,
    "load5": 2.04,
    "load15": 2.19
}
```

##  网卡信息
### 地址：
> ip:port/net/info

### 请求方法：
GET

### 返回参数：
```javascript
[
{
"mtu": 16384,
"name": "lo0",
"hardwareaddr": "",
"flags": [
"up",
"loopback",
"multicast"
],
"addrs": [
{
"addr": "127.0.0.1/8"
},
{
"addr": "::1/128"
},
{
"addr": "fe80::1/64"
}
]
},
{
"mtu": 1280,
"name": "gif0",
"hardwareaddr": "",
"flags": [
"pointtopoint",
"multicast"
],
"addrs": []
},
{
"mtu": 1280,
"name": "stf0",
"hardwareaddr": "",
"flags": null,
"addrs": []
},
{
"mtu": 1500,
"name": "en0",
"hardwareaddr": "b8:e8:56:46:61:ae",
"flags": [
"up",
"broadcast",
"multicast"
],
"addrs": [
{
"addr": "fe80::1010:8a1d:33c4:2e01/64"
},
{
"addr": "192.168.1.200/24"
}
]
},
{
"mtu": 1500,
"name": "en1",
"hardwareaddr": "72:00:01:2a:86:80",
"flags": [
"up",
"broadcast"
],
"addrs": []
},
{
"mtu": 1500,
"name": "en2",
"hardwareaddr": "72:00:01:2a:86:81",
"flags": [
"up",
"broadcast"
],
"addrs": []
},
{
"mtu": 1500,
"name": "bridge0",
"hardwareaddr": "72:00:01:2a:86:80",
"flags": [
"up",
"broadcast",
"multicast"
],
"addrs": []
},
{
"mtu": 2304,
"name": "p2p0",
"hardwareaddr": "0a:e8:56:46:61:ae",
"flags": [
"up",
"broadcast",
"multicast"
],
"addrs": []
},
{
"mtu": 1484,
"name": "awdl0",
"hardwareaddr": "9a:19:37:64:7e:5e",
"flags": [
"up",
"broadcast",
"multicast"
],
"addrs": [
{
"addr": "fe80::9819:37ff:fe64:7e5e/64"
}
]
},
{
"mtu": 2000,
"name": "utun0",
"hardwareaddr": "",
"flags": [
"up",
"pointtopoint",
"multicast"
],
"addrs": [
{
"addr": "fe80::e250:ac0f:d136:4ea4/64"
}
]
}
]
```

##  CPU 负载查看
### 地址：
> ip:port/process/info

### 请求方法：
GET

### 返回参数：
```javascript
[1,54,55,57,58,59,61,64,65,66,67,69,70,75,77,80,82,83,88,89,91,93,94,100,101,103,104,105,106,107,108,110,112,113,114,115,118,120,121,122,123,124,126,127,128,129,130,133,135,143,168,169,171,180,185,186,188,199,200,204,213,216,218,308,309,312,314,315,316,317,318,319,320,326,327,328,329,335,336,338,353,354,371,372,374,375,376,377,378,381,382,385,386,390,399,403,404,405,414,416,417,418,420,422,423,424,426,427,428,429,433,434,436,437,438,440,441,442,443,444,445,446,452,454,456,459,460,461,462,464,465,467,468,471,473,474,475,476,478,479,481,482,483,484,485,486,487,488,489,490,492,495,496,497,498,499,500,503,504,505,507,508,509,510,511,518,520,524,525,526,527,529,530,531,532,533,535,536,538,541,542,544,545,546,551,552,553,554,558,559,560,561,565,568,569,570,571,572,578,581,585,590,600,601,602,629,630,634,648,654,656,657,658,659,661,662,664,673,676,677,678,679,684,685,697,772,773,775,776,777,782,783,784,785,786,787,792,907,908,949,999,1176,1305,1306,1311,1314,1529,1967,2049,2096,2302,3327,4196,4197,4546,4665,4754,4755,4756,4885,4894,6163,8546,8547,8550,8552,9688,9958,9970,10959,10960,10963,10964,11418,11420,11433,11744,11846,11976,12353,12800,13330,13341,13342,14085,14722,15195,15473,16311,16457,16792,24820,24871,30092,31747,38492,42348,43723,52713,52770,52772,53240,53242,53243,53244,56206,59031,59043,59045,59337,59997,60204,60205,61610,68694,69770,69791,69792,70436,84070,84113,85402,85641,86393,86741,89172,89214,90047,92546,14256,16321,16358,16468]
```