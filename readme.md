Package `d3dmath` provides vector and matrix functions similar to those in [D3DX](https://learn.microsoft.com/en-us/windows/win32/direct3d9/d3dx), a helper library for doing 3D graphics maths.

It is designed to work well with [d3d9](https://github.com/gonutz/d3d9).

Matrices are layed out in row-major order, meaning a matrix defined as:

```
Mat4 {
	 1,  2,  3,  4,
	 5,  6,  7,  8,
	 9, 10, 11, 12,
	13, 14, 15, 16,
}
```

represents the matrix you would expect by reading the source code, i.e.:

```
 1  2  3  4
 5  6  7  8
 9 10 11 12
13 14 15 16
```

Vectors are row-vectors, so the vector:

```
Vec3{1, 2, 3}
```

is the vector:

```
1 2 3
```

and not the vector:

```
1
2    <- not this
3
```

Functions that work with view transformations, like `LookAt` or `Perspective` assume a left-handed coordinate system.
