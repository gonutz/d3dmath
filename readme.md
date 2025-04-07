Package `d3dmath` provides vector and matrix functions similar to those in
[D3DX](https://learn.microsoft.com/en-us/windows/win32/direct3d9/d3dx), a
helper library for doing 3D graphics maths.

It is designed to work well with [d3d9](https://github.com/gonutz/d3d9).

The top-level package `github.com/gonutz/d3dmath` is only for backwards
compatibility with version 1.0.0.

There are sub-packages `github.com/gonutz/d3dmath/column_major/d3dmath` and
`github.com/gonutz/d3dmath/row_major/d3dmath` which sort matrices either in
column or in row major order. This lets you work in both modes in Direct3D.
