# Polymorph
## Overview
See general overview here: [https://github.com/LimeChain/Polymorph-Contracts/](https://github.com/LimeChain/Polymorph-Contracts/)
## Genes interpretation
Each gene is 2 numbers read from right to left. Interpret follows:
- Gene 1 [0:2] - base character. Will not be morphable 
- Gene 2 [2:4] - background attribute
- Gene 3 [4:6] - pants attribute
- Gene 4 [6:8] - torso attribute
- Gene 5 [8:10] - shoes attribute
- Gene 6 [10:12] - face attribute
- Gene 7 [12:14] - head attribute
- Gene 8 [14:16] - right weapon attribute
- Gene 9 [16:18] - left weapon attribute

## Genes and their variations
```
const GENES_COUNT = 9
const BACKGROUND_GENE_COUNT int = 12
const BASE_GENES_COUNT int = 11	
const PANTS_GENES_COUNT int = 24
const TORSO_GENES_COUNT int = 22
const SHOES_GENES_COUNT int = 17
const FACE_GENES_COUNT int = 9
const HEAD_GENES_COUNT int = 28
const WEAPON_RIGHT_GENES_COUNT int = 5
const WEAPON_LEFT_GENES_COUNT int = 17
```