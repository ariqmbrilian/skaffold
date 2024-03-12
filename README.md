source: 
- https://skaffold.dev/docs/quickstart/
- https://dev.to/thenjdevopsguy/getting-started-with-skaffold-5120

### Bootstrap Skaffold configuration
```
skaffold init
```

### Use Skaffold for continuous development
```
skaffold dev
```

### Use Skaffold for continuous integration
#### Build an image
```
export STATE=$(git rev-list -1 HEAD --abbrev-commit)
skaffold build --file-output build-$STATE.json
```

### Test an image
```
skaffold test --build-artifacts build-$STATE.json
```

### Use Skaffold for continuous delivery
#### Deploy in a single step
```
skaffold deploy -a build-$STATE.json
```

### Render and apply in separate steps
#### Run the following command to render a hydrated manifest:
```
skaffold render -a build-$STATE.json --output render.yaml --digest-source local
```

### Next, run the following command to apply your hydrated manifest
```
skaffold apply render.yaml
```

### Delete deployment
```
skaffold delete
```