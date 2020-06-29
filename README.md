# spotify-dbus-infos
Simply display spotify infos.
Use font awesome icons.

## Example with polybar config

```
modules-center = spotify

;----------------------------------------------------------
;                SPOTIFY
;----------------------------------------------------------

[module/spotify]
type = custom/script
interval = 1
format = "<label>"
exec = spotify-dbus-infos
format-underline = #1db954
```
