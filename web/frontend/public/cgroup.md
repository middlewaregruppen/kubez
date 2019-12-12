## Control Groups
Control Groups or cgroups handles Recource limiting, Prioritzation, Accounting and Control of processes on a server.

It's main propuse in a container platform is to limit how much compute power (cpu and ram) a container can use.

In the Kernel a scheduler is responsible for deciding which process that is allowed to run on the CPU. The default scheduler in linux since 2007 is the Completely Fair Scheduler (CFS) and has the goal to maximize the overall CPU utilization while also maximizing interactive performance.


#### CPU: CFS Quota
Resouce Limiting. Shows the maximum time during each cpu cfs period in which the current group (container) will be allowed to run. 

#### CPU: CFS Period 
Resource Limiting. Duration of each scheduler period, for bandwidth decisions.

#### CPU: Nr_Throttled
Bandwidth statistics. Number of times we exausted the full allowed bandwidth.

#### CPU: Throttled_Time 
Bandwidth statistics. Total time the tasks were not run due to being over quota.

#### CPU: Nr_Periods
Bandwidth statistics. How many full periods have been elapsed

#### Memory: Usage
Current memory usage

#### Memory: Limit
Maximum memory that the group (container) is allowed to use

### Using control groups in docker

`docker run --memory=20m --cpus=0.5 containerimage:version` will give the container 20 MB of RAM memory and allows the container to be scheduled maximum 50% of CFS period. If 20 MB is exceeded the container will be killed.  
 
 ### Using control groups in kubernetes
 In the workloads pod template add spec.resources.limits.

```
    
apiVersion: v1/deployment
kind: Deployment
metadata:
  labels:
    app: myapp
  name: myapp
spec:
  replicas: 1
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
        image: containerimage:version
        name: myworkload
        resources:
          limits:
            cpu: 600m
            memory: 256Mi
 ```
