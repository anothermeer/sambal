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

Sambal V0.2
├── [ ] real HTTP file transfer
│   ├── [/] progress bar
│   ├── [/] file size validation
│   ├── [ ] overwrite check
│   └── [ ] checksum verifications
├── [ ] daemon/tray (sambal recv)
├── [ ] better connection handling
│   ├── [ ] pause transfer
│   ├── [ ] resume transfer
│   └── [ ] timeouts
├── [ ] protocol version compatibility
│   ├── [ ] minimum supported version
│   └── [ ] incompatible version warning
├── [/] better --help
├── [ ] better "sambal version"
├── [ ] auto update
├── [ ] auto device name/id
└── [/] double click -> temporary GUI

Sambal V0.3
├── [ ] main window + settings
├── [ ] Clipboard sync
├── [ ] transfer logs
├── [ ] large file transfer/streaming
├── [ ] multi-file support
│   └── [ ] transfer IDs
├── [ ] settings flags
│   ├── [ ] configurable port
│   │   ├── [ ] port conflict warning
│   │   └── [ ] port in use warning
│   ├── [ ] daemon configs
│   │   ├── [ ] configurable timeout
│   │   ├── [ ] auto accept on/off
│   │   └── [ ] discover/transfer on/off
│   └── [ ] targetted device
└── [ ] mDNS discovery (sambal list)

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