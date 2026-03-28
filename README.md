# Tugas Kecil 2 IF2211 Strategi Algoritma Semester II tahun 2025/2026

This repository is a project assignment about Divide and Conquer algorithm. This program implements Divide and Conquer to convert .obj file consisting of vertices and triangular faces into voxelized model .obj file.

## Requirements

- Go 1.25+ installed
- Internet connectivity to fetch module dependencies (if needed)

## Install dependencies

From repository root:

```bash
cd d:\Akademik\Semester_4\StrategiAlgoritma\OctreeVoxelization
go mod tidy
```

## Compile

```bash
go build -o bin/voxelizer.exe ./src/main
```

## Run

### Option A: Direct `go run`

```bash
go run ./src/main
```

### Option B: Run compiled executable

```bash
bin/voxelizer
```

## Usage

1. Program displays a menu:
   - `1. Start`
   - `2. Quit`
2. Choose `1` to start.
3. Program shows the current working directory.
4. Enter input `.obj` path (relative to current directory), e.g. `test/contoh/line.obj`.
5. Enter max octree depth:
   - Negative means no maximum.
   - Invalid (non-integer) defaults to `6`.
6. Enter minimum voxel size:
   - Must be positive float.
   - Invalid or non-positive defaults to `0.01`.
7. Program starts timer and performs voxelization.
8. Program prints statistics:
   - voxel count
   - vertex count
   - face count
   - nodes per level
   - empty nodes per level
   - max depth reached
   - processing time
9. Enter output `.obj` path (relative) to save result.
10. Program writes the output and returns to menu.

## Contact Author

- Email 1: 13524045@std.stei.itb.ac.id
- Email 2: ahmadzakyrobbani@gmail.com
