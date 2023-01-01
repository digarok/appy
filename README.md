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

Currently it uses an `appy.yaml` file in the current project directory. The format is as follows:
```
assemble: [main.s, grafix.s, snd.s, a.s, b.s]   # <--- list of files to assemble with Merlin   
indent: [c.s]                                   # <--- additional files to indent when running `appy fmt`               
disks:                                          # <--- define disks, can be more than one, handy for 140K + 800K
- name: mydiskimage                             #   <---- each disk has a name (ProDOS volume name)
  file: mydiskimage800.2mg                      #   <---- each disk has a filename for the image it creates
  size: 800KB                                   #   <---- each disk has a size (valid Cadius sizes)
  files:                                        #   <---- each disk has zero or more files
  - input: ../PRODOS.2.4.2/PRODOS               #     <---- each file has an input location (on your host system)
    output: /mydiskimage                        #     <----       ... and an output location (on the ProDOS disk image)
  - input: testsrc/sp.s
    output: /mydiskimage
  - input: ../modsearch/potential/eclipse.ntp
    output: /mydiskimage
    ftaux: #B30000                              #     <----       ... and option filetype/auxtype info (not sure if working)
```

## User Overrides
What if your copy of Merlin32 (assembler) is in a different location than your teammate's?  You can set up local binary overrides with an `appy.user.yaml` file in the same directory.  It allows the following 3 program settings:
```
# local system settings/overrides
programs:
  merlin32: 'G:\My Drive\appleiigs\tools\merlin32-windows-v1.1.10\Merlin32.exe'
  cadius: 'G:\My Drive\appleiigs\tools\Cadius.exe'
  gsplus: 'myproj\gsplus'
```

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

$ appy fmt          # format/indent your assembly file in appy.yaml
                    # you can also pass in filename(s) to format
```

### Version Notes
This is an early experimental version not intended for public use.

Versioning/vendoring binaries from external sources not yet implemented but all core functionality exists including local binary overrides.


### Dev Quickstart

1. Checkout
2. Run `go mod tidy` to install dependencies
3. Run `go build & go install` to build and install