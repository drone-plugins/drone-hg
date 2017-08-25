Use the Mercurial plugin to clone a mercurial repository. Note that Drone uses this plugin
by default for all Bitbucket mercurial repositories, without any configuration required. You can override
the default configuration with the following parameters:

* `path` relative path inside /drone/src where the repository is cloned

The following is a sample Git clone configuration in your .drone.yml file:

```yaml
clone:
  hg:
    image: plugins/drone-hg
    path: bitbucket.org/foo/bar
```
