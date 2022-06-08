## Simple Load Balancer using nginx docker compose and gin

In `main.go` the gin server simply returns the message Hello, World! on the `/` path.

In `nginx.conf` the two servers are assign weights in which way the load is to be distributed. By the `curl` command we can see in the terminal that `app1` is used more times than `app2`.
