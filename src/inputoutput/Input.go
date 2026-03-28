package inputoutput

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/SlackingSlothh/OctreeVoxelization/src/geometry"
	"github.com/SlackingSlothh/OctreeVoxelization/src/voxel"
	"github.com/golang/geo/r3"
)


func ReadObj(path string) (geometry.Cube, []geometry.Triangle3D) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var faces []voxel.Face
	vertices := make(map[int]r3.Vector)
	vertexNum := 0
	
	var minX, maxX, minY, maxY, minZ, maxZ float64

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if line[0] == 'v' {
			vertex, err := ParseVertex(line)
			if err != nil {continue}
			vertexNum++
			vertices[vertexNum] = vertex
			if (vertexNum == 1) {
				minX = vertex.X
				maxX = vertex.X
				minY = vertex.Y
				maxY = vertex.Y
				minZ = vertex.Z
				maxZ = vertex.Z
			} else {
				minX = min(minX, vertex.X)
				maxX = max(maxX, vertex.X)
				minY = min(minY, vertex.Y)
				maxY = max(maxY, vertex.Y)
				minZ = min(minZ, vertex.Z)
				maxZ = max(maxZ, vertex.Z)
			}
		} else if line[0] == 'f' {
			face, err := ParseFace(line)
			if err != nil {continue}
			faces = append(faces, face)
		}
	}

	var triangles []geometry.Triangle3D

	for _, face := range faces {
		var v1, v2, v3 r3.Vector
		if vector, ok := vertices[face.V1]; ok {
			v1 = vector
		} else {continue}
		if vector, ok := vertices[face.V2]; ok {
			v2 = vector
		} else {continue}
		if vector, ok := vertices[face.V3]; ok {
			v3 = vector
		} else {continue}
		triangles = append(triangles, geometry.Triangle3D{V1: v1, V2: v2, V3: v3})
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error encountered during file read: %s", err)
	}

	rootCubeSize := max(maxX - minX, maxY - minY, maxZ - minZ)
	rootMinPoint := r3.Vector{X: minX, Y: minY, Z: minZ}
	rootMaxPoint := rootMinPoint.Add(r3.Vector{X: rootCubeSize, Y: rootCubeSize, Z: rootCubeSize})
	return geometry.Cube{V0: rootMinPoint, V1: rootMaxPoint}, triangles
}


func ParseVertex(line string) (r3.Vector, error) {
	var x, y, z float64
	_, err := fmt.Sscanf(line, "v %f %f %f", &x, &y, &z)

	if err != nil {
		return r3.Vector{}, err
	}

	return r3.Vector{X: x, Y: y, Z: z}, nil
}

func ParseFace(line string) (voxel.Face, error) {
	var v1, v2, v3 int
	_, err := fmt.Sscanf(line, "f %d %d %d", &v1, &v2, &v3)

	if err != nil {
		return voxel.Face{}, err
	}

	return voxel.Face{V1: v1, V2: v2, V3: v3}, nil
}