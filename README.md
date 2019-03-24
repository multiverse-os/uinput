<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">

## Multiverse OS: `uinput` library
**URL** [multiverse-os.org](https://multiverse-os.org)

The Linux `uinput` subsystem provides low-level access to the Kernel for the
purposes of creating and manipulating virtual devices. This library was created
and is maintained to satisfy requirements of the Multiverse OS Go language
library `vinput`. The library `vinput` provides an abstraction layer ontop of
the `uinput` subsystem; specifically designed to simplfiy creation of virtual
devices for the purposes of user input automation. 

It makes sense to provides `uinput` as a separate library, one that simply
focuses on interaction with the `uinput` subsystem and provides very little
abstraction over the existing protocol. So ideally it can be used by a wide
range of applications beyond our use in `vinput`.

### Usage
Initialization of a uinput device is done using a function chain used to 
modify defaults. 

Unlike many other implementations, this one requires connecting the virtual
device after creating it. Which allows the developer to disconnect and 
reconnect the virtual devices as needed.

```
  // Simpliest initialization
  kbd := Keyboard.New("device-name").Connect()

```
