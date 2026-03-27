package voxel

import "github.com/SlackingSlothh/OctreeVoxelization/src/geometry"

type OctreeNode struct {
	boundary geometry.Cube
	triangles []geometry.Triangle3D
	isLeaf bool
	children []OctreeNode
}

