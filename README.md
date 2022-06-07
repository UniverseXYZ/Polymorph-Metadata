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
const SHOES_GENES_COUNT int = 25
const PANTS_GENES_COUNT int = 33
const TORSO_GENES_COUNT int = 34
const EYEWEAR_GENES_COUNT int = 13
const HEAD_GENES_COUNT int = 32
const WEAPON_RIGHT_GENES_COUNT int = 32
const WEAPON_LEFT_GENES_COUNT int = 32
```

## Variables
- app/domain/metadata/generator.go
  - GCLOUD_UPLOAD_BUCKET_NAME = "polymorph-images_test"
  - GCLOUD_SOURCE_BUCKET_NAME = "polymorph-source-imgs_test"
- app/domain/metadata/model.go
  - POLYMORPH_IMAGE_URL string = "https://storage.googleapis.com/polymorph-images_test/"

## Deployment
```bash
gcloud functions deploy images-function-test --entry-point TokenMetadata --runtime go116 --trigger-http --allow-unauthenticated --update-env-vars NODE_URL=https://rinkeby.infura.io/v3/72c7b07ef4c44fa79845fbbd526412ed,CONTRACT_ADDRESS=0x0c64049fd11bfb5c194662a4a6ED9BA7D6A389cA,DB_URL=polymorphraritydevclust.fyvje.mongodb.net/myFirstDatabase,USERNAME=thevikk,PASSWORD=2JkRAigzcaESkalt,POLYMORPH_DB=polymorphs-rarity-rinkeby-prod5,RARITY_COLLECTION=rarities-v2
```