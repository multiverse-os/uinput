<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">

## Multiverse OS: `uinput` library
**URL** [multiverse-os.org](https://multiverse-os.org)

The Linux `uinput` subsystem provides low-level access to the Kernel for the
purposes of creating and manipulating virtual devices. This library is a 
subcomponent of the Multiverse OS Go language library `vinput`, which provides
and abstraction layer ontop of the `uinput` subsystem. This abstraction layer 
is specifically designed to simplfiy creation of virtual devices for the 
purposes of user input automation. 

This library simply focuses on interaction with the `uinput` subsystem and 
provides very little abstraction over the exsiting protocol and ideally is 
designed so that it can be used by a wide range of applications beyond our 
use in `vinput`.

### Usage
Initialization of a uinput device is done using a function chain used to 
modify defaults. 

```
  // Simpliest initialization
  kbd := Keyboard.New("device-name") 

```

(Provide more details as progress continues on vinput which will can be used as
example code and general usage examples) 
