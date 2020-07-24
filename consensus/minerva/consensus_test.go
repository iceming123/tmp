// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package minerva

import (
	"encoding/json"
	"github.com/truechain/truechain-engineering-code/common"
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"
	"fmt"
	"github.com/truechain/truechain-engineering-code/common/math"
	"github.com/truechain/truechain-engineering-code/core/types"
	"github.com/truechain/truechain-engineering-code/params"
	osMath "math"
)

var (
	FrontierBlockReward = big.NewInt(5e+18) // Block reward in wei for successfully mining a block
	//SnailBlockRewardsInitial Snail block rewards initial 116.48733*10^18
	SnailBlockRewardsInitial = new(big.Int).Mul(big.NewInt(11648733), big.NewInt(1e13))
)

type diffTest struct {
	ParentTimestamp    uint64
	ParentDifficulty   *big.Int
	CurrentTimestamp   uint64
	CurrentBlocknumber *big.Int
	CurrentDifficulty  *big.Int
}

func (d *diffTest) UnmarshalJSON(b []byte) (err error) {
	var ext struct {
		ParentTimestamp    string
		ParentDifficulty   string
		CurrentTimestamp   string
		CurrentBlocknumber string
		CurrentDifficulty  string
	}
	if err := json.Unmarshal(b, &ext); err != nil {
		return err
	}

	d.ParentTimestamp = math.MustParseUint64(ext.ParentTimestamp)
	d.ParentDifficulty = math.MustParseBig256(ext.ParentDifficulty)
	d.CurrentTimestamp = math.MustParseUint64(ext.CurrentTimestamp)
	d.CurrentBlocknumber = math.MustParseBig256(ext.CurrentBlocknumber)
	d.CurrentDifficulty = math.MustParseBig256(ext.CurrentDifficulty)

	return nil
}

func TestAccountDiv(t *testing.T) {
	r := new(big.Int)
	println(r.Uint64())
	r = big.NewInt(600077777777777)
	println(r.Uint64())
	r.Div(r, big2999999)
	println(r.Uint64(), FrontierBlockReward.Uint64(), SnailBlockRewardsInitial.Bytes())
	fmt.Printf("%v", new(big.Int).Exp(new(big.Int).SetInt64(2),
		new(big.Int).Div(new(big.Int).Add(new(big.Int).SetInt64(5000), new(big.Int).SetInt64(12)), new(big.Int).SetInt64(5000)), nil))
}

//Calculate the reward distribution corresponding to the slow block height
//There is a new distribution incentive for every 4,500 blocks.
//The unit of output is wei
//6 bits at the end are cleared
func TestSnailAwardForHeight(t *testing.T) {
	for i := 1; i < 1000; i++ {
		snailBlockNumber := new(big.Int).SetInt64(int64(1 + 4500*(i-1)))
		fmt.Println("snailBlockNumber:", snailBlockNumber, "Award:", getCurrentCoin(snailBlockNumber))
		committeeAward, minerAward, minerFruitAward, _ := GetBlockReward(snailBlockNumber)
		fmt.Println("committeeAward:", committeeAward, "minerAward:", minerAward, "minerFruitAward:", minerFruitAward)
	}
}
func TestReward2(t *testing.T) {
	fmt.Println("addr:",types.FoundationAddress.String())
	snailNum := NewRewardBegin
	allReward := big.NewInt(0)
	snailReward := big.NewInt(0)
	rewardLimit := new(big.Int).Mul(big.NewInt(20000000),BaseBig)

	for i := 1; i < 2000000; i++ {
		num := big.NewInt(int64(i+snailNum))
		snailReward1 := getRewardCoin(num)
		if num.Cmp(big.NewInt(int64(NewRewardBegin+RewardEndSnailHeight))) >= 0{
			fmt.Println("last pos1:",i+1)
			break
		}
		allReward = new(big.Int).Add(allReward,snailReward1)
		if allReward.Cmp(rewardLimit) >= 0 {
			fmt.Println("last pos2:",i+1)
			break
		}
		if snailReward1.Cmp(snailReward) != 0 {
			fmt.Println("pos:",i+1,"preReward:",snailReward,"reward:",snailReward1)
			fmt.Println("pos:",i+1,"preReward:",toTrueCoin(snailReward).Text('f',6),
			"reward:",toTrueCoin(snailReward1).Text('f',6))
			snailReward = snailReward1

			cc, mm, mf,fc,_ := GetBlockReward3(num)
			fmt.Println("committeeAward:", cc, "minerAward:", mm, 
				"minerFruitAward:", mf,"found",fc)
			fmt.Println("committeeAward:", toTrueCoin(cc).Text('f',6), "minerAward:", toTrueCoin(mm).Text('f',6), 
				"minerFruitAward:", toTrueCoin(mf).Text('f',6),"found",toTrueCoin(fc).Text('f',6))
		}
	}
	fmt.Println("allReward",allReward)
	fmt.Println("allReward",toTrueCoin(allReward).Text('f',10))

	fmt.Println("finish")
}
func toTrueCoin(val *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(val),new(big.Float).SetInt(BaseBig))
}

func TestTime(t *testing.T) {
	t1 := time.Now()
	time.Sleep(time.Millisecond * time.Duration(600))
	t2 := time.Now()
	d := t2.Sub(t1)
	fmt.Println("d:",d.Seconds())
	if d.Seconds() > float64(0.5) {
		fmt.Println("good")
	}
	fmt.Println("finish")
}