
Prerequisites:
`sudo apt install rlwrap`.
Java, installed from https://adoptium.net/installation/linux

Installed, following instructions from [here](https://clojure.org/guides/install_clojure):
```
$ curl -O https://download.clojure.org/install/linux-install-1.11.1.1200.sh
$ chmod +x linux-install-1.11.1.1200.sh
$ sudo ./linux-install-1.11.1.1200.sh
```


`sudo apt install leiningen` - https://leiningen.org/, a package manager(?).

```
$ mkdir ~/.lein
$ cat ~/.lein/profiles.clj
{:user {:plugins [[cider/cider-nrepl "0.28.5"]]}}
```

and then `lein repl` starts `cider-nrepl`:

```
$ lein repl
Warning: cider-nrepl requires Leiningen 2.8.3 or greater.
Warning: cider-nrepl requires Clojure 1.8 or greater.
Warning: cider-nrepl will not be included in your project.
Warning: cider-nrepl requires Leiningen 2.8.3 or greater.
Warning: cider-nrepl requires Clojure 1.8 or greater.
Warning: cider-nrepl will not be included in your project.
nREPL server started on port 34877 on host 127.0.0.1 - nrepl://127.0.0.1:34877
WARNING: cat already refers to: #'clojure.core/cat in namespace: net.cgrand.regex, being replaced by: #'net.cgrand.regex/cat
REPL-y 0.3.7, nREPL
Clojure 1.8.0
OpenJDK 64-Bit Server VM 1.8.0_352-8u352-ga-1~18.04-b08
    Docs: (doc function-name-here)
          (find-doc "part-of-name-here")
  Source: (source function-name-here)
 Javadoc: (javadoc java-object-or-class-here)
    Exit: Control+D or (exit) or (quit)
 Results: Stored in vars *1, *2, *3, an exception in *e

user=>
```

Install Fireplace plugin: https://github.com/tpope/vim-fireplace

With `lein repl` running, say `:FireplaceConnect` in vim, and pass the nrepl
URL from lein repl output, e.g.: `nrepl://127.0.0.1:34877`.

