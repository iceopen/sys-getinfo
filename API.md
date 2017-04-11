# 接口文档说明

## swap 区使用情况
### 地址： 
> ip:port/swap/info

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
> ip:port/cpu/info

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