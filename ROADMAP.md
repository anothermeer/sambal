```
Sambal V0.1
├── [/] CLI (cobra)
├── [/] GUI (fyne, temporary)
├── [/] server/receive
├── [/] client/send
├── [/] handshake protocol
│   ├── [/] HELLO
│   ├── [/] FILE_OFFER
│   ├── [/] TEXT_SEND (template)
│   └── [/] CLIPBOARD_SYNC (template)
├── [/] tidy commands and functions
└── [/] JSON message system

Sambal V0.1.1
├── [/] real HTTP file transfer
│   ├── [/] progress bar
│   ├── [/] file size validation
│   ├── [/] overwrite check
│   └── [/] checksum verifications
├── [/] protocol version compatibility
│   ├── [/] minimum supported version
│   └── [/] incompatible version warning
├── [/] better --help
├── [/] auto device name/id
├── [/] large file transfer/streaming
└── [/] double click -> temporary GUI

Sambal V0.2
├── [ ] transfer logs
├── [ ] mDNS discovery (sambal list)
├── [ ] better connection handling
│   ├── [ ] pause transfer
│   ├── [ ] resume transfer
│   └── [ ] timeouts
├── [ ] multi-file support
│   └── [ ] transfer IDs
└── [ ] daemon/tray (sambal recv)

Planned features
├── [ ] main window
├── [ ] Clipboard sync
├── [ ] transfer logs
├── [ ] better "sambal version"
├── [ ] auto update
└── [ ] settings flags
    ├── [ ] configurable port
    │   ├── [ ] port conflict warning
    │   └── [ ] port in use warning
    ├── [ ] daemon configs
    │   ├── [ ] configurable timeout
    │   ├── [ ] auto accept on/off
    │   └── [ ] discover/transfer on/off
    └── [ ] targetted device

Future plans
├── [ ] Android support
├── [ ] iOS support
├── [ ] Windows & MacOS signing
├── [ ] More Linux distro compatibility
└── [ ] Better GUI

planned commands:
sambal version
sambal send <file>
sambal recv
sambal text <text>
sambal ping
sambal clip
sambal list

planned flags:
--version
--help
--debug
--console
--tcpport <port>
--httpport <port>
--ip <ip addr>
--fileloc <path>

(maybe open a Github Sponsors??)
...still planning...
```