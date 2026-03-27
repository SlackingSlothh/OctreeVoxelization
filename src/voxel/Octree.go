package voxel

import "github.com/SlackingSlothh/OctreeVoxelization/src/geometry"

type OctreeNode struct {
	boundary geometry.Cube
	triangles []geometry.Triangle3D
	children []OctreeNode
}

const VOXEL_SIZE = 1.0

func (node OctreeNode) Construct() {
	if len(node.triangles) == 0 {
		node.children = nil
	} else if node.boundary.GetEdgeLength() <= VOXEL_SIZE {
		node.children = nil
	} else {
		dividedCube := node.boundary.DivideCube()
		for _, childCube := range dividedCube {
			childNode := OctreeNode{childCube, nil, nil}
			for _, triangle := range node.triangles {
				if triangle.IsIntersecting(childCube) {
					childNode.triangles = append(childNode.triangles, triangle)
				}
			}
			childNode.Construct()
		}
	}
}