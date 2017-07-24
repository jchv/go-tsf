# Go-TSF
Go-TSF is a work-in-progress wrapper for the Microsoft Windows Text Services
Framework library. The TSF is an important part of the Windows operating
system; applications that implement text input can use TSF to provide advanced
integration with IMM, accessibility services, and other services like
handwriting and speech recognition.

## Why?
The main use case I envision is for enabling screen readers and east-asian
input method editors to interact with Go applications.

## About Go-TSF
Go-TSF is a pure Go (no CGO) wrapper around COM objects.

UNIX-like operating systems typically tie the runtime linking functionality
into the C library, which has the annoying side-effect that dynamically
calling into third party libraries requires linking with C, even to just
access the dlopen/dlclose APIs.

Interestingly enough, Windows does not have this limitation and we can freely
call the linking services from the kernel, enabling us to load modules at
will. This means it is possible to load OLE32 and initialize COM objects.
From that point, COM objects have a defined, simple C++-style layout using
vtables.

Therefore, Go-TSF exploits the fact that Go structs are already C-like and
simply contains transcribed versions of the interfaces that wrap the raw
VTables.

Aside from allowing us to eschew running or needing a C compiler, this is also
just a neat box to tick in my opinion :)

I did want to utilize as much of the Go ecosystem as possible, but it didn't
seem worth it to have a w32 dependency just for a few type definitions, and
the go-ole library doesn't seem to be doing quite what we want here. So, for
now, Go-TSF doesn't have any dependencies outside the standard library.

## Code generation?
Currently the code is painstakingly written by hand, scrapping together details
from MSDN as well as the Windows 10 SDK. It would seem like a good opportunity
to move to code generation.

Sadly, it doesn't seem there's a huge benefit for this. Unlike APIs like
OpenGL, there's not a huge demand for generating wrappers for COM APIs, less
TSF specifically. So, if we did the work of making a machine-readable version
of the API with documentation, it would mostly only be used for this project.
And even then, it's not clear the code generation would save significant
amounts of time depending on how close the method call generation could
actually get.

If there were machine readable definitions for COM APIs, this would make sense.
It hits me as nearly insane that Microsoft does not seem to even have this
internally, as there are some really strange inconsistencies in MSDN itself.
Some APIs don't seem to exist as documented and others are said to be
unimplemented but seem to be implemented. As an example, try to find ISoftKbd
outside of MSDN... by all accounts, it doesn't exist. Googling it turns up the
leaked Windows NT 4.0 source code, where you can only find references to
unrelated bits in IMM32. Very confusing.

That being said, there are over 100 interfaces to implement, so it may be
worth it to at least attempt to automate some stubs. The problem is mostly
that parsing Windows headers is masochistically hard :)