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
├── [ ] tidy commands and functions
└── [/] JSON message system

Sambal V0.2
├── [ ] real HTTP file transfer
│   ├── [ ] progress bar
│   ├── [ ] file size validation
│   ├── [ ] overwrite check
│   └── [ ] checksum verifications
├── [ ] large file transfer/streaming
├── [ ] multi-file support
└── [ ] mDNS discovery (sambal list)

Sambal V0.3
└── [ ] HTTPS file transfer?

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