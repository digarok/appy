# Appy

Local project builder for emulator disk images

## Getting Started

Install Appy then create an `appy.yaml` project definition file. 

Run Appy to assemble, build, and run your source.  Or do all 3 with:
```
appy brun
```

### What this solves

We often need to assemble projects locally, but that is only part of the equation for people writing project to run under emulation.  

The second step is taking our assembled binaries, and combining them with any other assets, data files, operating system files, et cetera, into a single disk image for use with an emulator. 

Often this is handled by a set of scripts the author has locally in the projects.  However this leads to problems when collaborating with other people or building pipelines. 

Problem 1: You need to not only provide the scripts, but also the build tools they use, at the correct version.

Problem 2: The scripts often contain bespoke logic for things like handling filetypes or where to place files on the disk image. 

Appy abstracts the tools away from a script mindset, and into a project mindset.  Currently the tool paths are hardcoded but next versions will include dynamic tool providers, including vendoring from internet sources!

## The project file

Currently it uses an `appy.yaml` file in the current project directory. 

### Running Appy

Appy has a verb command structure (sorta like git) for executing the various functions. 

Invoke like: `appy command`

Commands:
```
$ appy asm          # assemble all files

$ appy disk         # create all disk images and add all files

$ appy run          # launch an emulator

$ appy build        # assemble files and make disk, aka 'asm'+'disk'

$ appy brun         # assemble files, make disk, and launch emulator
                    #   aka 'asm'+'disk'+'run'
```

### Version Notes
This is an early experimental version not intended for public use.

Versioning from external sources and custom tool paths not yet implemented but all core functionality exists.
