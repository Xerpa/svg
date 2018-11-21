package svg

import mt "github.com/rustyoz/Mtransform"

type Rect struct {
	Id        string  `xml:"id,attr"`
	Width     float64 `xml:"width,attr"`
	Height    float64 `xml:"height,attr"`
	Transform string  `xml:"transform,attr"`
	Style     string  `xml:"style,attr"`
	Rx        string  `xml:"rx,attr"`
	Ry        string  `xml:"ry,attr"`
	X         float64 `xml:"x,attr"`
	Y         float64 `xml:"y,attr"`
	Title     string  `xml:"title"`

	transform mt.Transform
	group     *Group
}
