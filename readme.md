```bash
# or just bcb for linux
go build -o bcb.exe
```

modifying the files inside `/assets/` should be all you need to create a new compact book, after setting up how your project should look, render to `out.html` inputing your original source from a `txt` file.
```bash
bcb render ./path/to/your/raw-book.txt
```

before releasing, minify your `out.html` file, you can use [this tool by tdewolff](https://github.com/tdewolff/minify/releases/latest).