package aoc2019

import (
	"testing"

	"github.com/janreggie/AdventOfCode/structs"
	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	testCases := []structs.TestCase{
		{Input: "14",
			Result1: "2",
			Result2: "2"},
		{Input: "1969",
			Result1: "654",
			Result2: "966"},
		{Input: "100756",
			Result1: "33583",
			Result2: "50346"},
		{Details: "Y2019D01 my input",
			Input: "109067\n75007\n66030\n93682\n83818\n108891\n" +
				"139958\n129246\n80272\n119897\n112804\n69495\n" +
				"95884\n85402\n148361\n75986\n120063\n127683\n" +
				"146962\n76907\n61414\n98452\n134330\n53858\n" +
				"82662\n143258\n82801\n60279\n131782\n105989\n" +
				"102464\n96563\n71172\n113731\n90645\n94830\n" +
				"133247\n110149\n54792\n134863\n125919\n145490\n" +
				"69836\n108808\n87954\n148957\n110182\n126668\n" +
				"148024\n96915\n117727\n147378\n75967\n91915\n" +
				"60130\n85331\n66800\n103419\n72627\n72687\n" +
				"61606\n113160\n107082\n110793\n61589\n105005\n" +
				"73952\n65705\n117243\n140944\n117091\n113482\n" +
				"91379\n148185\n113853\n119822\n78179\n85407\n" +
				"119886\n109230\n68783\n63914\n51101\n93549\n" +
				"53361\n127984\n106315\n54997\n138941\n81075\n" +
				"120272\n120307\n98414\n115245\n105649\n89793\n" +
				"88421\n121104\n97084\n56928\n",
			Result1: "3330521",
			Result2: "4992931"},
	}
	for _, eachCase := range testCases {
		eachCase.Test(Day01, assert)
	}
}
