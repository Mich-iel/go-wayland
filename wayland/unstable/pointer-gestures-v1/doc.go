package pointer_gestures

//go:generate go run github.com/rajveermalviya/go-wayland/cmd/go-wayland-scanner -pkg pointer_gestures -prefix zwp -suffix v1 -o pointer_gestures.go -i https://raw.githubusercontent.com/wayland-project/wayland-protocols/1.31/unstable/pointer-gestures/pointer-gestures-unstable-v1.xml
