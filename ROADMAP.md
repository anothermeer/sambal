```
Sambal V0.1 [released]
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

Sambal V0.1.1 [released]
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
├── [ ] self tests
├── [ ] mDNS discovery
│   ├── [ ] sambal recv advertiser
│   └── [ ] sambal list discovery
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
├── [ ] end-to-end encryption pipeline
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

OS to support:
- Windows x86_64 (from vista)
- windows ARM/x32 (vista-7)
- Linux x86_64 (kernel>3.1)
- Linux ARM (especially SBC OSes)
- MacOS/iOS/iPadOS
- Blackberry OS
- TempleOS (personal challenge)
[do tell me what other OS I should support ;) ]

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