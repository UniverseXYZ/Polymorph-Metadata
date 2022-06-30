# Polymorphic Faces
## Overview
- This branch is based on master's Polymorphs but is modified for Polymorphic Faces metadata
- Overview of contracts: [https://github.com/UniverseXYZ/Polymorphic-Faces-Contracts/](https://github.com/UniverseXYZ/Polymorphic-Faces-Contracts/)
## Genes interpretation
Each gene is 2 numbers read from right to left. Interpret follows:
- Gene 1 [0:2] - Background attribute. 
- Gene 2 [2:4] - Hair Left attribute
- Gene 3 [4:6] - Hair Right attribute
- Gene 4 [6:8] - Ear Left attribute
- Gene 5 [8:10] - Ear Right attribute
- Gene 6 [10:12] - Eye Left attribute
- Gene 7 [12:14] - Eye Right attribute
- Gene 8 [14:16] - Beard Top Left attribute
- Gene 9 [16:18] - Beard Top Right attribute
- Gene 10 [18:20] - Lips Left attribute
- Gene 11 [20:22] - Lips Right attribute
- Gene 12 [22:24] - Beard Bottom Left attribute
- Gene 13 [24:26] - Beard Bottom Right attribute

## Genes and their variations
```
const GENES_COUNT = 13
const BACKGROUND_GENE_COUNT int = 34
const HAIR_LEFT int = 34
const HAIR_RIGHT int = 34
const EAR_LEFT int = 34
const EAR_RIGHT int = 34
const BEARD_TOP_LEFT int = 34
const BEARD_TOP_RIGHT int = 34
const LIPS_LEFT int = 34
const LIPS_RIGHT int = 34
const BEARD_BOTTOM_LEFT int = 34
const BEARD_BOTTOM_RIGHT int = 34
const EYE_RIGHT int = 34
const EYE_LEFT int = 34
```