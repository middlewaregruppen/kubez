# Welcome to Dr. Kubez

Dr. Kubez is a tool to test, diagnose and load k8s clusters.

## Features

**Load**
- Load the API Server with gazillions of calls to fill up the etcd database.
- Max out the CPU
- Allocate RAM-memory 

**Diagnostics**
- View information about CGroups.
- View HTTP Headers that comes in to the container.
- TCP/IP Connection test

**Test**
- Create services that mock endpoints and introduces errors and delays in them.

## Upcoming Feature
- Hijack k8s services and introduce random errors and delays to them. 
- DNS Lookup from inside the cluster.