# vinput uinput sub-package task list
===============================================================================

// TODO: Emulate/Model the virtual keyboard in a more realistic way;
//       such as, add a boolean to track <capslock>, <numlock>, and
//       status of other similar toggle keys.
//
// TODO: Support simple multi-key pressing
//
// TODO: Support Unicode
//
// TODO: Support simpler API allowing TypeString("string to type")...
//
// TODO: Support macro definition, as one could with a more advanced
//       keyboard
//
// TODO: Begin to sketch out some of the additional hardware functionality
//       in the Multiverse open source hardware HID spec

// Keep logic in uinput releated to uninput, logic that abstracts
// and keeps track of window/process being interacted with, should
// be at the higher level of virtui

// TODO: Instead of taking in Int perhaps take in KeyCode type
