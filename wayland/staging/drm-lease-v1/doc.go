package drm_lease

//go:generate go run github.com/rajveermalviya/go-wayland/go-wayland-scanner -pkg drm_lease -prefix wp -suffix v1 -o drm_lease.go -i https://raw.githubusercontent.com/wayland-project/wayland-protocols/1.26/staging/drm-lease/drm-lease-v1.xml
