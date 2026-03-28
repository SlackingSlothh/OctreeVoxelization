package voxel

import "github.com/SlackingSlothh/OctreeVoxelization/src/geometry"

type OctreeNode struct {
	Boundary geometry.Cube
	Triangles []geometry.Triangle3D
	Children []OctreeNode
}

const MAX_DEPTH = 6
const MIN_EDGE = 0.005

func (node *OctreeNode) Construct(currentDepth int) int {
	if len(node.Triangles) == 0 {
		node.Children = nil
		return currentDepth
	}
	if currentDepth >= MAX_DEPTH || node.Boundary.GetEdgeLength() <= MIN_EDGE {
		node.Children = nil
		return currentDepth
	}

	node.Children = nil
	dividedCube := node.Boundary.DivideCube()
	maxDepth := currentDepth
	for _, childCube := range dividedCube {
		childNode := OctreeNode{Boundary: childCube}
		for _, triangle := range node.Triangles {
			if triangle.IsIntersecting(childCube) {
				childNode.Triangles = append(childNode.Triangles, triangle)
			}
		}
		if len(childNode.Triangles) == 0 {
			continue
		}
		childDepth := childNode.Construct(currentDepth + 1)
		if childDepth > maxDepth {
			maxDepth = childDepth
		}
		node.Children = append(node.Children, childNode)
	}
	return maxDepth
}
