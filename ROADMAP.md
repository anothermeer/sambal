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
│   ├── [ ] progress bar
│   ├── [ ] file size validation
│   ├── [ ] overwrite check
│   └── [ ] checksum verifications
├── [ ] daemon/tray (sambal recv)
├── [ ] better --help
├── [ ] better "sambal version"
├── [ ] changable IP/port
├── [ ] auto device name/id
└── [ ] double click -> GUI

Sambal V0.3
├── [ ] Clipboard sync
├── [ ] large file transfer/streaming
├── [ ] multi-file support
└── [ ] mDNS discovery (sambal list)

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

...still planning...
```