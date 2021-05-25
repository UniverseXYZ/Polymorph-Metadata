# Polymorph
## Overview
See general overview here: [https://github.com/LimeChain/Polymorph-Contracts/](https://github.com/LimeChain/Polymorph-Contracts/)
## Genes interpretation
Each gene is 2 numbers read from right to left. Interpret follows:
- Gene 1 [0:2] - base character. Will not be morphable 
- Gene 2 [2:4] - head attribute
- Gene 3 [4:6] - chest attribute
- Gene 4 [6:8] - pants attribute
- Gene 5 [8:10] - shoes attribute
- Gene 6 [10:12] - weapon attribute
- Gene 7 [12:14] - face attribute
- Gene 8 [14:16] - indicates whether tattoo should be applied. The tattoo is different for evert character
- Gene 9 [16:18] - background attribute