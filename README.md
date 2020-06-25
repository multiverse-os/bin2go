[<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">](https://github.com/multiverse-os)

## Multiverse OS: Go Binary Asset Embedding
**URL** [multiverse-os.org](https://multiverse-os.org)

A variety of tools exist to accomplish this task, but this is not just another
fork, and not just a re-implementation; but a re-imagining how assets can be
included in a Go project. 

This library has very simple goals: 

    1) **Security** as with all Multiverse OS libraries is the first
    consideration, and other asset embedding tools fail to supply a way to check
    against existing third parties to ensure the binaries included are not
    tampered with.
  
    **Verifiability** is essential if the item being embedded is a binary, or some
    other object that can easily be verifeid against a third party. 
    
    Additionally, if for example, we want to embed the `ruby` binary in our Go
    application, we should not be relying on the developer to include it, but
    rather use a third party resource like the debian package system or similar
    and simply generate the files based on configurations (perferably YAML or
    ruby based configurations). 
  
    2) **Flexibility** again following Multiverse OS design guidelines, developers
    should not be hard-coding opinions into the library; for example, most
    implementations use `gzip` o compress the data before storing in binary inline
    in the Go source, and decompressing when using it in the application runtime. 
  
    However, Multiverse OS design guidelines, do not enforce developer
    preferences, with only security being except from this design guideline. 
    Decisions for users (developers in the case of libraries) regarding
    compression library should not be hard-coded or enforced, it should be
    flexible and software-definable. There are a variety of reasons for this that
    can be found argued in the Multiverse OS design documentation. 
  
    **Middleware-oriented design** so that not just the compression codec can be
    easily dropped in and replaced, removed or whatever is desired by the
    developer calling the library, but this type of functionality should be
    implemented using a middleware style design pattern common with Go. 
  
    This allows for library developers to implement several different drop-in
    middleware solutions like Reed-solomon/Raptor error checking, encryption based
    on whichever cipher the developer requires. 
  
    3) **Drop all functionality beyond storing data within the binary, and
    recovering data from the binary** so that the domain of this library can
    remain focused. 
  
    Many embedding libraries are very opinionated on how the data once recovered
    from the binary should be made available to the applicaion. However this
    should make no difference, as this should clearly be outside the scope of this
    specific functionality; otherwise the library will be far less modular, and
    will make it unable to take advantage of many existing libraries which provide
    this type of functionality. 
  
    It is a very short sighted design decision to pack in logic to handle how the
    data pulled out of the binary is shared, as it makes a lot of assumptions
    about the final application. 
  
    4) **Simplifity** Incredibly simple, and asset embedding, include the assets within a
    specific folder that is watched, so that there is no expectation of outside
    interaction which can be either confusing, and more importantly, has not yet
    resutled in a standard for asset embedding that is either secure or
    verifiable. 
  
And with just four points, we can identify clear issues with the existing
implementations and identify exactly why a new library is necessary for
Multiverse OS applications requiring embedding binary data. 

Multiverse OS development has also expored using memFD, and implementing an
memexec library that is a drop-in for os/exec allowing memory executation of
binary data. 

#### Not red-team style tools
This library also is not a downloader, a "crypter", or similar red-team style
tools. While it can serve for that if a a separate module was added-on, this
module is outside the scope of this project. But this also goes to show how this
design is more versatile, by combining small modules with it, it can be used to
acheive a wide variety of goals.  

#### Commentary on Go Library Ecosystem
A lot of Go libraries suffer similar problems, of failing to appropriately scope
their project so that it maximally versatile, re-usable, and focused on specific
functionality. Instead often we find libraries sprawling across functionality,
often heavily opinionated due to hard-coding developer preferences, and
therefore result in libraries that are not as broadly useful as they could be,
do not easily upgrade when new technology is released. 

For example, if the compression library component in this case is implemented as 
a middleware-style codec, when new compression libraries come out, instead of 
simply using the previous `gzip` middleware module was a skeleton (or better, 
provide middleware skeletons to speed up development and assist other developers 
working on or around your projects), upgraded to a newly released or popularized 
compression with superprior preformance like `zstd`. 

This allows these new changes be added, without breaking existing projects 
using the library with the older middleware, and provide ways for those 
projects to upgrade, and let new projects use new or experimental compressions. 
But this compression example, is meant to illustrate a wider issue. 

*Our hope is that pointing out these issues, pointing out good libraries that
already exist, creating new examples of libraries that are more broadly usuable
by a wider variety of projects; we can encourage new practices within the
Go library ecosystem that faciliate better collaboration, and cooperation; and
less re-implementing the same code when it can be avoided.* 


#### Contributions
All Multiverse OS projects are open to the public to contribute pull requests,
request support, make suggestions, and feature requests. As a show of
appreciation and mutual respect for developers contributing pull requests to
Multiverse OS projects, code reviews will be provided for any submitted code,
and a discussion generated. 

All contributions are taken seriously, but will be
required to meet the Multiverse OS design minimum guidelines, where they can not
be satisfied, existing Multiverse OS developers can step in to contribute time
and edits to help all submisisons fall in-line with the minimum requirements. 

