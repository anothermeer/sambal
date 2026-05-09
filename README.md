# Sambal
> An easy-to-use tool for sharing files and text across devices using local network.

Fun Fact: the name `Sambal` is actually inspired from the file transfer tool, samba, and Malaysia's iconic local cuisine, nasi lemak, which has a chili paste called sambal.

## Building Requirements
- Go = 1.26.2
- Cobra >= 1.10.2
- Fyne >= 2.7.3
- internet connection
- you

## Installation (I actually don't know how)
Windows:
```powershell
git clone https://github.com/anothermeer/sambal
cd sambal
go build . -o sambal.exe
```
Linux/Unix:
```bash
sudo apt update
sudo apt install gcc libgl1-mesa-dev xorg-dev libwayland-dev libxkbcommon-dev
git clone https://github.com/anothermeer/sambal
cd sambal
go build . -o sambal
```

## Troubleshooting

### Linux build error:
`build constraints exclude all Go files`

Install OpenGL/X11 development libraries:

```bash
sudo apt install gcc libgl1-mesa-dev xorg-dev
```