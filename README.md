# concourse-resource-library

Collection of Go packages for Concourse Resources

## Resource Bootstrapping

This library contains a simple utility for bootstrapping new resources written in Go. It assumes the directory for the resource has not been created.

```bash
go run main.go new -r git@github.com:digitalocean/artifactory-resource.git -m github.com/digitalocean/artifactory-resource ../artifactory-resource
```

## Credits

Parts of this library is comprised of code which was originated from the `telia-oss/github-pr-resource` project.
