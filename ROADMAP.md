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
├── [ ] auto device name
└── [ ] double click -> GUI

Sambal V0.3
├── [ ] Clipboard sync
├── [ ] large file transfer/streaming
├── [ ] multi-file support
├── [ ] mDNS discovery (sambal list)

planned commands:
sambal version
sambal send <file>
sambal recv
sambal text <text>
sambal clip
sambal ping

planned flags:
--version
--help
--debug
--console
--port <port>
--tcpip <ip addr>
--httpip <ip addr>

...still planning...
```