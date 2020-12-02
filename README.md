# Entemplate


Entemplate (a friendly fork of [Extemplate](https://github.com/dannyvankooten/extemplate)) is a small wrapper package around [html/template](https://golang.org/pkg/html/template/) to allow for easy file-based template inheritance.

File: `templates/parent.tmpl`
```text
<html>
<head>
	<title>{{ block "title" }}Default title{{ end }}</title>
</head>
<body>
	{{ block "content" }}Default content{{ end }}
</body>
</html>
```

File: `templates/child.tmpl`
```text
{{ extends "parent.tmpl" }}
{{ define "title" }}Child title{{ end }}
{{ define "content" }}Hello world!{{ end }}
```

File: `main.go`
```go
xt := entemplate.New()
xt.ParseDir("templates/", []string{".tmpl"})
_ = xt.ExecuteTemplate(os.Stdout, "child.tmpl", "no data needed")
// Output: <html>.... Hello world! ....</html>
```

Entemplate recursively walks all files in the given directory and will parse the files matching the given extensions as a template. Templates are named by path and filename, relative to the root directory.

For example, calling `ParseDir("templates/", []string{".tmpl"})` on the following directory structure:

```text
templates/
  |__ admin/
  |      |__ index.tmpl
  |      |__ edit.tmpl
  |__ index.tmpl
```

Will result in the following templates:

```text
admin/index.tmpl
admin/edit.tmpl
index.tmpl
```

Check out the [tests](https://github.com/kalafut/entemplate/blob/master/template_test.go) and [examples directory](https://github.com/dannyvankooten/extemplate/tree/master/examples) for more examples.

### Developer Conveniences

- Enable automatic template reloading with `xt.AutoReload(true)`. When enabled, the template folder will be reparsed
on execution to allow rapid iteration of templates.

- A bit of syntactic sugar was added to the standard template execution function. In addition to passing in a
single data object (often a `map[string]interface{}`), you may also provide key/value pairs directly:

```
xt.ExecuteTemplate(os.Stdout, "child.tmpl", "foo", 42, "bar", "some value")
```

### Static Templates

The contents of your template folder may be bundled into the application binary directly to allow easier
distribution. The included `enstatic` application will create a `.go` file from the specified template folder:

```
enstatic -p my_package_name my_templates
```

When the resulting `.go` file is built with your application, the bundled template data will be used during
`ParseDir` instead of reading the file system.

See `enstatic` help for configuration options.

### Benchmarks

You will most likely never have to worry about performance, when using this package properly.
The benchmarks are purely listed here so we have a place to keep track of progress.

```
BenchmarkEntemplateGetLayoutForTemplate-8   	 2000000	       923 ns/op	     104 B/op	       3 allocs/op
BenchmarkEntemplateParseDir-8               	    5000	    227898 ns/op	   34864 B/op	     325 allocs/op
```

### License

MIT
