// Copyright © 2019 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

package openfoodfacts

// Nutriment is the collection of all retruned nutriment associated with the product
type Nutriment struct {
	Salt                    float64 `json:"salt"`
	Salt100G                float64 `json:"salt_100g"`
	SaltValue               float64 `json:"salt_value"`
	SaltServing             float64 `json:"salt_serving"`
	SaltUnit                string  `json:"salt_unit"`
	Sugars100G              float64 `json:"sugars_100g"`
	Sugars                  float64 `json:"sugars"`
	SugarsUnit              string  `json:"sugars_unit"`
	SugarsServing           float64 `json:"sugars_serving"`
	SugarsValue             float64 `json:"sugars_value"`
	Iron                    float64 `json:"iron"`
	IronValue               float64 `json:"iron_value"`
	IronLabel               string  `json:"iron_label"`
	IronUnit                string  `json:"iron_unit"`
	Iron100G                float64 `json:"iron_100g"`
	IronServing             float64 `json:"iron_serving"`
	CalciumUnit             string  `json:"calcium_unit"`
	CalciumServing          float64 `json:"calcium_serving"`
	Calcium                 float64 `json:"calcium"`
	CalciumValue            float64 `json:"calcium_value"`
	CalciumLabel            string  `json:"calcium_label"`
	Calcium100G             float64 `json:"calcium_100g"`
	SaturatedFat100G        float64 `json:"saturated-fat_100g"`
	SaturatedFatServing     float64 `json:"saturated-fat_serving"`
	SaturatedFat            float64 `json:"saturated-fat"`
	SaturatedFatValue       float64 `json:"saturated-fat_value"`
	SaturatedFatUnit        string  `json:"saturated-fat_unit"`
	Fat100G                 float64 `json:"fat_100g"`
	FatServing              float64 `json:"fat_serving"`
	FatValue                float64 `json:"fat_value"`
	FatUnit                 string  `json:"fat_unit"`
	Fat                     float64 `json:"fat"`
	TransFatLabel           string  `json:"trans-fat_label"`
	TransFatUnit            string  `json:"trans-fat_unit"`
	TransFat                float64 `json:"trans-fat"`
	TransFat100G            float64 `json:"trans-fat_100g"`
	TransFatServing         float64 `json:"trans-fat_serving"`
	TransFatValue           float64 `json:"trans-fat_value"`
	VitaminA                float64 `json:"vitamin-a"`
	VitaminA100G            float64 `json:"vitamin-a_100g"`
	VitaminAValue           float64 `json:"vitamin-a_value"`
	VitaminAServing         float64 `json:"vitamin-a_serving"`
	VitaminAUnit            string  `json:"vitamin-a_unit"`
	VitaminALabel           string  `json:"vitamin-a_label"`
	VitaminCValue           float64 `json:"vitamin-c_value"`
	VitaminCUnit            string  `json:"vitamin-c_unit"`
	VitaminC100G            float64 `json:"vitamin-c_100g"`
	VitaminC                float64 `json:"vitamin-c"`
	VitaminCServing         float64 `json:"vitamin-c_serving"`
	VitaminCLabel           string  `json:"vitamin-c_label"`
	Proteins100G            float64 `json:"proteins_100g"`
	ProteinsServing         float64 `json:"proteins_serving"`
	ProteinsValue           float64 `json:"proteins_value"`
	ProteinsUnit            string  `json:"proteins_unit"`
	Proteins                float64 `json:"proteins"`
	Sodium                  float64 `json:"sodium"`
	SodiumServing           float64 `json:"sodium_serving"`
	SodiumValue             float64 `json:"sodium_value"`
	Sodium100G              float64 `json:"sodium_100g"`
	SodiumUnit              string  `json:"sodium_unit"`
	CarbohydratesUnit       string  `json:"carbohydrates_unit"`
	CarbohydratesValue      float64 `json:"carbohydrates_value"`
	Carbohydrates100G       float64 `json:"carbohydrates_100g"`
	Carbohydrates           float64 `json:"carbohydrates"`
	CarbohydratesServing    float64 `json:"carbohydrates_serving"`
	AlcoholValue            float64 `json:"alcohol_value"`
	AlcoholServing          float64 `json:"alcohol_serving"`
	AlcoholUnit             string  `json:"alcohol_unit"`
	Alcohol100G             float64 `json:"alcohol_100g"`
	Alcohol                 float64 `json:"alcohol"`
	NovaGroup               float64 `json:"nova-group"`
	NovaGroupServing        float64 `json:"nova-group_serving"`
	NovaGroup100G           float64 `json:"nova-group_100g"`
	Energy                  float64 `json:"energy"`
	EnergyServing           float64 `json:"energy_serving"`
	EnergyKcalServing       float64 `json:"energy-kcal_serving"`
	EnergyKcal              float64 `json:"energy-kcal"`
	Energy100G              float64 `json:"energy_100g"`
	EnergyUnit              string  `json:"energy_unit"`
	EnergyKcalValue         float64 `json:"energy-kcal_value"`
	EnergyKcalUnit          string  `json:"energy-kcal_unit"`
	EnergyKcal100G          float64 `json:"energy-kcal_100g"`
	EnergyValue             float64 `json:"energy_value"`
	NutritionScoreUk100G    float64 `json:"nutrition-score-uk_100g"`
	NutritionScoreFrServing float64 `json:"nutrition-score-fr_serving"`
	NutritionScoreFr        float64 `json:"nutrition-score-fr"`
	NutritionScoreFr100G    float64 `json:"nutrition-score-fr_100g"`
	NutritionScoreUkServing float64 `json:"nutrition-score-uk_serving"`
	NutritionScoreUk        float64 `json:"nutrition-score-uk"`
	Fiber                   float64 `json:"fiber"`
	Fiber100G               float64 `json:"fiber_100g"`
	FiberValue              float64 `json:"fiber_value"`
	FiberServing            float64 `json:"fiber_serving"`
	FiberUnit               string  `json:"fiber_unit"`
}
