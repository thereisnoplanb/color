# color

A Go package for representing and working with colors in the ARGB color space.

## Installation

```sh
go get github.com/thereisnoplanb/color
```

## Overview

The `Color` type is a `uint32` value whose bits encode the four 8-bit ARGB channels:

| bits 31–24 | bits 23–16 | bits 15–8 | bits 7–0 |
|:---:|:---:|:---:|:---:|
| Alpha | Red | Green | Blue |

## Predefined colors

Over 140 named constants are provided, matching the standard .NET / CSS named-color set.

<details>
<summary>Full color table (141 colors)</summary>

| Color | Hex | Color | Hex | Color | Hex |
|:------|:----|:------|:----|:------|:----|
| Transparent ¹ | `#FFFFFFFF` | <img src="doc/colors/AliceBlue.png" height="16" style="vertical-align:middle"> AliceBlue | `#FFF0F8FF` | <img src="doc/colors/AntiqueWhite.png" height="16" style="vertical-align:middle"> AntiqueWhite | `#FFFAEBD7` |
| <img src="doc/colors/Aqua.png" height="16" style="vertical-align:middle"> Aqua | `#FF00FFFF` | <img src="doc/colors/Aquamarine.png" height="16" style="vertical-align:middle"> Aquamarine | `#FF7FFFD4` | <img src="doc/colors/Azure.png" height="16" style="vertical-align:middle"> Azure | `#FFF0FFFF` |
| <img src="doc/colors/Beige.png" height="16" style="vertical-align:middle"> Beige | `#FFF5F5DC` | <img src="doc/colors/Bisque.png" height="16" style="vertical-align:middle"> Bisque | `#FFFFE4C4` | <img src="doc/colors/Black.png" height="16" style="vertical-align:middle"> Black | `#FF000000` |
| <img src="doc/colors/BlanchedAlmond.png" height="16" style="vertical-align:middle"> BlanchedAlmond | `#FFFFEBCD` | <img src="doc/colors/Blue.png" height="16" style="vertical-align:middle"> Blue | `#FF0000FF` | <img src="doc/colors/BlueViolet.png" height="16" style="vertical-align:middle"> BlueViolet | `#FF8A2BE2` |
| <img src="doc/colors/Brown.png" height="16" style="vertical-align:middle"> Brown | `#FFA52A2A` | <img src="doc/colors/BurlyWood.png" height="16" style="vertical-align:middle"> BurlyWood | `#FFDEB887` | <img src="doc/colors/CadetBlue.png" height="16" style="vertical-align:middle"> CadetBlue | `#FF5F9EA0` |
| <img src="doc/colors/Chartreuse.png" height="16" style="vertical-align:middle"> Chartreuse | `#FF7FFF00` | <img src="doc/colors/Chocolate.png" height="16" style="vertical-align:middle"> Chocolate | `#FFD2691E` | <img src="doc/colors/Coral.png" height="16" style="vertical-align:middle"> Coral | `#FFFF7F50` |
| <img src="doc/colors/CornflowerBlue.png" height="16" style="vertical-align:middle"> CornflowerBlue | `#FF6495ED` | <img src="doc/colors/Cornsilk.png" height="16" style="vertical-align:middle"> Cornsilk | `#FFFFF8DC` | <img src="doc/colors/Crimson.png" height="16" style="vertical-align:middle"> Crimson | `#FFDC143C` |
| <img src="doc/colors/Cyan.png" height="16" style="vertical-align:middle"> Cyan | `#FF00FFFF` | <img src="doc/colors/DarkBlue.png" height="16" style="vertical-align:middle"> DarkBlue | `#FF00008B` | <img src="doc/colors/DarkCyan.png" height="16" style="vertical-align:middle"> DarkCyan | `#FF008B8B` |
| <img src="doc/colors/DarkGoldenrod.png" height="16" style="vertical-align:middle"> DarkGoldenrod | `#FFB8860B` | <img src="doc/colors/DarkGray.png" height="16" style="vertical-align:middle"> DarkGray | `#FFA9A9A9` | <img src="doc/colors/DarkGreen.png" height="16" style="vertical-align:middle"> DarkGreen | `#FF006400` |
| <img src="doc/colors/DarkKhaki.png" height="16" style="vertical-align:middle"> DarkKhaki | `#FFBDB76B` | <img src="doc/colors/DarkMagenta.png" height="16" style="vertical-align:middle"> DarkMagenta | `#FF8B008B` | <img src="doc/colors/DarkOliveGreen.png" height="16" style="vertical-align:middle"> DarkOliveGreen | `#FF556B2F` |
| <img src="doc/colors/DarkOrange.png" height="16" style="vertical-align:middle"> DarkOrange | `#FFFF8C00` | <img src="doc/colors/DarkOrchid.png" height="16" style="vertical-align:middle"> DarkOrchid | `#FF9932CC` | <img src="doc/colors/DarkRed.png" height="16" style="vertical-align:middle"> DarkRed | `#FF8B0000` |
| <img src="doc/colors/DarkSalmon.png" height="16" style="vertical-align:middle"> DarkSalmon | `#FFE9967A` | <img src="doc/colors/DarkSeaGreen.png" height="16" style="vertical-align:middle"> DarkSeaGreen | `#FF8FBC8F` | <img src="doc/colors/DarkSlateBlue.png" height="16" style="vertical-align:middle"> DarkSlateBlue | `#FF483D8B` |
| <img src="doc/colors/DarkSlateGray.png" height="16" style="vertical-align:middle"> DarkSlateGray | `#FF2F4F4F` | <img src="doc/colors/DarkTurquoise.png" height="16" style="vertical-align:middle"> DarkTurquoise | `#FF00CED1` | <img src="doc/colors/DarkViolet.png" height="16" style="vertical-align:middle"> DarkViolet | `#FF9400D3` |
| <img src="doc/colors/DeepPink.png" height="16" style="vertical-align:middle"> DeepPink | `#FFFF1493` | <img src="doc/colors/DeepSkyBlue.png" height="16" style="vertical-align:middle"> DeepSkyBlue | `#FF00BFFF` | <img src="doc/colors/DimGray.png" height="16" style="vertical-align:middle"> DimGray | `#FF696969` |
| <img src="doc/colors/DodgerBlue.png" height="16" style="vertical-align:middle"> DodgerBlue | `#FF1E90FF` | <img src="doc/colors/Firebrick.png" height="16" style="vertical-align:middle"> Firebrick | `#FFB22222` | <img src="doc/colors/FloralWhite.png" height="16" style="vertical-align:middle"> FloralWhite | `#FFFFFAF0` |
| <img src="doc/colors/ForestGreen.png" height="16" style="vertical-align:middle"> ForestGreen | `#FF228B22` | <img src="doc/colors/Fuchsia.png" height="16" style="vertical-align:middle"> Fuchsia | `#FFFF00FF` | <img src="doc/colors/Gainsboro.png" height="16" style="vertical-align:middle"> Gainsboro | `#FFDCDCDC` |
| <img src="doc/colors/GhostWhite.png" height="16" style="vertical-align:middle"> GhostWhite | `#FFF8F8FF` | <img src="doc/colors/Gold.png" height="16" style="vertical-align:middle"> Gold | `#FFFFD700` | <img src="doc/colors/Goldenrod.png" height="16" style="vertical-align:middle"> Goldenrod | `#FFDAA520` |
| <img src="doc/colors/Gray.png" height="16" style="vertical-align:middle"> Gray | `#FF808080` | <img src="doc/colors/Green.png" height="16" style="vertical-align:middle"> Green | `#FF008000` | <img src="doc/colors/GreenYellow.png" height="16" style="vertical-align:middle"> GreenYellow | `#FFADFF2F` |
| <img src="doc/colors/Honeydew.png" height="16" style="vertical-align:middle"> Honeydew | `#FFF0FFF0` | <img src="doc/colors/HotPink.png" height="16" style="vertical-align:middle"> HotPink | `#FFFF69B4` | <img src="doc/colors/IndianRed.png" height="16" style="vertical-align:middle"> IndianRed | `#FFCD5C5C` |
| <img src="doc/colors/Indigo.png" height="16" style="vertical-align:middle"> Indigo | `#FF4B0082` | <img src="doc/colors/Ivory.png" height="16" style="vertical-align:middle"> Ivory | `#FFFFFFF0` | <img src="doc/colors/Khaki.png" height="16" style="vertical-align:middle"> Khaki | `#FFF0E68C` |
| <img src="doc/colors/Lavender.png" height="16" style="vertical-align:middle"> Lavender | `#FFE6E6FA` | <img src="doc/colors/LavenderBlush.png" height="16" style="vertical-align:middle"> LavenderBlush | `#FFFFF0F5` | <img src="doc/colors/LawnGreen.png" height="16" style="vertical-align:middle"> LawnGreen | `#FF7CFC00` |
| <img src="doc/colors/LemonChiffon.png" height="16" style="vertical-align:middle"> LemonChiffon | `#FFFFFACD` | <img src="doc/colors/LightBlue.png" height="16" style="vertical-align:middle"> LightBlue | `#FFADD8E6` | <img src="doc/colors/LightCoral.png" height="16" style="vertical-align:middle"> LightCoral | `#FFF08080` |
| <img src="doc/colors/LightCyan.png" height="16" style="vertical-align:middle"> LightCyan | `#FFE0FFFF` | <img src="doc/colors/LightGoldenrodYellow.png" height="16" style="vertical-align:middle"> LightGoldenrod | `#FFFAFAD2` | <img src="doc/colors/LightGray.png" height="16" style="vertical-align:middle"> LightGray | `#FFD3D3D3` |
| <img src="doc/colors/LightGreen.png" height="16" style="vertical-align:middle"> LightGreen | `#FF90EE90` | <img src="doc/colors/LightPink.png" height="16" style="vertical-align:middle"> LightPink | `#FFFFB6C1` | <img src="doc/colors/LightSalmon.png" height="16" style="vertical-align:middle"> LightSalmon | `#FFFFA07A` |
| <img src="doc/colors/LightSeaGreen.png" height="16" style="vertical-align:middle"> LightSeaGreen | `#FF20B2AA` | <img src="doc/colors/LightSkyBlue.png" height="16" style="vertical-align:middle"> LightSkyBlue | `#FF87CEFA` | <img src="doc/colors/LightSlateGray.png" height="16" style="vertical-align:middle"> LightSlateGray | `#FF778899` |
| <img src="doc/colors/LightSteelBlue.png" height="16" style="vertical-align:middle"> LightSteelBlue | `#FFB0C4DE` | <img src="doc/colors/LightYellow.png" height="16" style="vertical-align:middle"> LightYellow | `#FFFFFFE0` | <img src="doc/colors/Lime.png" height="16" style="vertical-align:middle"> Lime | `#FF00FF00` |
| <img src="doc/colors/LimeGreen.png" height="16" style="vertical-align:middle"> LimeGreen | `#FF32CD32` | <img src="doc/colors/Linen.png" height="16" style="vertical-align:middle"> Linen | `#FFFAF0E6` | <img src="doc/colors/Magenta.png" height="16" style="vertical-align:middle"> Magenta | `#FFFF00FF` |
| <img src="doc/colors/Maroon.png" height="16" style="vertical-align:middle"> Maroon | `#FF800000` | <img src="doc/colors/MediumAquamarine.png" height="16" style="vertical-align:middle"> MediumAquamarine | `#FF66CDAA` | <img src="doc/colors/MediumBlue.png" height="16" style="vertical-align:middle"> MediumBlue | `#FF0000CD` |
| <img src="doc/colors/MediumOrchid.png" height="16" style="vertical-align:middle"> MediumOrchid | `#FFBA55D3` | <img src="doc/colors/MediumPurple.png" height="16" style="vertical-align:middle"> MediumPurple | `#FF9370DB` | <img src="doc/colors/MediumSeaGreen.png" height="16" style="vertical-align:middle"> MediumSeaGreen | `#FF3CB371` |
| <img src="doc/colors/MediumSlateBlue.png" height="16" style="vertical-align:middle"> MediumSlateBlue | `#FF7B68EE` | <img src="doc/colors/MediumSpringGreen.png" height="16" style="vertical-align:middle"> MediumSpringGreen | `#FF00FA9A` | <img src="doc/colors/MediumTurquoise.png" height="16" style="vertical-align:middle"> MediumTurquoise | `#FF48D1CC` |
| <img src="doc/colors/MediumVioletRed.png" height="16" style="vertical-align:middle"> MediumVioletRed | `#FFC71585` | <img src="doc/colors/MidnightBlue.png" height="16" style="vertical-align:middle"> MidnightBlue | `#FF191970` | <img src="doc/colors/MintCream.png" height="16" style="vertical-align:middle"> MintCream | `#FFF5FFFA` |
| <img src="doc/colors/MistyRose.png" height="16" style="vertical-align:middle"> MistyRose | `#FFFFE4E1` | <img src="doc/colors/Moccasin.png" height="16" style="vertical-align:middle"> Moccasin | `#FFFFE4B5` | <img src="doc/colors/NavajoWhite.png" height="16" style="vertical-align:middle"> NavajoWhite | `#FFFFDEAD` |
| <img src="doc/colors/Navy.png" height="16" style="vertical-align:middle"> Navy | `#FF000080` | <img src="doc/colors/OldLace.png" height="16" style="vertical-align:middle"> OldLace | `#FFFDF5E6` | <img src="doc/colors/Olive.png" height="16" style="vertical-align:middle"> Olive | `#FF808000` |
| <img src="doc/colors/OliveDrab.png" height="16" style="vertical-align:middle"> OliveDrab | `#FF6B8E23` | <img src="doc/colors/Orange.png" height="16" style="vertical-align:middle"> Orange | `#FFFFA500` | <img src="doc/colors/OrangeRed.png" height="16" style="vertical-align:middle"> OrangeRed | `#FFFF4500` |
| <img src="doc/colors/Orchid.png" height="16" style="vertical-align:middle"> Orchid | `#FFDA70D6` | <img src="doc/colors/PaleGoldenrod.png" height="16" style="vertical-align:middle"> PaleGoldenrod | `#FFEEE8AA` | <img src="doc/colors/PaleGreen.png" height="16" style="vertical-align:middle"> PaleGreen | `#FF98FB98` |
| <img src="doc/colors/PaleTurquoise.png" height="16" style="vertical-align:middle"> PaleTurquoise | `#FFAFEEEE` | <img src="doc/colors/PaleVioletRed.png" height="16" style="vertical-align:middle"> PaleVioletRed | `#FFDB7093` | <img src="doc/colors/PapayaWhip.png" height="16" style="vertical-align:middle"> PapayaWhip | `#FFFFEFD5` |
| <img src="doc/colors/PeachPuff.png" height="16" style="vertical-align:middle"> PeachPuff | `#FFFFDAB9` | <img src="doc/colors/Peru.png" height="16" style="vertical-align:middle"> Peru | `#FFCD853F` | <img src="doc/colors/Pink.png" height="16" style="vertical-align:middle"> Pink | `#FFFFC0CB` |
| <img src="doc/colors/Plum.png" height="16" style="vertical-align:middle"> Plum | `#FFDDA0DD` | <img src="doc/colors/PowderBlue.png" height="16" style="vertical-align:middle"> PowderBlue | `#FFB0E0E6` | <img src="doc/colors/Purple.png" height="16" style="vertical-align:middle"> Purple | `#FF800080` |
| <img src="doc/colors/Red.png" height="16" style="vertical-align:middle"> Red | `#FFFF0000` | <img src="doc/colors/RosyBrown.png" height="16" style="vertical-align:middle"> RosyBrown | `#FFBC8F8F` | <img src="doc/colors/RoyalBlue.png" height="16" style="vertical-align:middle"> RoyalBlue | `#FF4169E1` |
| <img src="doc/colors/SaddleBrown.png" height="16" style="vertical-align:middle"> SaddleBrown | `#FF8B4513` | <img src="doc/colors/Salmon.png" height="16" style="vertical-align:middle"> Salmon | `#FFFA8072` | <img src="doc/colors/SandyBrown.png" height="16" style="vertical-align:middle"> SandyBrown | `#FFF4A460` |
| <img src="doc/colors/SeaGreen.png" height="16" style="vertical-align:middle"> SeaGreen | `#FF2E8B57` | <img src="doc/colors/SeaShell.png" height="16" style="vertical-align:middle"> SeaShell | `#FFFFF5EE` | <img src="doc/colors/Sienna.png" height="16" style="vertical-align:middle"> Sienna | `#FFA0522D` |
| <img src="doc/colors/Silver.png" height="16" style="vertical-align:middle"> Silver | `#FC0C0C0` | <img src="doc/colors/SkyBlue.png" height="16" style="vertical-align:middle"> SkyBlue | `#FF87CEEB` | <img src="doc/colors/SlateBlue.png" height="16" style="vertical-align:middle"> SlateBlue | `#F6A5ACD` |
| <img src="doc/colors/SlateGray.png" height="16" style="vertical-align:middle"> SlateGray | `#FF708090` | <img src="doc/colors/Snow.png" height="16" style="vertical-align:middle"> Snow | `#FFFFFAFA` | <img src="doc/colors/SpringGreen.png" height="16" style="vertical-align:middle"> SpringGreen | `#F00FF7F` |
| <img src="doc/colors/SteelBlue.png" height="16" style="vertical-align:middle"> SteelBlue | `#FF4682B4` | <img src="doc/colors/Tan.png" height="16" style="vertical-align:middle"> Tan | `#FFD2B48C` | <img src="doc/colors/Teal.png" height="16" style="vertical-align:middle"> Teal | `#FF008080` |
| <img src="doc/colors/Thistle.png" height="16" style="vertical-align:middle"> Thistle | `#FFD8BFD8` | <img src="doc/colors/Tomato.png" height="16" style="vertical-align:middle"> Tomato | `#FFFF6347` | <img src="doc/colors/Turquoise.png" height="16" style="vertical-align:middle"> Turquoise | `#FF40E0D0` |
| <img src="doc/colors/Violet.png" height="16" style="vertical-align:middle"> Violet | `#FFEE82EE` | <img src="doc/colors/Wheat.png" height="16" style="vertical-align:middle"> Wheat | `#FFF5DEB3` | <img src="doc/colors/White.png" height="16" style="vertical-align:middle"> White | `#FFFFFFFF` |
| <img src="doc/colors/WhiteSmoke.png" height="16" style="vertical-align:middle"> WhiteSmoke | `#FFF5F5F5` | <img src="doc/colors/Yellow.png" height="16" style="vertical-align:middle"> Yellow | `#FFFFFF00` | <img src="doc/colors/YellowGreen.png" height="16" style="vertical-align:middle"> YellowGreen | `#FF9ACD32` |

¹ `Transparent` — RGB `#FFFFFF`, alpha = 0 (fully transparent).

</details>

```go
c := color.Crimson          // 0xFFDC143C
c  = color.DodgerBlue       // 0xFF1E90FF
c  = color.Transparent      // 0x00FFFFFF
```

## Creating colors

### From RGB / ARGB / RGBA components

```go
c := color.FromRGB(255, 0, 0)                  // fully opaque red
c  = color.FromARGB(128, 255, 0, 0)            // semi-transparent red
c  = color.FromRGBA(255, 0, 0, 128)            // same, RGBA argument order
```

### From HSL / HSLA components

```go
c, err := color.FromHSL(0, 1.0, 0.5)          // red  (hue=0°, sat=100%, light=50%)
c, err  = color.FromHSLA(120, 1.0, 0.5, 255)  // fully opaque green
```

`hue` must be in **[0, 360)**, `saturation` and `lightness` must be in **[0, 1]**.

## Reading color components

```go
c := color.CornflowerBlue

r := c.Red()        // uint8
g := c.Green()      // uint8
b := c.Blue()       // uint8
a := c.Alpha()      // uint8  (0 = transparent, 255 = opaque)

h := c.Hue()        // float64  [0, 360)
s := c.Saturation() // float64  [0, 1]
l := c.Lightness()  // float64  [0, 1]
```

## Formatting

`Format` converts a `Color` to a string in the requested layout:

| Constant | Example output |
|---|---|
| `color.INT` | `4278190335` |
| `color.HARGB` | `#FF0000FF` |
| `color.ARGB` | `FF0000FF` |
| `color.HRGB` | `#0000FF` |
| `color.RGB` | `0000FF` |
| `color.HRGBA` | `#0000FFFF` |
| `color.RGBA` | `0000FFFF` |
| `color.HSL` | `hsl(240, 100%, 50%)` |
| `color.HSLA` | `hsla(240, 100%, 50%, 1)` |
| `color.CSSRGB` | `rgb(0, 0, 255)` |
| `color.CSSRGBA` | `rgba(0, 0, 255, 1)` |
| `color.CSS4RGB` | `rgb(0 0 255)` |
| `color.CSS4RGBA` | `rgba(0 0 255 / 1)` |

```go
c := color.Blue

fmt.Println(c.Format(color.HRGB))    // #0000FF
fmt.Println(c.Format(color.HSL))     // hsl(240, 100%, 50%)
fmt.Println(c.Format(color.CSSRGBA)) // rgba(0, 0, 255, 1)
fmt.Println(c.String())              // FF0000FF  (default: AARRGGBB)
```

## Parsing

`Parse` converts a formatted string back to a `Color`:

```go
c, err := color.Parse(color.HRGB,    "#FF0000")
c, err  = color.Parse(color.HARGB,   "#FF0000FF")
c, err  = color.Parse(color.HSL,     "hsl(0, 100%, 50%)")
c, err  = color.Parse(color.CSSRGBA, "rgba(255, 0, 0, 1)")
```

## Color-space conversion helpers

```go
r, g, b        := color.HSLToRGB(0, 1.0, 0.5)   // → 255, 0, 0
h, s, l        := color.RGBToHSL(255, 0, 0)     // → 0, 1, 0.5
```

## JSON

`Color` implements `json.Marshaler` and `json.Unmarshaler`.  
Serialized format is `#RRGGBBAA`.

```go
type Widget struct {
    Background color.Color `json:"background"`
}

// {"background":"#FF0000FF"}
```

## SQL

`Color` implements `driver.Valuer` and `sql.Scanner`.  
The database value is a `#RRGGBB` hex string.

```go
var c color.Color
row.Scan(&c)           // reads "#FF0000" → color.Red
db.Exec("...", c)      // stores "#FF0000"
```

## License

See [LICENSE](LICENSE).