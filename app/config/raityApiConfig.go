package config

import "github.com/polymorph-metadata/constants"

var MORPHS_NO_PROJECTION_FIELDS []string = []string{
	constants.MorphFieldNames.ObjId,
	constants.MorphFieldNames.TokenId,
	constants.MorphFieldNames.CurrentGene,
	constants.MorphFieldNames.Headwear,
	constants.MorphFieldNames.Eyewear,
	constants.MorphFieldNames.Torso,
	constants.MorphFieldNames.Pants,
	constants.MorphFieldNames.Footwear,
	constants.MorphFieldNames.LeftHand,
	constants.MorphFieldNames.RightHand,
	constants.MorphFieldNames.Character,
	constants.MorphFieldNames.MainSetName,
	constants.MorphFieldNames.SecSetName,
	constants.MorphFieldNames.IsVirgin,
	constants.MorphFieldNames.ColorMismatches,
	constants.MorphFieldNames.MainSetName,
	constants.MorphFieldNames.MainMatchingTraits,
	constants.MorphFieldNames.SecSetName,
	constants.MorphFieldNames.SecMatchingTraits,
	constants.MorphFieldNames.HasCompletedSet,
	constants.MorphFieldNames.HandsScaler,
	constants.MorphFieldNames.HandsSetName,
	constants.MorphFieldNames.MatchingHands,
	constants.MorphFieldNames.NoColorMismatchScaler,
	constants.MorphFieldNames.ColorMismatchScaler,
	constants.MorphFieldNames.DegenScaler,
	constants.MorphFieldNames.VirginScaler,
	constants.MorphFieldNames.BaseRarity,
	constants.MorphFieldNames.Scrambles,
	constants.MorphFieldNames.Morphs,
	constants.MorphFieldNames.OldGenes,
}

const RESULTS_LIMIT int64 = 10000
