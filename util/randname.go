package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	// 定义形容词列表
	adjectives := []string{"美丽的", "不可一世的", "可怕的", "养猫的",
		"聪明的", "优雅的", "迷人的", "善良的", "快乐的", "勇敢的", "刻薄的",
		"幸福的", "无情的", "认贼作父的", "狡猾的", "谨慎的", "老练的", "摸不着头脑的",
		"霸气的", "在流汗的", "停不下来的", "自信的", "聪慧的", "机智的", "勤奋的",
		"毅力的", "豁达的", "耿直的", "直率的", "决心的", "会说英语的", "机灵的",
		"精明的", "讨厌香菜的", "热心的", "没有1美刀的", "真诚的", "拳头很大的", "怕热的", "被封印的", "宽容的"}

	adj := adjectives[rng.Intn(len(adjectives))]

	fantasyCreatures := []string{"哥布林", "小精灵", "矮人", "精灵女王",
		"地底怪兽", "水晶巨人", "独角兽", "恶魔猎手", "深渊魔王", "冰霜巨龙",
		"火焰元素", "树人", "黑暗法师", "天使", "魔术师", "海妖", "狼人",
		"石像鬼", "炎魔", "雷神", "食人魔", "死灵法师", "飞行巨蝠", "吸血鬼",
		"仙女", "骷髅战士", "史莱姆", "穴居人", "狂暴野兽", "魅魔"}

	noun := fantasyCreatures[rng.Intn(len(fantasyCreatures))]

	return adj + noun
}
